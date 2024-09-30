//go:build prod

package internal

import (
	"os"
	"runtime"
)

var PathSeparator string
var PolynHomeDir string

func init() {
	if runtime.GOOS == "windows" {
		PathSeparator = "\\"
		PolynHomeDir = os.Getenv("LOCALAPPDATA") + "\\Programs\\PolyNode"
	} else {
		PathSeparator = "/"
		PolynHomeDir = os.Getenv("HOME") + "/.PolyNode"
	}
}
