package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {
	browser := "/mnt/c/Program Files (x86)/Google/Chrome/Application/Chrome.exe"
	branchName, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD", "--").Output()
	if err != nil {
		log.Fatal("Error git cmd: ", err)
	} else {
		branchNameS := string(branchName)
		branchNameS = strings.Split(branchNameS, "\n")[0]
		fmt.Println(branchNameS)
		cmd := exec.Command(browser, "http://sh-w81-001:8080/search/?q="+string(branchNameS))
		cmd.Start()
	}
}
