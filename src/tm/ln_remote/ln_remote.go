// Package main implements the "ln" command:
// "Links the image from the datastore (persistent images)"
// For IPFS, TODO
package main

import (
	"os"

	helpers "github.com/hsanjuan/one-ipfs/helpers"
)

func main() {
	helpers.TMCmdParseArgs(os.Args, "ln")
	helpers.ExitWithError("IPFS Datastore does not support persistent images")
}
