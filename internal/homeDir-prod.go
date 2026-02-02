//go:build prod

package internal

import (
	"os"
	"runtime"
)

var PathSeparator = string(os.PathSeparator)
var PolynHomeDir string

func init() {
	if runtime.GOOS == "windows" {
		PolynHomeDir = os.Getenv("LOCALAPPDATA") + "\\Programs\\PolyNode"
	} else {
		PolynHomeDir = os.Getenv("HOME") + "/.PolyNode"
	}
}
