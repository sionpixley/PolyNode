package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"

	"github.com/sionpixley/PolyNode/internal/constants"
	"github.com/sionpixley/PolyNode/internal/models"
	"github.com/sionpixley/PolyNode/internal/node"
	"github.com/sionpixley/PolyNode/internal/utilities"
	"github.com/sionpixley/PolyNode/pkg/polynrc"
)

func main() {
	operatingSystem := convertToOperatingSystem(runtime.GOOS)
	arch := convertToArchitecture(runtime.GOARCH)

	defer fmt.Println()

	if !isSupportedOperatingSystem(operatingSystem) {
		log.Fatal(constants.UNSUPPORTED_OS_ERROR)
	} else if !isSupportedArchitecture(arch) {
		log.Fatal(constants.UNSUPPORTED_ARCH_ERROR)
	}

	if len(os.Args) == 1 {
		fmt.Println(constants.HELP)
		return
	}

	config := polynrc.LoadPolyNodeConfig()

	args := []string{}
	for _, arg := range os.Args {
		args = append(args, strings.ToLower(arg))
	}

	if args[1] == "version" {
		fmt.Println(constants.VERSION)
	} else if args[1] == "upgrade" {
		err := utilities.UpgradePolyNode(operatingSystem, arch)
		if err != nil {
			log.Fatal(err.Error())
		}
	} else if utilities.IsKnownCommand(args[1]) {
		node.Handle(args[1:], operatingSystem, arch, config)
	} else {
		fmt.Println(constants.HELP)
	}
}

func convertToArchitecture(archStr string) models.Architecture {
	switch archStr {
	case "amd64":
		return constants.X64
	case "arm64":
		return constants.ARM64
	case "ppc64":
		return constants.PPC64
	default:
		return constants.NA_ARCH
	}
}

func convertToOperatingSystem(osStr string) models.OperatingSystem {
	switch osStr {
	case "aix":
		return constants.AIX
	case "darwin":
		return constants.MAC
	case "linux":
		return constants.LINUX
	case "windows":
		return constants.WINDOWS
	default:
		return constants.NA_OS
	}
}

func isSupportedArchitecture(arch models.Architecture) bool {
	return arch != constants.NA_ARCH
}

func isSupportedOperatingSystem(operatingSystem models.OperatingSystem) bool {
	return operatingSystem != constants.NA_OS
}
