// Package main implements the "stat" command:
// "returns the size of an image in Mb"
package main

import (
	"fmt"
	"os"

	ipfs "github.com/hsanjuan/go-ipfs-api"
	helpers "github.com/hsanjuan/one-ipfs/helpers"
)

func main() {
	args := helpers.DsCmdParseArgs(os.Args)
	src := helpers.Resolve(helpers.ExtractIPFSID(args.ImgDump))
	sh := ipfs.NewShell(helpers.IPFSUrl)
	stat, err := sh.ObjectStat(src)
	if err != nil {
		helpers.ExitWithError("IPFS object not found")
	}
	fmt.Println(stat.CumulativeSize / (1024 * 1024)) // in MBs
}
