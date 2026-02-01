//go:build !prod

package internal

import "os"

var PathSeparator = string(os.PathSeparator)
var PolynHomeDir = "."
