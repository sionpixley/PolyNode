//go:build !prod

package internal

import "os"

var PathSeparator string = string(os.PathSeparator)
var PolynHomeDir string = "."
