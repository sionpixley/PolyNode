package internal

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func IsKnownCommand(command string) bool {
	return convertToCommand(command) != _NA_COMM
}

func LoadPolyNodeConfig() PolyNodeConfig {
	if _, err := os.Stat(polynHomeDir + pathSeparator + ".polynrc"); os.IsNotExist(err) {
		// Default config
		return PolyNodeConfig{NodeMirror: _DEFAULT_NODE_MIRROR}
	} else {
		content, err := os.ReadFile(polynHomeDir + pathSeparator + ".polynrc")
		if err != nil {
			// Default config
			fmt.Println(err.Error())
			return PolyNodeConfig{NodeMirror: _DEFAULT_NODE_MIRROR}
		}

		config := PolyNodeConfig{}
		err = config.UnmarshalJSON(content)
		if err != nil {
			// Default config
			fmt.Println(err.Error())
			return PolyNodeConfig{NodeMirror: _DEFAULT_NODE_MIRROR}
		}
		return config
	}
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
