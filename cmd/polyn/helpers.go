package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/sionpixley/PolyNode/internal"
	"github.com/sionpixley/PolyNode/internal/constants"
	"github.com/sionpixley/PolyNode/internal/constants/arch"
	"github.com/sionpixley/PolyNode/internal/constants/opsys"
	"github.com/sionpixley/PolyNode/internal/models"
	"github.com/sionpixley/PolyNode/internal/node"
	"github.com/sionpixley/PolyNode/internal/utilities"
	flag "github.com/spf13/pflag"
)

const isoDateTimeFormat = "2006-01-02T15:04:05.000Z07:00"

func autoUpdate(operatingSystem models.OperatingSystem, architecture models.Architecture) error {
	now := time.Now().UTC()
	lastUpdated := getLastUpdate()
	if now.Sub(lastUpdated).Hours() >= 720 {
		err := updatePolyNode(operatingSystem, architecture)
		if err != nil {
			return err
		}
	}

	return nil
}

func checkArchitecture() models.Architecture {
	architecture := convertToArchitecture(runtime.GOARCH)
	if !supportedArchitecture(architecture) {
		log.Fatalln(constants.UnsupportedArchError)
	}
	return architecture
}

func checkOS() models.OperatingSystem {
	operatingSystem := convertToOperatingSystem(runtime.GOOS)
	if !supportedOS(operatingSystem) {
		log.Fatalln(constants.UnsupportedOSError)
	}
	return operatingSystem
}

func convertToArchitecture(archStr string) models.Architecture {
	switch archStr {
	case "amd64":
		return arch.X64
	case "arm64":
		return arch.ARM64
	case "ppc64":
		return arch.PPC64
	case "ppc64le":
		return arch.PPC64LE
	case "s390x":
		return arch.S390X
	default:
		return arch.Other
	}
}

func convertToOperatingSystem(osStr string) models.OperatingSystem {
	switch osStr {
	case "aix":
		return opsys.AIX
	case "darwin":
		return opsys.Mac
	case "linux":
		return opsys.Linux
	case "windows":
		return opsys.Windows
	default:
		return opsys.Other
	}
}

func downloadPolyNodeFile(filename string) error {
	fmt.Print("downloading the latest release of PolyNode...")

	client := new(http.Client)
	request, err := http.NewRequest(http.MethodGet, "https://github.com/sionpixley/PolyNode/releases/latest/download/"+filename, nil)
	if err != nil {
		return err
	}

	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer func() { _ = response.Body.Close() }()

	filename = internal.PolynHomeDir + internal.PathSeparator + filename
	err = os.RemoveAll(filename)
	if err != nil {
		return err
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer func() { _ = file.Close() }()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	fmt.Println("done")
	return nil
}

func execute(args []string, operatingSystem models.OperatingSystem, architecture models.Architecture, config models.PolyNodeConfig) {
	var err error
	switch {
	case args[0] == "version":
		fmt.Println(constants.Version)
	case args[0] == "update":
		err = updatePolyNode(operatingSystem, architecture)
		if err != nil {
			log.Fatalln(err)
		}
	case utilities.KnownCommand(args[0]):
		node.Handle(args, operatingSystem, architecture, config)
	default:
		err = fmt.Errorf(constants.UnknownCommandError, args[0])
		log.Fatalln(err)
	}

	if config.AutoUpdate {
		err = autoUpdate(operatingSystem, architecture)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func getLastUpdate() time.Time {
	updateFilePath := internal.PolynHomeDir + internal.PathSeparator + "last-update.txt"
	if _, err := os.Stat(updateFilePath); os.IsNotExist(err) {
		return time.Now().UTC().AddDate(0, 0, -30)
	} else if err != nil {
		return time.Now().UTC().AddDate(0, 0, -30)
	}

	content, err := os.ReadFile(updateFilePath)
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

func parseCLIArgs() []string {
	flag.Usage = func() {
		w := flag.CommandLine.Output()
		_, _ = fmt.Fprintln(w, constants.Help)
	}

	flag.Parse()

	if flag.NArg() < 1 {
		log.Fatalln("polyn error: no command specified")
	}

	args := make([]string, flag.NArg())
	for i := range flag.NArg() {
		args[i] = strings.ToLower(flag.Arg(i))
	}

	return args
}

func runUpdateScript(operatingSystem models.OperatingSystem) error {
	fmt.Print("running update...")

	if operatingSystem == opsys.Windows {
		batchfilePath := internal.PolynHomeDir + "\\polyn-update-temp.cmd"
		updateBatchfile := `@echo off
timeout /t 1 /nobreak > nul
cd %LOCALAPPDATA%\Programs\PolyNode\update-temp
.\setup
timeout /t 1 /nobreak > nul
cd %LOCALAPPDATA%
del %LOCALAPPDATA%\Programs\PolyNode\update-temp /s /f /q > nul
rmdir %LOCALAPPDATA%\Programs\PolyNode\update-temp /s /q
if exist %LOCALAPPDATA%\Programs\PolyNode\update-temp\ (
  del %LOCALAPPDATA%\Programs\PolyNode\update-temp /s /f /q > nul
  rmdir %LOCALAPPDATA%\Programs\PolyNode\update-temp /s /q
)
(goto) 2>nul & del "%~f0"`

		err := os.WriteFile(batchfilePath, []byte(updateBatchfile), 0744)
		if err != nil {
			return err
		}

		err = exec.Command("cmd", "/c", "start", "/b", batchfilePath).Run()
		if err != nil {
			return err
		}
	} else {
		updateTemp := internal.PolynHomeDir + internal.PathSeparator + "update-temp"
		err := exec.Command(updateTemp + internal.PathSeparator + "setup").Run()
		if err != nil {
			return err
		}

		err = os.RemoveAll(updateTemp)
		if err != nil {
			return err
		}
	}

	fmt.Println("done")
	return nil
}

func supportedArchitecture(architecture models.Architecture) bool {
	return architecture != arch.Other
}

func supportedOS(operatingSystem models.OperatingSystem) bool {
	return operatingSystem != opsys.Other
}

func updatePolyNode(operatingSystem models.OperatingSystem, architecture models.Architecture) error {
	var filename string
	switch operatingSystem {
	case opsys.AIX:
		filename = "PolyNode-aix-ppc64.tar.gz"
	case opsys.Linux:
		switch architecture {
		case arch.ARM64:
			filename = "PolyNode-linux-arm64.tar.gz"
		case arch.PPC64LE:
			filename = "PolyNode-linux-ppc64le.tar.gz"
		case arch.S390X:
			filename = "PolyNode-linux-s390x.tar.gz"
		case arch.X64:
			filename = "PolyNode-linux-x64.tar.gz"
		default:
			return errors.New(constants.UnsupportedArchError)
		}
	case opsys.Mac:
		switch architecture {
		case arch.ARM64:
			filename = "PolyNode-darwin-arm64.tar.gz"
		case arch.X64:
			filename = "PolyNode-darwin-x64.tar.gz"
		default:
			return errors.New(constants.UnsupportedArchError)
		}
	case opsys.Windows:
		switch architecture {
		case arch.ARM64:
			filename = "PolyNode-win-arm64.zip"
		case arch.X64:
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

	fmt.Printf("extracting %s...", filename)
	filename = internal.PolynHomeDir + internal.PathSeparator + filename
	err = utilities.ExtractFile(filename, internal.PolynHomeDir+internal.PathSeparator+"update-temp")
	if err != nil {
		return err
	}
	fmt.Println("done")

	return runUpdateScript(operatingSystem)
}
