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

	config := models.LoadPolyNodeConfig()

	args := []string{}
	for _, arg := range os.Args {
		args = append(args, strings.ToLower(arg))
	}

	switch {
	case args[1] == "version":
		fmt.Println(constants.VERSION)
	case args[1] == "upgrade":
		err := upgradePolyNode(operatingSystem, arch)
		if err != nil {
			log.Fatal(err.Error())
		}
	case utilities.IsKnownCommand(args[1]):
		node.Handle(args[1:], operatingSystem, arch, config)
	default:
		fmt.Println(constants.HELP)
	}
}
