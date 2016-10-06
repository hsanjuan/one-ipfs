package helpers

import (
	"encoding/xml"
	"fmt"
	"os"
	"strings"
)

const IPFSUrl = "/ip4/127.0.0.1/tcp/5001"

func ExitWithError(msg string) {
	fmt.Fprint(os.Stderr, msg)
	os.Exit(2)
}

type Image struct {
	XMLName xml.Name `xml:"IMAGE"`
	Source  string   `xml:"SOURCE"`
	Path    string   `xml:"PATH"`
}

type Datastore struct {
	XMLName xml.Name `xml:"DATASTORE"`
	DSMad   string   `xml:"DS_MAD"`
	TMMad   string   `xml:"TM_MAD"`
}

type ImgDescription struct {
	XMLName xml.Name `xml:"DS_DRIVER_ACTION_DATA"`
	Img     Image
	DS      Datastore
}

func ParseDsImgDump(dsImgDump string) ImgDescription {
	result := ImgDescription{}
	err := xml.Unmarshal([]byte(dsImgDump), &result)
	if err != nil {
		ExitWithError(err.Error())
	}
	return result
}

type DsCmdArgs struct {
	Cmd     string
	ImgId   string
	ImgDump ImgDescription
}

func DsCmdParseArgs(args []string) *DsCmdArgs {
	if len(args) < 3 {
		ExitWithError(
			"Not all arguments are provided to the driver command")
	}
	return &DsCmdArgs{
		Cmd:     args[0],
		ImgId:   args[2],
		ImgDump: ParseDsImgDump(args[1]),
	}
}

func ExtractSource(img ImgDescription) string {
	src := img.Img.Source
	if src == "" {
		ExitWithError("Must provide an IPFS address in the SOURCE field")
	}
	src = strings.TrimPrefix(src, "/ipfs/")
	return src
}
