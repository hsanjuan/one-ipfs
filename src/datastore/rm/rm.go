// Package main implements the "rm" command: "removes an image from the datastore"
// For IPFS, we cannot do anything. But we unpin it in case it was pinned.
package main

import (
	"os"

	ipfs "github.com/hsanjuan/go-ipfs-api"
	helpers "github.com/hsanjuan/one-ipfs/helpers"
)

func main() {
	args := helpers.DsCmdParseArgs(os.Args)
	src := helpers.ExtractSource(args.ImgDump)
	sh := ipfs.NewShell(helpers.IPFSUrl)
	sh.Unpin(src) // Ignore errors.
}
