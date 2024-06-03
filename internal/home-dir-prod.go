//go:build prod

package internal

import (
	"os"
	"runtime"
)

var pathSeparator string
var polynHomeDir string

func init() {
	if runtime.GOOS == "windows" {
		pathSeparator = "\\"
		polynHomeDir = "C:\\Program Files\\PolyNode"
	} else {
		pathSeparator = "/"
		home := os.Getenv("HOME")
		polynHomeDir = home + "/.PolyNode"
	}
}
