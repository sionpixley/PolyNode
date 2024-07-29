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

func IsSupportedArchitecture(arch Architecture) bool {
	return arch != c_NA_ARCH
}

func IsSupportedOperatingSystem(operatingSystem OperatingSystem) bool {
	return operatingSystem != c_NA_OS
}

func PrintHelp() {
	help := c_DESCRIPTION + "\n\n" + c_USAGE + "\n\n" + c_COMMANDS
	fmt.Println(help)
}

// Windows automatically adds a new line at the end of stdout.
// Linux and macOS need an extra line printed to make the output look better.
func PrintOptionalLine(operatingSystem OperatingSystem) {
	if operatingSystem != c_WIN {
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

func convertToSemanticVersion(version string) (string, error) {
	if len(version) < 6 {
		return "", errors.New("invalid version format")
	} else if version[0] == 'v' {
		return version, nil
	} else {
		return "v" + version, nil
	}
}

func extractFile(source string, destination string, operatingSystem OperatingSystem) error {
	var err error
	if operatingSystem == c_WIN {
		// err = exec.Command(command, "x", source, "-o"+polynHomeDir).Run()
		// if err != nil {
		// 	return err
		// }

		parts := strings.Split(source, "\\")
		folderName := parts[len(parts)-1]
		folderName = polynHomeDir + "\\" + folderName[:len(folderName)-3]

		err = exec.Command("xcopy", "/s", "/i", folderName+"\\", destination+"\\").Run()
		if err != nil {
			return err
		}

		err = os.RemoveAll(folderName)
	} else {
		err = exec.Command("tar", "-xf", source, "-C", destination).Run()
		if err != nil {
			return err
		}

		err = os.RemoveAll(source)
		if err != nil {
			return err
		}
	}

	return err
}

func printError(err error) {
	fmt.Println(err.Error())
}
