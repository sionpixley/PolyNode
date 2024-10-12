package utilities

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/sionpixley/PolyNode/internal"
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

func DownloadPolyNodeFile(filename string) error {
	fmt.Print("Downloading the latest release of PolyNode...")

	client := new(http.Client)
	request, err := http.NewRequest(http.MethodGet, "https://github.com/sionpixley/PolyNode/releases/latest/download/"+filename, nil)
	if err != nil {
		return err
	}

	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	filename = internal.PolynHomeDir + internal.PathSeparator + filename
	err = os.RemoveAll(filename)
	if err != nil {
		return err
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	_, err = io.Copy(file, response.Body)
	if err != nil {
		file.Close()
		return err
	}
	// Closing the file explicitly to avoid lock errors.
	file.Close()

	fmt.Println("Done.")
	return nil
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

func UpgradePolyNode(operatingSystem models.OperatingSystem, arch models.Architecture) error {
	var guiInstalled bool
	if _, err := os.Stat(internal.PolynHomeDir + internal.PathSeparator + "gui"); os.IsNotExist(err) {
		guiInstalled = false
	} else if err != nil {
		guiInstalled = false
	} else {
		guiInstalled = true
	}

	var filename string
	switch operatingSystem {
	case constants.AIX:
		if guiInstalled {
			filename = "PolyNode-GUI-aix-ppc64.tar.gz"
		} else {
			filename = "PolyNode-aix-ppc64.tar.gz"
		}
	case constants.LINUX:
		switch arch {
		case constants.ARM64:
			if guiInstalled {
				filename = "PolyNode-GUI-linux-arm64.tar.xz"
			} else {
				filename = "PolyNode-linux-arm64.tar.xz"
			}
		case constants.X64:
			if guiInstalled {
				filename = "PolyNode-GUI-linux-x64.tar.xz"
			} else {
				filename = "PolyNode-linux-x64.tar.xz"
			}
		default:
			return errors.New(constants.UNSUPPORTED_ARCH_ERROR)
		}
	case constants.MAC:
		switch arch {
		case constants.ARM64:
			if guiInstalled {
				filename = "PolyNode-GUI-darwin-arm64.tar.gz"
			} else {
				filename = "PolyNode-darwin-arm64.tar.gz"
			}
		case constants.X64:
			if guiInstalled {
				filename = "PolyNode-GUI-darwin-x64.tar.gz"
			} else {
				filename = "PolyNode-darwin-x64.tar.gz"
			}
		default:
			return errors.New(constants.UNSUPPORTED_ARCH_ERROR)
		}
	case constants.WINDOWS:
		switch arch {
		case constants.ARM64:
			if guiInstalled {
				filename = "PolyNode-GUI-win-arm64.zip"
			} else {
				filename = "PolyNode-win-arm64.zip"
			}
		case constants.X64:
			if guiInstalled {
				filename = "PolyNode-GUI-win-x64.zip"
			} else {
				filename = "PolyNode-win-x64.zip"
			}
		default:
			return errors.New(constants.UNSUPPORTED_ARCH_ERROR)
		}
	default:
		return errors.New(constants.UNSUPPORTED_OS_ERROR)
	}

	err := DownloadPolyNodeFile(filename)
	if err != nil {
		return err
	}

	filename = internal.PolynHomeDir + internal.PathSeparator + filename
	err = ExtractFile(filename, internal.PolynHomeDir+internal.PathSeparator+"upgrade-temp")
	if err != nil {
		return err
	}

	return runUpgradeScript(operatingSystem)
}

func runUpgradeScript(operatingSystem models.OperatingSystem) error {
	if operatingSystem == constants.WINDOWS {
		batchfilePath := internal.PolynHomeDir + internal.PathSeparator + "polyn-upgrade-temp.cmd"
		upgradeBatchfile := `@echo off

timeout /t 1 /nobreak > nul
%LOCALAPPDATA%\Programs\PolyNode\upgrade-temp\setup
del %LOCALAPPDATA%\Programs\PolyNode\upgrade-temp /s /f /q > nul
rmdir %LOCALAPPDATA%\Programs\PolyNode\upgrade-temp /s /q
(goto) 2>nul & del "%~f0%"`

		err := os.WriteFile(batchfilePath, []byte(upgradeBatchfile), 0700)
		if err != nil {
			return err
		}

		return exec.Command("cmd", "/c", "start", "/b", batchfilePath).Run()
	} else {
		scriptPath := internal.PolynHomeDir + internal.PathSeparator + "polyn-upgrade-temp"
		upgradeScript := `#!/bin/sh

sleep 1
$HOME/.PolyNode/upgrade-temp/setup
rm -rf $HOME/.PolyNode/upgrade-temp
rm $HOME/.PolyNode/polyn-upgrade-temp`

		err := os.WriteFile(scriptPath, []byte(upgradeScript), 0700)
		if err != nil {
			return err
		}

		return exec.Command(scriptPath).Run()
	}
}
