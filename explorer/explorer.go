package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func launchExplorer(filename string) {

	fmt.Println("filename :" + filename)
	parts := strings.Split(strings.TrimSpace(filename), "/")
	if parts[1] == "mnt" {
		drv := parts[2]
		if len(drv) == 1 {
			winFilename := drv + ":" + filename[6:]
			winFilename = strings.Replace(winFilename, "/", "\\", -1)
			cmd := exec.Command("/mnt/c/windows/explorer.exe", winFilename)
			cmd.Start()
		}
	}
}
func main() {
	//current directory
	dir, _ := os.Getwd()
	launchExplorer(dir)
}
