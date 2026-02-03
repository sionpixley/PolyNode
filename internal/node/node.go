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
			err = fmt.Errorf(constants.MissingVersionKeywordOrPrefixError, args[0])
			utilities.LogError(err)
		}
	case command.Current:
		current()
	case command.Default:
		if len(args) > 1 {
			err = def(args[1], operatingSystem)
		} else {
			err = fmt.Errorf(constants.MissingVersionOrPrefixError, args[0])
			utilities.LogError(err)
		}
	case command.Install:
		if len(args) > 1 {
			err = install(convertKeywordToVersion(args[1], operatingSystem, arch, config), operatingSystem, arch, config)
		} else {
			err = fmt.Errorf(constants.MissingVersionKeywordOrPrefixError, args[0])
			utilities.LogError(err)
		}
	case command.List:
		list()
	case command.Remove:
		if len(args) > 1 {
			err = remove(args[1])
		} else {
			err = fmt.Errorf(constants.MissingVersionOrPrefixError, args[0])
			utilities.LogError(err)
		}
	case command.Search:
		if len(args) > 1 {
			err = search(args[1], operatingSystem, arch, config)
		} else {
			err = searchDefault(operatingSystem, arch, config)
		}
	case command.Use:
		if len(args) > 1 {
			err = use(args[1], operatingSystem)
		} else {
			err = fmt.Errorf(constants.MissingVersionOrPrefixError, args[0])
			utilities.LogError(err)
		}
	default:
		err = fmt.Errorf(constants.UnknownCommandError, args[0])
		utilities.LogError(err)
	}

	if err != nil {
		log.Fatalln(err)
	}
}
