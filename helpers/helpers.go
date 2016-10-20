package helpers

import (
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"os"
	"strings"

	"github.com/hsanjuan/go-ipfs-api"
)

// IPFSUrl is the multiaddress string where the IPFS daemon is listening.
const IPFSUrl = "/ip4/127.0.0.1/tcp/5001"

// ExitWithError logs an error to stderr and exits.
func ExitWithError(msg string) {
	LogError(msg)
	os.Exit(2)
}

// LogInfo logs a message to stderr prepended by INFO
func LogInfo(msg string) {
	Log(fmt.Sprintf("INFO: %s", msg))
}

// LogError logs a message to stderr prepended by ERROR
func LogError(msg string) {
	Log(fmt.Sprintf("ERROR: %s", msg))
}

// Log logs a message to stderr
func Log(msg string) {
	fmt.Fprintln(os.Stderr, msg)
}

// Image is used to parse several fields in the XML object of an Image.
type Image struct {
	XMLName xml.Name `xml:"IMAGE"`
	Source  string   `xml:"SOURCE"`
	Path    string   `xml:"PATH"`
}

// Datastore is used to parse fields in the XML object of a Datastore.
type Datastore struct {
	XMLName xml.Name `xml:"DATASTORE"`
	DSMad   string   `xml:"DS_MAD"`
	TMMad   string   `xml:"TM_MAD"`
}

// ImgDescription is used to parse fields in the XML of an Image description.
type ImgDescription struct {
	XMLName xml.Name `xml:"DS_DRIVER_ACTION_DATA"`
	Img     Image
	DS      Datastore
}

// ParseDsImgDump parses base64-encoded XML datastore image template which is
// passed to ds_mad commands
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

// DsCmdArgs holds values for arguments which are usually passed to
// Datastore manager (ds_mad) commands.
type DsCmdArgs struct {
	Cmd     string
	ImgID   string
	ImgDump ImgDescription
}

// DsCmdParseArgs parses the arguments provided to a transfer manager
// command and returns them as a DsCmdArgs object.
func DsCmdParseArgs(args []string) *DsCmdArgs {
	if len(args) < 3 {
		ExitWithError(
			"Not all arguments are provided to the driver command")
	}
	return &DsCmdArgs{
		Cmd:     args[0],
		ImgID:   args[2],
		ImgDump: ParseDsImgDump(args[1]),
	}
}

// TMCmdArgs holds values for arguments which are usually passed to
// transfer manager (tm_mad) commands.
type TMCmdArgs struct {
	Fe         string
	Source     string
	Host       string
	RemotePath string
	VMID       string
	DsID       string
	SnapID     string
}

// TMCmdParseArgs parses the arguments provided to a transfer manager
// command and returns them as a TMCmdArgs object.
func TMCmdParseArgs(args []string, op string) *TMCmdArgs {
	tmArgs := TMCmdArgs{}
	switch op {
	case "clone", "ln", "mvds":
		if len(args) < 5 {
			ExitWithError("Not enough arguments")
		}
		feSource := strings.Split(args[1], ":")
		hostRemotePath := strings.Split(args[2], ":")
		if len(feSource) != 2 || len(hostRemotePath) != 2 {
			ExitWithError("Arguments not properly formatted")
		}
		tmArgs.Fe = feSource[0]
		tmArgs.Source = feSource[1]
		tmArgs.Host = hostRemotePath[0]
		tmArgs.RemotePath = hostRemotePath[1]
		tmArgs.VMID = args[3]
		tmArgs.DsID = args[4]
	case "cpds":
		ExitWithError("test")
	default:
		ExitWithError("Operation not supported")
	}
	return &tmArgs
}

// ExtractIPFSID returns the value of the Source field of an Image description
// or falls back to the value of the Path field. It checks that the
// values are a valid URI (fs:/ipns/ or fs:/ipfs/) or exits with an error.
func ExtractIPFSID(img ImgDescription) string {
	imgSrc := img.Img.Source
	imgPath := img.Img.Path
	var src string
	if imgSrc != "" {
		src = imgSrc
	} else if imgPath != "" {
		src = imgPath
	} else {
		ExitWithError("Must provide an IPFS address as SOURCE or PATH")
	}

	if !strings.HasPrefix(src, "fs:/ipfs/") &&
		!strings.HasPrefix(src, "fs:/ipns/") &&
		!strings.HasPrefix(src, "/ipfs/") &&
		!strings.HasPrefix(src, "/ipns/") {
		ExitWithError("Wrong IPFS/IPNS path")
	}
	return strings.TrimPrefix(src, "fs:")
}

// Resolve receives an ipfsID in the form fs:/ipfs/[hash] or fs:/ipns/[id].
// It returns the [hash] for IPFS URIs. For IPNS URIs, it resolves them and
// returns the hash they are pointing to.
func Resolve(ipfsID string) string {
	if strings.HasPrefix(ipfsID, "/ipfs/") {
		return strings.TrimPrefix(ipfsID, "/ipfs/")
	} else if strings.HasPrefix(ipfsID, "/ipns/") {
		sh := shell.NewShell(IPFSUrl)
		hash, err := sh.Resolve(ipfsID)
		if err != nil {
			ExitWithError(fmt.Sprintf("Error resolving: %s", err))
		}
		return strings.TrimPrefix(hash, "/ipfs/")
	}
	ExitWithError(fmt.Sprintf("Wrong IPFS URI: %s", ipfsID))
	return ""
}
