package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
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

func autoUpdate(operatingSystem models.OperatingSystem, architecture models.Architecture, httpWrapper models.HTTPWrapper, ioWrapper models.IOWrapper, osWrapper models.OSWrapper) error {
	now := time.Now().UTC()
	lastUpdated := getLastUpdate(osWrapper)
	if now.Sub(lastUpdated).Hours() >= 720 {
		err := updatePolyNode(operatingSystem, architecture, httpWrapper, ioWrapper, osWrapper)
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

func downloadPolyNodeFile(filename string, httpWrapper models.HTTPWrapper, ioWrapper models.IOWrapper, osWrapper models.OSWrapper) error {
	fmt.Print("downloading the latest release of PolyNode...")

	client := httpWrapper.NewClient()
	request, err := httpWrapper.NewRequest(http.MethodGet, "https://github.com/sionpixley/PolyNode/releases/latest/download/"+filename, nil)
	if err != nil {
		return err
	}

	response, err := httpWrapper.Do(client, request)
	if err != nil {
		return err
	}
	defer func() { _ = response.Body.Close() }()

	filename = internal.PolynHomeDir + internal.PathSeparator + filename
	err = osWrapper.RemoveAll(filename)
	if err != nil {
		return err
	}

	file, err := osWrapper.Create(filename)
	if err != nil {
		return err
	}
	defer func() { _ = file.Close() }()

	_, err = ioWrapper.Copy(file, response.Body)
	if err != nil {
		return err
	}

	fmt.Println("done")
	return nil
}

func execute(args []string, operatingSystem models.OperatingSystem, architecture models.Architecture, config *models.PolyNodeConfig, httpWrapper models.HTTPWrapper, ioWrapper models.IOWrapper, osWrapper models.OSWrapper) {
	var err error
	if args[0] == "update" {
		err = updatePolyNode(operatingSystem, architecture, httpWrapper, ioWrapper, osWrapper)
		if err != nil {
			log.Fatalf("polyn: %v\n", err)
		}
	} else if utilities.KnownCommand(args[0]) {
		node.Handle(args, operatingSystem, architecture, config, httpWrapper)
	} else {
		err = fmt.Errorf(constants.UnknownCommandError, args[0])
		utilities.LogUserError(err)
	}

	if config.AutoUpdate {
		err = autoUpdate(operatingSystem, architecture, httpWrapper, ioWrapper, osWrapper)
		if err != nil {
			log.Fatalf("polyn: %v\n", err)
		}
	}
}

func getLastUpdate(osWrapper models.OSWrapper) time.Time {
	updateFilePath := internal.PolynHomeDir + internal.PathSeparator + "last-update.txt"
	if _, err := osWrapper.Stat(updateFilePath); osWrapper.IsNotExist(err) {
		return time.Now().UTC().AddDate(0, 0, -30)
	} else if err != nil {
		return time.Now().UTC().AddDate(0, 0, -30)
	}

	content, err := osWrapper.ReadFile(updateFilePath)
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

func parseCLIArgs(osWrapper models.OSWrapper) []string {
	flag.Usage = func() {
		w := flag.CommandLine.Output()
		_, _ = fmt.Fprintln(w, constants.Help)
	}

	var version bool
	flag.BoolVarP(&version, "version", "v", false, "print the version and exit")

	flag.Parse()

	if version {
		fmt.Println(constants.Version)
		osWrapper.Exit(0)
	}

	if flag.NArg() < 1 {
		flag.CommandLine.SetOutput(osWrapper.Stdout())
		flag.Usage()
		osWrapper.Exit(0)
	}

	args := make([]string, flag.NArg())
	for i := range flag.NArg() {
		args[i] = strings.ToLower(flag.Arg(i))
	}

	return args
}

func runUpdateScript(operatingSystem models.OperatingSystem, osWrapper models.OSWrapper) error {
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
timeout /t 1 /nobreak > nul
if exist %LOCALAPPDATA%\Programs\PolyNode\update-temp\ (
  del %LOCALAPPDATA%\Programs\PolyNode\update-temp /s /f /q > nul
  rmdir %LOCALAPPDATA%\Programs\PolyNode\update-temp /s /q
)
(goto) 2>nul & del "%~f0"`

		err := osWrapper.WriteFile(batchfilePath, []byte(updateBatchfile), 0744)
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

		err = osWrapper.RemoveAll(updateTemp)
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

func updatePolyNode(operatingSystem models.OperatingSystem, architecture models.Architecture, httpWrapper models.HTTPWrapper, ioWrapper models.IOWrapper, osWrapper models.OSWrapper) error {
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

	err := downloadPolyNodeFile(filename, httpWrapper, ioWrapper, osWrapper)
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

	return runUpdateScript(operatingSystem, osWrapper)
}
