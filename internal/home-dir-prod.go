//go:build prod

package internal

import (
	"os"
)

var polynHomeDir string

func init() {
	home := os.Getenv("HOME")
	polynHomeDir = home + "/.PolyNode"
}
