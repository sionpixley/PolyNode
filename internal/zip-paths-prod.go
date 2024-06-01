//go:build prod

package internal

import (
	"os"
	"runtime"
)

var polynHomeDir string

func init() {
	if runtime.GOOS == "windows" {
		polynHomeDir = "C:\\Program Files\\PolyNode"
	} else {
		home := os.Getenv("HOME")
		polynHomeDir = home + "/.PolyNode"
	}
}
