package helpers

import (
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"os"
	"strings"
)

const IPFSUrl = "/ip4/127.0.0.1/tcp/5001"

func ExitWithError(msg string) {
	LogError(msg)
	os.Exit(2)
}

func LogInfo(msg string) {
	Log(fmt.Sprintf("INFO: %s", msg))
}

func LogError(msg string) {
	Log(fmt.Sprintf("ERROR: %s", msg))
}

func Log(msg string) {
	fmt.Fprintln(os.Stderr, msg)
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
	data, err := base64.StdEncoding.DecodeString(dsImgDump)
	if err != nil {
		ExitWithError("Error decoding base64 string")
	}

	result := ImgDescription{}
	err = xml.Unmarshal([]byte(data), &result)
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

type TMCmdArgs struct {
	Fe         string
	Source     string
	Host       string
	RemotePath string
	VMId       string
	DSId       string
	SnapId     string
}

func TMCmdParseArgs(args []string, op string) *TMCmdArgs {
	tmArgs := TMCmdArgs{}
	switch op {
	case "clone", "ln", "mvds":
		if len(args) < 5 {
			ExitWithError("Not enough arguments")
		}
		fe_source := strings.Split(args[1], ":")
		host_rp := strings.Split(args[2], ":")
		if len(fe_source) != 2 || len(host_rp) != 2 {
			ExitWithError("Arguments not properly formatted")
		}
		tmArgs.Fe = fe_source[0]
		tmArgs.Source = fe_source[1]
		tmArgs.Host = host_rp[0]
		tmArgs.RemotePath = host_rp[1]
		tmArgs.VMId = args[3]
		tmArgs.DSId = args[4]
	case "cpds":
		ExitWithError("test")
	default:
		ExitWithError("Operation not supported")
	}
	return &tmArgs
}

func ExtractSource(img ImgDescription) string {
	src := img.Img.Source
	path := img.Img.Path
	if src != "" {
		src = strings.TrimPrefix(src, "/ipfs/")
	} else if path != "" {
		// Upon clone, OpenNebula places SOURCE into PATH
		// Maybe a bug who knows.
		src = strings.TrimPrefix(path, "/ipfs/")
	} else {
		ExitWithError("Must provide an IPFS address as SOURCE or PATH")
	}
	return src
}
