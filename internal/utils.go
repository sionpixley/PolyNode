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
	help := "\n" + c_DESCRIPTION + "\n\n" + c_USAGE + "\n\n" + c_COMMANDS
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

	if operatingSystem == c_WIN {
		err = exec.Command(command, "x", source, "-o"+destination).Run()
	} else {
		err = exec.Command(command, "x", source).Run()
		if err != nil {
			return err
		}

		tarball := strings.Split(source, "/")[2]
		tarball = tarball[:len(tarball)-3]

		err = exec.Command(command, "x", tarball).Run()
		if err != nil {
			return err
		}

		err = deleteFileIfExists(source)
		if err != nil {
			return err
		}

		err = deleteFileIfExists(tarball)
		if err != nil {
			return err
		}

		tarball = tarball[:len(tarball)-4]
		err = exec.Command("mv", tarball, destination).Run()
	}

	return err
}

func get7ZipCmdLocation(operatingSystem OperatingSystem, arch Architecture) (string, error) {
	switch operatingSystem {
	case c_LINUX:
		if arch == c_ARM64 {
			return polynHomeDir + "/emb/7z/linux/arm64/7zzs", nil
		} else if arch == c_X64 {
			return polynHomeDir + "/emb/7z/linux/x64/7zzs", nil
		} else {
			return "", errors.New(c_UNSUPPORTED_ARCH)
		}
	case c_MAC:
		return polynHomeDir + "/emb/7z/mac/7zz", nil
	case c_WIN:
		return polynHomeDir + "\\emb\\7z\\win\\7za", nil
	default:
		return "", errors.New(c_UNSUPPORTED_OS)
	}
}

func printError(err error) {
	fmt.Println(err.Error())
}
