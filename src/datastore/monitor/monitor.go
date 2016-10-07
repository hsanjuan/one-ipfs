// Package main implements the "monitor" command: "monitors a datastore"
// For IPFS, we are going to return 0 for everything until we figure out a
// better way (i.e. listing and adding sizes for all images in DS).
package main

import (
	"fmt"
	"os"

	helpers "github.com/hsanjuan/one-ipfs/helpers"
)

func main() {
	// Make sure we have valid arguments anyway
	helpers.DsCmdParseArgs(os.Args)
	helpers.LogInfo("monitor: IPFS datastore monitoring is meaningless")
	fmt.Println("USED_MB=0")
	fmt.Println("TOTAL_MB=100000000")
	fmt.Println("FREE_MB=100000000")
}
