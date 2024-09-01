package internal

import (
	"errors"
	"os"
	"os/exec"
	"strings"
)

func ConvertToArchitecture(archStr string) Architecture {
	switch archStr {
	case "amd64":
		return _X64
	case "arm64":
		return _ARM64
	default:
		return _NA_ARCH
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
	return convertToCommand(command) != _NA_COMM
}

func IsSupportedArchitecture(arch Architecture) bool {
	return arch != _NA_ARCH
}

func IsSupportedOperatingSystem(operatingSystem OperatingSystem) bool {
	return operatingSystem != NA_OS
}

func convertToCommand(commandStr string) command {
	switch commandStr {
	case "add":
		return _ADD
	case "current":
		return _CURRENT
	case "install":
		return _INSTALL
	case "ls":
		fallthrough
	case "list":
		return _LIST
	case "rm":
		fallthrough
	case "remove":
		return _REMOVE
	case "search":
		return _SEARCH
	case "use":
		return _USE
	default:
		return _NA_COMM
	}
}

func convertToSemanticVersion(version string) (string, error) {
	parts := strings.Split(version, ".")
	if len(parts) != 3 {
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
