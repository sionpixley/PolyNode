package node

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"

	"github.com/sionpixley/polyn/internal"
	"github.com/sionpixley/polyn/internal/constants"
	"github.com/sionpixley/polyn/internal/models"
)

// Main function for Node.js actions.
// The args parameter should not include the optional runtime.
func Handle(args []string, operatingSystem models.Os, arch models.Architecture) {
	if args == nil || len(args) == 0 {
		internal.PrintHelp()
		return
	}

	var err error
	command := internal.ConvertToCommand(args[0])
	switch command {
	case constants.ADD:
		if len(args) > 1 {
			err = add(args[1], operatingSystem, arch)
		} else {
			internal.PrintHelp()
		}
	case constants.CURRENT:
		printCurrent()
	case constants.INSTALL:
		if len(args) > 1 {
			err = install(args[1], operatingSystem, arch)
		} else {
			internal.PrintHelp()
		}
	case constants.LIST:
		listDownloaded()
	case constants.REMOVE:
		if len(args) > 1 {
			err = remove(args[1])
		} else {
			internal.PrintHelp()
		}
	case constants.USE:
		if len(args) > 1 {
			err = use(args[1])
		} else {
			internal.PrintHelp()
		}
	default:
		internal.PrintHelp()
	}

	if err != nil {
		internal.PrintError(err)
	}
}

func add(version string, operatingSystem models.Os, arch models.Architecture) error {
	archiveName, err := getTargetArchiveName(operatingSystem, arch)
	if err != nil {
		return err
	}

	fileName := "node-" + version + "-" + archiveName
	fmt.Println("\nDownloading " + fileName + "...")

	url := "https://nodejs.org/dist/" + version + "/" + fileName

	client := new(http.Client)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	err = os.Mkdir("node", os.ModePerm)
	if err != nil {
		return err
	}

	filePath := "./node/" + fileName
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}
	// Calling file.Close() explicitly instead of with defer because the 7-zip command was getting a lock error on the zip file.
	file.Close()

	err = internal.UnzipFile(filePath, "./node/"+version, operatingSystem, arch)
	return err
}

func getTargetArchiveName(operatingSystem models.Os, arch models.Architecture) (string, error) {
	archiveName := ""
	switch operatingSystem {
	case constants.LINUX:
		if arch == constants.ARM64 {
			archiveName = "linux-arm64.tar.xz"
		} else if arch == constants.X64 {
			archiveName = "linux-x64.tar.xz"
		}
	case constants.MAC:
		if arch == constants.ARM64 {
			archiveName = "darwin-arm64.tar.gz"
		} else if arch == constants.X64 {
			archiveName = "darwin-x64.tar.gz"
		}
	case constants.WIN:
		archiveName = "win-x64.zip"
	default:
		return "", errors.New(constants.UNSUPPORTED_OS)
	}

	return archiveName, nil
}

func install(version string, operatingSystem models.Os, arch models.Architecture) error {
	err := add(version, operatingSystem, arch)
	if err != nil {
		return err
	}

	err = use(version)
	return err
}

func listDownloaded() {

}

func printCurrent() {
	output, err := exec.Command("node", "-v").Output()
	if err != nil {
		fmt.Println("There aren't any Node.js versions set as the current version.")
	} else {
		fmt.Print(string(output))
	}
}

func remove(version string) error {
	return nil
}

func use(version string) error {
	return nil
}
