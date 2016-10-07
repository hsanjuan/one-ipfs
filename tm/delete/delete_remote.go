// Package main implements the "delete" command:
// " removes the either system datastore’s directory of the VM or a disk itself."
// This script runs on the deployment host.
// We simply delete the disk.
package main

import (
	"os"
	"strings"

	helpers "github.com/hsanjuan/one-ipfs/helpers"
)

func main() {
	if len(os.Args) < 4 {
		helpers.ExitWithError("Wrong number of arguments")
	}
	host_disk := strings.Split(os.Args[1], ":")
	disk := host_disk[1]
	err := os.Remove(disk)
	if err != nil { // Just log an error if it happens
		helpers.LogInfo(err.Error())
	}
}
