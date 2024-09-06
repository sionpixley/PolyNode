package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/sionpixley/PolyNode/internal"
)

func main() {
	operatingSystem := convertToOperatingSystem(runtime.GOOS)
	arch := convertToArchitecture(runtime.GOARCH)

	defer fmt.Println()

	if !isSupportedOperatingSystem(operatingSystem) {
		fmt.Println(internal.UNSUPPORTED_OS_ERROR)
		return
	} else if !isSupportedArchitecture(arch) {
		fmt.Println("unsupported CPU architecture")
		return
	}

	if len(os.Args) == 1 {
		fmt.Println(internal.HELP)
		return
	}

	config := internal.LoadPolyNodeConfig()

	args := []string{}
	for _, arg := range os.Args {
		args = append(args, strings.ToLower(arg))
	}

	if args[1] == "version" {
		fmt.Println(internal.VERSION)
	} else if internal.IsKnownCommand(args[1]) {
		internal.HandleNode(args[1:], operatingSystem, arch, config)
	} else {
		fmt.Println(internal.HELP)
	}
}

func convertToArchitecture(archStr string) internal.Architecture {
	switch archStr {
	case "amd64":
		return internal.X64
	case "arm64":
		return internal.ARM64
	default:
		return internal.NA_ARCH
	}
}

func convertToOperatingSystem(osStr string) internal.OperatingSystem {
	switch osStr {
	case "darwin":
		return internal.MAC
	case "linux":
		return internal.LINUX
	case "windows":
		return internal.WINDOWS
	default:
		return internal.NA_OS
	}
}

func isSupportedArchitecture(arch internal.Architecture) bool {
	return arch != internal.NA_ARCH
}

func isSupportedOperatingSystem(operatingSystem internal.OperatingSystem) bool {
	return operatingSystem != internal.NA_OS
}
