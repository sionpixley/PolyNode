package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"

	"github.com/sionpixley/PolyNode/internal"
	"github.com/sionpixley/PolyNode/internal/constants"
	"github.com/sionpixley/PolyNode/internal/models"
	"github.com/sionpixley/PolyNode/internal/utilities"
)

func convertToArchitecture(archStr string) models.Architecture {
	switch archStr {
	case "amd64":
		return constants.X64
	case "arm64":
		return constants.ARM64
	case "ppc64":
		return constants.PPC64
	case "ppc64le":
		return constants.PPC64LE
	case "s390x":
		return constants.S390X
	default:
		return constants.NA_ARCH
	}
}

func convertToOperatingSystem(osStr string) models.OperatingSystem {
	switch osStr {
	case "aix":
		return constants.AIX
	case "darwin":
		return constants.MAC
	case "linux":
		return constants.LINUX
	case "windows":
		return constants.WINDOWS
	default:
		return constants.NA_OS
	}
}

func downloadPolyNodeFile(filename string) error {
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

func isSupportedArchitecture(arch models.Architecture) bool {
	return arch != constants.NA_ARCH
}

func isSupportedOperatingSystem(operatingSystem models.OperatingSystem) bool {
	return operatingSystem != constants.NA_OS
}

func runUpgradeScript(operatingSystem models.OperatingSystem) error {
	fmt.Print("Running upgrade script...")
	if operatingSystem == constants.WINDOWS {
		batchfilePath := internal.PolynHomeDir + internal.PathSeparator + "polyn-upgrade-temp.cmd"
		upgradeBatchfile := `@echo off

timeout /t 1 /nobreak > nul
cd %LOCALAPPDATA%\Programs\PolyNode\upgrade-temp
.\setup
cd %LOCALAPPDATA%
del %LOCALAPPDATA%\Programs\PolyNode\upgrade-temp /s /f /q > nul
rmdir %LOCALAPPDATA%\Programs\PolyNode\upgrade-temp /s /q
(goto) 2>nul & del "%~f0"`

		err := os.WriteFile(batchfilePath, []byte(upgradeBatchfile), 0700)
		if err != nil {
			return err
		}

		fmt.Println("Done.")
		return exec.Command("cmd", "/c", "start", "/b", batchfilePath).Run()
	} else {
		scriptPath := internal.PolynHomeDir + internal.PathSeparator + "polyn-upgrade-temp"
		upgradeScript := `#!/bin/sh

sleep 1
cd $HOME/.PolyNode/upgrade-temp
./setup
cd $HOME
rm -rf $HOME/.PolyNode/upgrade-temp
rm $HOME/.PolyNode/polyn-upgrade-temp`

		err := os.WriteFile(scriptPath, []byte(upgradeScript), 0700)
		if err != nil {
			return err
		}

		fmt.Println("Done.")
		return exec.Command(scriptPath).Run()
	}
}

func upgradePolyNode(operatingSystem models.OperatingSystem, arch models.Architecture) error {
	var filename string
	switch operatingSystem {
	case constants.AIX:
		filename = "PolyNode-aix-ppc64.tar.gz"
	case constants.LINUX:
		switch arch {
		case constants.ARM64:
			filename = "PolyNode-linux-arm64.tar.xz"
		case constants.PPC64LE:
			filename = "PolyNode-linux-ppc64le.tar.xz"
		case constants.S390X:
			filename = "PolyNode-linux-s390x.tar.xz"
		case constants.X64:
			filename = "PolyNode-linux-x64.tar.xz"
		default:
			return errors.New(constants.UNSUPPORTED_ARCH_ERROR)
		}
	case constants.MAC:
		switch arch {
		case constants.ARM64:
			filename = "PolyNode-darwin-arm64.tar.gz"
		case constants.X64:
			filename = "PolyNode-darwin-x64.tar.gz"
		default:
			return errors.New(constants.UNSUPPORTED_ARCH_ERROR)
		}
	case constants.WINDOWS:
		switch arch {
		case constants.ARM64:
			filename = "PolyNode-win-arm64.zip"
		case constants.X64:
			filename = "PolyNode-win-x64.zip"
		default:
			return errors.New(constants.UNSUPPORTED_ARCH_ERROR)
		}
	default:
		return errors.New(constants.UNSUPPORTED_OS_ERROR)
	}

	err := downloadPolyNodeFile(filename)
	if err != nil {
		return err
	}

	fmt.Print("Extracting " + filename + "...")
	filename = internal.PolynHomeDir + internal.PathSeparator + filename
	err = utilities.ExtractFile(filename, internal.PolynHomeDir+internal.PathSeparator+"upgrade-temp")
	if err != nil {
		return err
	}
	fmt.Println("Done.")

	return runUpgradeScript(operatingSystem)
}
