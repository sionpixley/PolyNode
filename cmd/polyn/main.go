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

	if !isSupportedOperatingSystem(operatingSystem) {
		log.Fatalln(constants.UnsupportedOSError)
	} else if !isSupportedArchitecture(arch) {
		log.Fatalln(constants.UnsupportedArchError)
	}

	if len(os.Args) == 1 {
		fmt.Println(constants.Help)
		return
	}

	config := models.LoadPolyNodeConfig()

	if config.AutoUpdate {
		e := autoUpdate(operatingSystem, arch)
		if e != nil {
			log.Fatalln(e.Error())
		}
	}

	args := make([]string, len(os.Args)-1)
	for i, arg := range os.Args {
		if i == 0 {
			continue
		} else {
			args[i-1] = strings.ToLower(arg)
		}
	}

	switch {
	case args[0] == "version":
		fmt.Println(constants.Version)
	case args[0] == "update":
		err := updatePolyNode(operatingSystem, arch)
		if err != nil {
			log.Fatalln(err.Error())
		}
	case utilities.IsKnownCommand(args[0]):
		node.Handle(args, operatingSystem, arch, config)
	default:
		fmt.Println(constants.Help)
	}
}
