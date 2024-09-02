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
		polynHomeDir = os.Getenv("LOCALAPPDATA") + "\\Programs\\PolyNode"
	} else {
		pathSeparator = "/"
		polynHomeDir = os.Getenv("HOME") + "/.PolyNode"
	}
}
