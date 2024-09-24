package node

import (
	"fmt"
	"log"

	"github.com/sionpixley/PolyNode/internal/constants"
	"github.com/sionpixley/PolyNode/internal/models"
	"github.com/sionpixley/PolyNode/internal/utilities"
)

// Main function for Node.js actions.
func Handle(args []string, operatingSystem models.OperatingSystem, arch models.Architecture, config models.PolyNodeConfig) {
	if len(args) == 0 {
		fmt.Println(constants.HELP)
		return
	}

	var err error
	command := utilities.ConvertToCommand(args[0])
	switch command {
	case constants.ADD:
		if len(args) > 1 {
			err = add(convertKeywordToVersion(args[1], config), operatingSystem, arch, config)
		} else {
			fmt.Println(constants.HELP)
		}
	case constants.CURRENT:
		current()
	case constants.INSTALL:
		if len(args) > 1 {
			err = install(convertKeywordToVersion(args[1], config), operatingSystem, arch, config)
		} else {
			fmt.Println(constants.HELP)
		}
	case constants.LIST:
		list()
	case constants.REMOVE:
		if len(args) > 1 {
			err = remove(args[1])
		} else {
			fmt.Println(constants.HELP)
		}
	case constants.SEARCH:
		if len(args) > 1 {
			err = search(args[1], config)
		} else {
			err = searchDefault(config)
		}
	case constants.USE:
		if len(args) > 1 {
			err = use(args[1], operatingSystem)
		} else {
			fmt.Println(constants.HELP)
		}
	default:
		fmt.Println(constants.HELP)
	}

	if err != nil {
		log.Fatal(err.Error())
	}
}
