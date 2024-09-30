//go:build prod

package internal

import (
	"os"
	"runtime"
)

var PolyNodeHomeDir string

func init() {
	if runtime.GOOS == "windows" {
		PolyNodeHomeDir = os.Getenv("LOCALAPPDATA") + "/Programs/PolyNode"
	} else {
		PolyNodeHomeDir = os.Getenv("HOME") + "/.PolyNode"
	}
}
