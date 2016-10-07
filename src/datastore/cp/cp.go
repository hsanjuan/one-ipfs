// Package main implements the "cp" command: "copies/dumps the image to the
// datastore"
// For IPFS, we simply return the hash of the image (must be provided as SOURCE)
// in the image xml and the size.
package main

import (
	"fmt"
	"os"

	ipfs "github.com/hsanjuan/go-ipfs-api"
	helpers "github.com/hsanjuan/one-ipfs/helpers"
)

func main() {
	args := helpers.DsCmdParseArgs(os.Args)
	src := helpers.ExtractSource(args.ImgDump)
	sh := ipfs.NewShell(helpers.IPFSUrl)
	stat, err := sh.ObjectStat(src)
	if err != nil {
		helpers.ExitWithError("IPFS object not found")
	}
	fmt.Println(src, stat.CumulativeSize/(1024*1024)) // in MBs
}
