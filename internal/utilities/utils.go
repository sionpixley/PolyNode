package utilities

import (
	"os"
	"os/exec"
	"strings"

	"github.com/sionpixley/PolyNode/internal/constants"
	"github.com/sionpixley/PolyNode/internal/models"
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
	case "search":
		return constants.SEARCH
	case "use":
		return constants.USE
	default:
		return constants.NA_COMM
	}
}

func ConvertToSemanticVersion(version string) string {
	if version[0] == 'v' {
		return version
	} else {
		return "v" + version
	}
}

func ExtractFile(source string, destination string) error {
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

func IsKnownCommand(command string) bool {
	return ConvertToCommand(command) != constants.NA_COMM
}

func IsValidVersionFormat(version string) bool {
	if version[0] == 'v' {
		version = version[1:]
	}

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
			if _, exists := validChars[char]; !exists {
				return false
			}
		}
	}

	return true
}
