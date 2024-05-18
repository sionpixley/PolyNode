package internal

import (
	"fmt"
	"os/exec"

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
		return constants.NA_COMM
	}
}

func ConvertToOs(osStr string) models.Os {
	switch osStr {
	case "darwin":
		return constants.MAC
	case "linux":
		return constants.LINUX
	case "windows":
		return constants.WIN
	default:
		return constants.NA_OS
	}
}

func IsKnownCommand(command string) bool {
	return ConvertToCommand(command) != constants.NA_COMM
}

func PrintError(err error) {
	fmt.Println(err.Error())
}

func PrintHelp() {
	help := constants.DESCRIPTION + "\n\n" + constants.USAGE + "\n\n" + constants.RUNTIMES + "\n\n" + constants.COMMANDS
	fmt.Println(help)
}

func UnzipFile(source string, destination string) error {
	output, err := exec.Command("./emb/7za.exe", "x", source, "-o"+destination).Output()
	if err != nil {
		return err
	}
	fmt.Print(string(output))
	return nil
}
