//go:build !prod

package internal

import "runtime"

var PathSeparator string
var PolynHomeDir string = "."

func init() {
	if runtime.GOOS == "windows" {
		PathSeparator = "\\"
	} else {
		PathSeparator = "/"
	}
}
