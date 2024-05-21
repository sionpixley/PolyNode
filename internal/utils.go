package internal

import (
	"errors"
	"fmt"
	"os"
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

func ConvertToOs(osStr string) models.OperatingSystem {
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

func DeleteFileIfExists(filePath string) error {
	var err error
	if doesFileExist(filePath) {
		err = os.Remove(filePath)
	}
	return err
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

// Works with more than just .zip files.
func UnzipFile(source string, destination string, operatingSystem models.OperatingSystem, arch models.Architecture) error {
	command, err := get7zipCmdLocation(operatingSystem, arch)
	if err != nil {
		return err
	}

	_, err = exec.Command(command, "x", source, "-o"+destination).Output()
	return err
}

func doesFileExist(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}

func get7zipCmdLocation(operatingSystem models.OperatingSystem, arch models.Architecture) (string, error) {
	command := ""
	switch operatingSystem {
	case constants.LINUX:
		if arch == constants.ARM64 {
			command = "./emb/7z/linux/arm/7zzs"
		} else if arch == constants.X64 {
			command = "./emb/7z/linux/x64/7zzs"
		}
	case constants.MAC:
		command = "./emb/7z/mac/7zz"
	case constants.WIN:
		command = "./emb/7z/win/7za"
	default:
		return "", errors.New(constants.UNSUPPORTED_OS)
	}

	return command, nil
}
