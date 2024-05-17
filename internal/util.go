package internal

import (
	"fmt"

	"github.com/sionpixley/polyn/internal/constants"
)

func IsKnownCommand(command string) bool {
	return constants.ConvertToCommand(command) != constants.NA
}

func PrintError(err error) {
	fmt.Println(err.Error())
}

func PrintHelp() {
	help := constants.DESCRIPTION + "\n\n" + constants.USAGE + "\n\n" + constants.ENVS + "\n\n" + constants.COMMANDS
	fmt.Println(help)
}
