package internal

import (
	"errors"
	"fmt"
	"os/exec"

	"github.com/sionpixley/polyn/internal/constants"
	"github.com/sionpixley/polyn/internal/models"
)

func ConvertToArchitecture(archStr string) models.Architecture {
	switch archStr {
	case "amd64":
		return constants.X64
	case "arm64":
		return constants.ARM64
	default:
		return constants.NA_ARCH
	}
}

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

func UnzipFile(source string, destination string, operatingSystem models.Os, arch models.Architecture) error {
	command, err := get7zipCmdLocation(operatingSystem, arch)
	if err != nil {
		return err
	}

	output, err := exec.Command(command, "x", source, "-o"+destination).Output()
	if err != nil {
		return err
	}
	fmt.Print(string(output))
	return nil
}

func get7zipCmdLocation(operatingSystem models.Os, arch models.Architecture) (string, error) {
	command := ""
	switch operatingSystem {
	case constants.LINUX:
		if arch == constants.ARM64 {
			command = "./emb/7z/linux/7zzs-arm64"
		} else if arch == constants.X64 {
			command = "./emb/7z/linux/7zzs-x64"
		}
	case constants.MAC:
		command = "./emb/7z/mac/7zz"
	case constants.WIN:
		command = "./emb/7z/win/7za.exe"
	default:
		return "", errors.New(constants.UNSUPPORTED_OS)
	}

	return command, nil
}
