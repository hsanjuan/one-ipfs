// Package main implements the "clone" command:
// "clones the image from the datastore (non-persistent images)"
// This script runs on the deployment host.
// We get the image from IPFS and place it in the system datastore
// as indicated by RemotePath.
package main

import (
	"os"
	"path/filepath"

	ipfs "github.com/hsanjuan/go-ipfs-api"
	helpers "github.com/hsanjuan/one-ipfs/helpers"
)

func main() {
	args := helpers.TMCmdParseArgs(os.Args, "clone")
	src := args.Source
	dst := args.RemotePath
	dir := filepath.Dir(dst)
	err := os.Mkdir(dir, 0700)
	sh := ipfs.NewShell(helpers.IPFSUrl)
	err = sh.Get(src, dst)
	if err != nil {
		helpers.LogError(err.Error())
		helpers.ExitWithError("clone: Error getting IPFS object")
	}
	// Exit with 0
}
