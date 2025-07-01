package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/sionpixley/PolyNode/internal"
	"github.com/sionpixley/PolyNode/internal/constants"
	"github.com/sionpixley/PolyNode/internal/models"
	"github.com/sionpixley/PolyNode/internal/utilities"
)

const isoDateTimeFormat = "2006-01-02T15:04:05.000Z07:00"

func autoUpdate(operatingSystem models.OperatingSystem, arch models.Architecture) error {
	now := time.Now().UTC()
	lastUpdated := getLastAutoUpdate()
	if now.Sub(lastUpdated).Hours() > 719 {
		err := updatePolyNode(operatingSystem, arch)
		if err != nil {
			return err
		}
	}

	return nil
}

func convertToArchitecture(archStr string) models.Architecture {
	switch archStr {
	case "amd64":
		return constants.X64
	case "arm64":
		return constants.Arm64
	case "ppc64":
		return constants.Ppc64
	case "ppc64le":
		return constants.Ppc64Le
	case "s390x":
		return constants.S390x
	default:
		return constants.OtherArch
	}
}

func convertToOperatingSystem(osStr string) models.OperatingSystem {
	switch osStr {
	case "aix":
		return constants.Aix
	case "darwin":
		return constants.Mac
	case "linux":
		return constants.Linux
	case "windows":
		return constants.Windows
	default:
		return constants.OtherOS
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

func getLastAutoUpdate() time.Time {
	if _, err := os.Stat(internal.PolynHomeDir + internal.PathSeparator + "lastAutoUpdate.txt"); os.IsNotExist(err) {
		return time.Now().UTC().AddDate(0, 0, -30)
	} else if err != nil {
		return time.Now().UTC().AddDate(0, 0, -30)
	} else {
		content, err := os.ReadFile(internal.PolynHomeDir + internal.PathSeparator + "lastAutoUpdate.txt")
		if err != nil {
			return time.Now().UTC().AddDate(0, 0, -30)
		}

		timeStr := strings.TrimSpace(string(content))
		t, err := time.Parse(isoDateTimeFormat, timeStr)
		if err != nil {
			return time.Now().UTC().AddDate(0, 0, -30)
		}

		return t
	}
}

func isSupportedArchitecture(arch models.Architecture) bool {
	return arch != constants.OtherArch
}

func isSupportedOperatingSystem(operatingSystem models.OperatingSystem) bool {
	return operatingSystem != constants.OtherOS
}

func runUpdateScript(operatingSystem models.OperatingSystem) error {
	fmt.Print("Running update...")

	if operatingSystem == constants.Windows {
		batchfilePath := internal.PolynHomeDir + "\\polyn-upgrade-temp.cmd"
		upgradeBatchfile := `@echo off
timeout /t 1 /nobreak > nul
cd %LOCALAPPDATA%\Programs\PolyNode\upgrade-temp
.\setup
cd %LOCALAPPDATA%
del %LOCALAPPDATA%\Programs\PolyNode\upgrade-temp /s /f /q > nul
rmdir %LOCALAPPDATA%\Programs\PolyNode\upgrade-temp /s /q
if exist %LOCALAPPDATA%\Programs\PolyNode\upgrade-temp (
  del %LOCALAPPDATA%\Programs\PolyNode\upgrade-temp /s /f /q > nul
  rmdir %LOCALAPPDATA%\Programs\PolyNode\upgrade-temp /s /q
)
(goto) 2>nul & del "%~f0"`

		err := os.WriteFile(batchfilePath, []byte(upgradeBatchfile), 0744)
		if err != nil {
			return err
		}

		err = exec.Command("cmd", "/c", "start", "/b", batchfilePath).Run()
		if err != nil {
			return err
		}
	} else {
		err := exec.Command(internal.PolynHomeDir + internal.PathSeparator + "update-temp" + internal.PathSeparator + "setup").Run()
		if err != nil {
			return err
		}

		err = os.RemoveAll(internal.PolynHomeDir + internal.PathSeparator + "update-temp")
		if err != nil {
			return err
		}
	}

	fmt.Println("Done.")
	return nil
}

func updatePolyNode(operatingSystem models.OperatingSystem, arch models.Architecture) error {
	var filename string
	switch operatingSystem {
	case constants.Aix:
		filename = "PolyNode-aix-ppc64.tar.gz"
	case constants.Linux:
		switch arch {
		case constants.Arm64:
			filename = "PolyNode-linux-arm64.tar.gz"
		case constants.Ppc64Le:
			filename = "PolyNode-linux-ppc64le.tar.gz"
		case constants.S390x:
			filename = "PolyNode-linux-s390x.tar.gz"
		case constants.X64:
			filename = "PolyNode-linux-x64.tar.gz"
		default:
			return errors.New(constants.UnsupportedArchError)
		}
	case constants.Mac:
		switch arch {
		case constants.Arm64:
			filename = "PolyNode-darwin-arm64.tar.gz"
		case constants.X64:
			filename = "PolyNode-darwin-x64.tar.gz"
		default:
			return errors.New(constants.UnsupportedArchError)
		}
	case constants.Windows:
		switch arch {
		case constants.Arm64:
			filename = "PolyNode-win-arm64.zip"
		case constants.X64:
			filename = "PolyNode-win-x64.zip"
		default:
			return errors.New(constants.UnsupportedArchError)
		}
	default:
		return errors.New(constants.UnsupportedOSError)
	}

	err := downloadPolyNodeFile(filename)
	if err != nil {
		return err
	}

	fmt.Print("Extracting " + filename + "...")
	filename = internal.PolynHomeDir + internal.PathSeparator + filename
	err = utilities.ExtractFile(filename, internal.PolynHomeDir+internal.PathSeparator+"update-temp")
	if err != nil {
		return err
	}
	fmt.Println("Done.")

	return runUpdateScript(operatingSystem)
}
