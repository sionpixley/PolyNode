package node

import (
	"errors"
	"fmt"
	"log"

	"github.com/sionpixley/PolyNode/internal/constants"
	"github.com/sionpixley/PolyNode/internal/constants/command"
	"github.com/sionpixley/PolyNode/internal/constants/subcomm"
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
	case command.Config:
		if len(args) > 1 {
			scomm := utilities.ConvertToSubcommand(args[1])
			if scomm == subcomm.Get {
				if len(args) > 2 {
					configGet(args[2])
				} else {
					err = fmt.Errorf(constants.MissingConfigFieldError, "config get")
					utilities.LogUserError(err)
				}
			} else if scomm == subcomm.Set {
				if len(args) > 3 {
					err = configSet(args[2], args[3])
				} else if len(args) > 2 {
					err = fmt.Errorf("missing argument: 'config set %s' requires a new value", args[2])
					utilities.LogUserError(err)
				} else {
					err = fmt.Errorf(constants.MissingConfigFieldError, "config set")
					utilities.LogUserError(err)
				}
			}
		} else {
			err = errors.New("missing argument: 'config' command requires a subcommand 'get' or 'set'")
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
