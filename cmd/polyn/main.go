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

	defer internal.PrintOptionalLine(operatingSystem)

	if !internal.IsSupportedOperatingSystem(operatingSystem) {
		fmt.Println("Not a supported operating system.")
		return
	} else if !internal.IsSupportedArchitecture(arch) {
		fmt.Println("Not a supported CPU architecture.")
		return
	}

	if len(os.Args) == 1 {
		internal.PrintHelp()
		return
	}

	args := []string{}
	for _, arg := range os.Args {
		args = append(args, strings.ToLower(arg))
	}

	if internal.IsKnownCommand(args[1]) {
		internal.HandleNode(args[1:], operatingSystem, arch)
	} else {
		internal.PrintHelp()
	}
}

func printVersion() {
	fmt.Println(internal.VERSION)
}
