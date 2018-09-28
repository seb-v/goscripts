package tools

import (
	"strings"
)

func WslToWinPath(filename string) string {
	winFilename := ""
	parts := strings.Split(strings.TrimSpace(filename), "/")
	if parts[1] == "mnt" {
		drv := parts[2]
		if len(drv) == 1 {
			winFilename = drv + ":" + filename[6:]
			winFilename = strings.Replace(winFilename, "/", "\\", -1)
		}
	}
	return winFilename
}
