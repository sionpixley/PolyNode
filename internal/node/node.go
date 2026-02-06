package node

import (
	"errors"
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
			utilities.LogUserError(err)
		}
	case command.ConfigGet:
		if len(args) > 1 {
			configGet(config, args[1])
		} else {
			configGetAll(config)
		}
	case command.ConfigSet:
		if len(args) > 2 {
			err = configSet(config, args[1], args[2])
		} else if len(args) > 1 {
			err = fmt.Errorf("missing argument: 'config-set %s' requires a new value", args[1])
			utilities.LogUserError(err)
		} else {
			err = errors.New("missing argument: 'config-set' command requires a config field")
			utilities.LogUserError(err)
		}
	case command.Current:
		current()
	case command.Default:
		if len(args) > 1 {
			err = def(args[1], operatingSystem)
		} else {
			err = fmt.Errorf(constants.MissingVersionOrPrefixError, args[0])
			utilities.LogUserError(err)
		}
	case command.Install:
		if len(args) > 1 {
			err = install(convertKeywordToVersion(args[1], operatingSystem, arch, config), operatingSystem, arch, config)
		} else {
			err = fmt.Errorf(constants.MissingVersionKeywordOrPrefixError, args[0])
			utilities.LogUserError(err)
		}
	case command.List:
		list()
	case command.Remove:
		if len(args) > 1 {
			err = remove(args[1])
		} else {
			err = fmt.Errorf(constants.MissingVersionOrPrefixError, args[0])
			utilities.LogUserError(err)
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
			utilities.LogUserError(err)
		}
	default:
		err = fmt.Errorf(constants.UnknownCommandError, args[0])
		utilities.LogUserError(err)
	}

	if err != nil {
		log.Fatalln(err)
	}
}
