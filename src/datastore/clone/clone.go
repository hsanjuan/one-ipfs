// Package main implements the "clone" command: " clones an image"
// For IPFS, we no need to clone! The source for the new image is the same.
package main

import (
	"fmt"
	"os"

	ipfs "github.com/hsanjuan/go-ipfs-api"
	helpers "github.com/hsanjuan/one-ipfs/helpers"
)

func main() {
	args := helpers.DsCmdParseArgs(os.Args)
	ipfsId := helpers.ExtractIPFSID(args.ImgDump)
	src := helpers.Resolve(ipfsId)
	sh := ipfs.NewShell(helpers.IPFSUrl)
	_, err := sh.ObjectStat(src)
	if err != nil {
		helpers.ExitWithError("IPFS object not found")
	}
	fmt.Println(ipfsId) // Seems a valid source, so we keep it
}
