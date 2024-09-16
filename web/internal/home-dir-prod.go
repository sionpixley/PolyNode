//go:build prod

package internal

import (
	"os"
	"runtime"
)

var PolyNodeHomeDir string

func init() {
	if runtime.GOOS == "windows" {
		polyNodeHomeDir = os.Getenv("LOCALAPPDATA") + "/Programs/PolyNode"
	} else {
		polyNodeHomeDir = os.Getenv("HOME") + "/.PolyNode"
	}
}
