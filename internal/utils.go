package internal

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func ConvertToArchitecture(archStr string) Architecture {
	switch archStr {
	case "amd64":
		return c_X64
	case "arm64":
		return c_ARM64
	default:
		return c_NA_ARCH
	}
}

func ConvertToOperatingSystem(osStr string) OperatingSystem {
	switch osStr {
	case "darwin":
		return c_MAC
	case "linux":
		return c_LINUX
	case "windows":
		return c_WIN
	default:
		return c_NA_OS
	}
}

func IsKnownCommand(command string) bool {
	return convertToCommand(command) != c_NA_COMM
}

func PrintHelp(operatingSystem OperatingSystem) {
	help := c_DESCRIPTION + "\n\n" + c_USAGE + "\n\n" + c_RUNTIMES + "\n\n" + c_COMMANDS
	fmt.Println(help)

	if operatingSystem == c_LINUX {
		fmt.Println()
	}
}

func convertToCommand(commandStr string) command {
	switch commandStr {
	case "add":
		return c_ADD
	case "current":
		return c_CURRENT
	case "install":
		return c_INSTALL
	case "ls":
		fallthrough
	case "list":
		return c_LIST
	case "rm":
		fallthrough
	case "remove":
		return c_REMOVE
	case "use":
		return c_USE
	default:
		return c_NA_COMM
	}
}

func deleteFileIfExists(filePath string) error {
	var err error
	if doesFileExist(filePath) {
		err = os.Remove(filePath)
	}
	return err
}

func doesFileExist(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}

func extractFile(source string, destination string, operatingSystem OperatingSystem, arch Architecture) error {
	command, err := get7ZipCmdLocation(operatingSystem, arch)
	if err != nil {
		return err
	}

	if strings.Contains(source, ".tar.") {
		_, err = exec.Command(command, "x", source, "-o"+destination).Output()
		if err != nil {
			return err
		}
		_, err = exec.Command(command, "x", source[:len(source)-3], "-o"+destination).Output()
	} else {
		_, err = exec.Command(command, "x", source, "-o"+destination).Output()
	}

	return err
}

func get7ZipCmdLocation(operatingSystem OperatingSystem, arch Architecture) (string, error) {
	command := ""
	switch operatingSystem {
	case c_LINUX:
		if arch == c_ARM64 {
			command = "./emb/7z/linux/arm64/7zzs"
		} else if arch == c_X64 {
			command = "./emb/7z/linux/x64/7zzs"
		}
	case c_MAC:
		command = "./emb/7z/mac/7zz"
	case c_WIN:
		command = "./emb/7z/win/7za"
	default:
		return "", errors.New(c_UNSUPPORTED_OS)
	}

	return command, nil
}

func printError(err error) {
	fmt.Println(err.Error())
}
