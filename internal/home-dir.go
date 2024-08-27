//go:build !prod

package internal

import "runtime"

var pathSeparator string
var polynHomeDir string = "."

func init() {
	if runtime.GOOS == "windows" {
		pathSeparator = "\\"
	} else {
		pathSeparator = "/"
	}
}
