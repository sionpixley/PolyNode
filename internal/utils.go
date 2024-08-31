package internal

import (
	"errors"
	"os"
	"os/exec"
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
		return MAC
	case "linux":
		return LINUX
	case "windows":
		return WINDOWS
	default:
		return NA_OS
	}
}

func IsKnownCommand(command string) bool {
	return convertToCommand(command) != c_NA_COMM
}

func IsSupportedArchitecture(arch Architecture) bool {
	return arch != c_NA_ARCH
}

func IsSupportedOperatingSystem(operatingSystem OperatingSystem) bool {
	return operatingSystem != NA_OS
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
	case "search":
		return c_SEARCH
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

func extractFile(source string, destination string) error {
	err := os.RemoveAll(destination)
	if err != nil {
		return err
	}

	err = os.MkdirAll(destination, os.ModePerm)
	if err != nil {
		return err
	}

	err = exec.Command("tar", "-xf", source, "-C", destination, "--strip-components=1").Run()
	if err != nil {
		return err
	}

	return os.RemoveAll(source)
}
