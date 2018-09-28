package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
	"github.com/seb-v/goscripts/tools"
)

func isValidConfig(s string) bool {
	s = strings.ToLower(s)
	return s == "release" || s == "debug" || s == "RelWithDebInfo"
}

func getBuildParams() (string, string) {
	// default build configuration
	config := "debug"
	target := "SceCoreTests"

	params := os.Args[1:]
	for i := len(params) - 1; i >= 0; i-- {
		if isValidConfig(params[i]) {
			config = params[i]
		} else {
			target = params[i]
		}
	}
	return config, target
}

func main() {

	config, target := getBuildParams()

	color.Cyan("*********************************")
	color.Cyan("Building " + target + " : " + config)
	color.Cyan("*********************************")

	rootPath, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
	if err == nil {
		buildPath := strings.TrimSuffix(string(rootPath), "\n") + "/sources/code/SceCore/build64"
		buildPath = tools.WslToWinPath(buildPath)
		color.Cyan("Build path : " + string(buildPath))

		cmd := exec.Command("cmake.exe", "--build", buildPath, "--target", target, "--", "/verbosity:minimal", "/property:configuration="+config)
		if err != nil {
			log.Fatal(err)
		}
		out, _ := cmd.StdoutPipe()
		rd := bufio.NewReader(out)
		if err := cmd.Start(); err != nil {
			log.Fatal("Buffer Error:", err)
		}

		for {
			str, err := rd.ReadString('\n')
			if err != nil {
				log.Fatal("Read Error:", err)
				return
			}
			if strings.Contains(str, "warning") {
				color.Yellow(str)
			} else if strings.Contains(str, "error") {
				color.Red(str)
			} else {
				fmt.Print(str)
			}
		}
	} else {
		log.Fatal(err)
	}

}
