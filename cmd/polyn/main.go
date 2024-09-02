package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/sionpixley/PolyNode/internal"
)

func main() {
	operatingSystem := internal.ConvertToOperatingSystem(runtime.GOOS)
	arch := internal.ConvertToArchitecture(runtime.GOARCH)

	defer optionallyPrintExtraLine(operatingSystem)

	if !internal.IsSupportedOperatingSystem(operatingSystem) {
		fmt.Println("Not a supported operating system.")
		return
	} else if !internal.IsSupportedArchitecture(arch) {
		fmt.Println("Not a supported CPU architecture.")
		return
	}

	if len(os.Args) == 1 {
		fmt.Println(internal.HELP)
		return
	}

	args := []string{}
	for _, arg := range os.Args {
		args = append(args, strings.ToLower(arg))
	}

	if args[1] == "version" {
		fmt.Println(internal.VERSION)
	} else if internal.IsKnownCommand(args[1]) {
		internal.HandleNode(args[1:], operatingSystem, arch)
	} else {
		fmt.Println(internal.HELP)
	}
}

// Command Prompt on Windows automatically adds an extra line to the output.
// Adding an extra line to the output for all shells except Command Prompt.
func optionallyPrintExtraLine(operatingSystem internal.OperatingSystem) {
	// The PSModulePath environment variable is only set in PowerShell.
	// If the OS is Windows, and the variable is not set, then it's running in Command Prompt.
	if operatingSystem == internal.WINDOWS && os.Getenv("PSModulePath") != "" {
		// Command Prompt already adds an extra line to the output, so do nothing.
	} else {
		fmt.Println()
	}
}
