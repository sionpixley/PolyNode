package internal

import (
	"fmt"

	"github.com/sionpixley/polyn/internal/constants"
	"github.com/sionpixley/polyn/internal/models"
)

func ConvertToCommand(commandStr string) models.Command {
	switch commandStr {
	case "add":
		return constants.ADD
	case "current":
		return constants.CURRENT
	case "install":
		return constants.INSTALL
	case "ls":
		fallthrough
	case "list":
		return constants.LIST
	case "rm":
		fallthrough
	case "remove":
		return constants.REMOVE
	case "use":
		return constants.USE
	default:
		return constants.NA
	}
}

func IsKnownCommand(command string) bool {
	return ConvertToCommand(command) != constants.NA
}

func PrintError(err error) {
	fmt.Println(err.Error())
}

func PrintHelp() {
	help := constants.DESCRIPTION + "\n\n" + constants.USAGE + "\n\n" + constants.ENVS + "\n\n" + constants.COMMANDS
	fmt.Println(help)
}
