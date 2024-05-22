package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/sionpixley/polyn/internal"
)

func main() {
	operatingSystem := internal.ConvertToOperatingSystem(runtime.GOOS)
	arch := internal.ConvertToArchitecture(runtime.GOARCH)

	if len(os.Args) == 1 {
		internal.PrintHelp(operatingSystem)
		return
	}

	args := []string{}
	for _, arg := range os.Args {
		args = append(args, strings.ToLower(arg))
	}

	runtime := args[1]
	switch runtime {
	case "bun":
		internal.HandleBun(args[2:], operatingSystem)
	case "deno":
		internal.HandleDeno(args[2:], operatingSystem)
	case "node":
		internal.HandleNode(args[2:], operatingSystem, arch)
	case "version":
		printVersion()
	default:
		if internal.IsKnownCommand(runtime) {
			// We default to Node.js.
			// Also, we slice starting with index 1 instead of 2 because the command is missing a runtime.
			internal.HandleNode(args[1:], operatingSystem, arch)
		} else {
			internal.PrintHelp(operatingSystem)
		}
	}
}

func printVersion() {
	fmt.Println(internal.VERSION)
}
