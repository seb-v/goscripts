package main

import (
	"os"
	"os/exec"

	"github.com/seb-v/goscripts/tools"
)

func launchExplorer(filename string) {
	winFilename := tools.WslToWinPath(filename)
	if winFilename != "" {
		cmd := exec.Command("/mnt/c/windows/explorer.exe", winFilename)
		cmd.Start()
	}
}

func main() {
	//current directory
	dir, _ := os.Getwd()
	launchExplorer(dir)
}
