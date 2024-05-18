package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/sionpixley/polyn/internal"
	"github.com/sionpixley/polyn/internal/bun"
	"github.com/sionpixley/polyn/internal/constants"
	"github.com/sionpixley/polyn/internal/deno"
	"github.com/sionpixley/polyn/internal/node"
)

func main() {
	if len(os.Args) == 1 {
		internal.PrintHelp()
		return
	}

	operatingSystem := internal.ConvertToOs(runtime.GOOS)

	args := []string{}
	for _, arg := range os.Args {
		args = append(args, strings.ToLower(arg))
	}

	env := args[1]
	switch env {
	case "bun":
		bun.Handle(args[2:], operatingSystem)
	case "deno":
		deno.Handle(args[2:], operatingSystem)
	case "node":
		node.Handle(args[2:], operatingSystem)
	case "version":
		printVersion()
	default:
		if internal.IsKnownCommand(env) {
			// We default to Node.js.
			// Also, we slice starting with index 1 instead of 2 because the command is missing an env.
			node.Handle(args[1:], operatingSystem)
		} else {
			internal.PrintHelp()
		}
	}
}

func printVersion() {
	fmt.Println(constants.VERSION)
}
