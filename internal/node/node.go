package node

import (
	"fmt"
	"log"

	"github.com/sionpixley/PolyNode/internal/constants"
	"github.com/sionpixley/PolyNode/internal/constants/command"
	"github.com/sionpixley/PolyNode/internal/models"
	"github.com/sionpixley/PolyNode/internal/utilities"
)

// Handle function is the main function for Node.js actions.
func Handle(args []string, operatingSystem models.OperatingSystem, arch models.Architecture, config models.PolyNodeConfig) {
	var err error
	comm := utilities.ConvertToCommand(args[0])
	switch comm {
	case command.Add:
		if len(args) > 1 {
			err = add(convertKeywordToVersion(args[1], operatingSystem, arch, config), operatingSystem, arch, config)
		} else {
			fmt.Println(constants.Help)
		}
	case command.Current:
		current()
	case command.Install:
		if len(args) > 1 {
			err = install(convertKeywordToVersion(args[1], operatingSystem, arch, config), operatingSystem, arch, config)
		} else {
			fmt.Println(constants.Help)
		}
	case command.List:
		list()
	case command.Remove:
		if len(args) > 1 {
			err = remove(args[1])
		} else {
			fmt.Println(constants.Help)
		}
	case command.Search:
		if len(args) > 1 {
			err = search(args[1], operatingSystem, arch, config)
		} else {
			err = searchDefault(operatingSystem, arch, config)
		}
	case command.Temp:
		if len(args) > 1 {
			err = temp(args[1], operatingSystem)
		} else {
			fmt.Println(constants.Help)
		}
	case command.Use:
		if len(args) > 1 {
			err = use(args[1], operatingSystem)
		} else {
			fmt.Println(constants.Help)
		}
	default:
		fmt.Println(constants.Help)
	}

	if err != nil {
		log.Fatalln(err)
	}
}
