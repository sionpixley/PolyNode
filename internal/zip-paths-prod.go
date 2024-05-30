//go:build prod

package internal

import "runtime"

var polynHomeDir string

func init() {
	if runtime.GOOS == "windows" {
		polynHomeDir = "C:\\Program Files\\PolyNode"
	} else {
		polynHomeDir = "/opt/PolyNode"
	}
}

const c_POLYN_HOME_DIR string = "/opt/PolyNode"
