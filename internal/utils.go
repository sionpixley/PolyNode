package internal

import (
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

func convertToSemanticVersion(version string) string {
	if version[0] == 'v' {
		return version
	} else {
		return "v" + version
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

func isValidVersionFormat(version string) bool {
	version = version[1:]

	parts := strings.Split(version, ".")
	if len(parts) != 3 {
		return false
	}

	validChars := map[rune]struct{}{
		'0': {},
		'1': {},
		'2': {},
		'3': {},
		'4': {},
		'5': {},
		'6': {},
		'7': {},
		'8': {},
		'9': {},
	}
	for _, part := range parts {
		for _, char := range part {
			if _, exists := validChars[char]; exists {
				// char exists in the validChars map, so it's a valid character.
				continue
			} else {
				return false
			}
		}
	}

	return true
}
