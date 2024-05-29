package internal

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
)

// Main function for Node.js actions.
func HandleNode(args []string, operatingSystem OperatingSystem, arch Architecture) {
	if len(args) == 0 {
		PrintHelp()
		return
	}

	var err error
	command := convertToCommand(args[0])
	switch command {
	case c_ADD:
		if len(args) > 1 {
			err = addNode(args[1], operatingSystem, arch)
		} else {
			PrintHelp()
		}
	case c_CURRENT:
		printCurrentNode()
	case c_INSTALL:
		if len(args) > 1 {
			err = installNode(args[1], operatingSystem, arch)
		} else {
			PrintHelp()
		}
	case c_LIST:
		listDownloadedNodes()
	case c_REMOVE:
		if len(args) > 1 {
			err = removeNode(args[1])
		} else {
			PrintHelp()
		}
	case c_USE:
		if len(args) > 1 {
			err = useNode(args[1])
		} else {
			PrintHelp()
		}
	default:
		PrintHelp()
	}

	if err != nil {
		printError(err)
	} else {
		fmt.Println("\nDone.")
	}
}

func addNode(version string, operatingSystem OperatingSystem, arch Architecture) error {
	version, err := convertToSemanticVersion(version)
	if err != nil {
		return err
	}

	archiveName, err := getNodeTargetArchiveName(operatingSystem, arch)
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

	err = os.MkdirAll("node", os.ModePerm)
	if err != nil {
		return err
	}

	filePath := "./node/" + fileName
	err = deleteFileIfExists(filePath)
	if err != nil {
		return err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}
	// Calling file.Close() explicitly instead of with defer because the 7-Zip command was getting a lock error on the zip file.
	file.Close()

	folderPath := "./node/" + version
	err = os.RemoveAll(folderPath)
	if err != nil {
		return err
	}

	fmt.Println("Extracting " + fileName + "...")
	err = extractFile(filePath, folderPath, operatingSystem, arch)
	if err != nil {
		return err
	}

	err = deleteFileIfExists(filePath)
	return err
}

func getNodeTargetArchiveName(operatingSystem OperatingSystem, arch Architecture) (string, error) {
	archiveName := ""
	switch operatingSystem {
	case c_LINUX:
		if arch == c_ARM64 {
			archiveName = "linux-arm64.tar.xz"
		} else if arch == c_X64 {
			archiveName = "linux-x64.tar.xz"
		}
	case c_MAC:
		if arch == c_ARM64 {
			archiveName = "darwin-arm64.tar.gz"
		} else if arch == c_X64 {
			archiveName = "darwin-x64.tar.gz"
		}
	case c_WIN:
		archiveName = "win-x64.zip"
	default:
		return "", errors.New(c_UNSUPPORTED_OS)
	}

	return archiveName, nil
}

func installNode(version string, operatingSystem OperatingSystem, arch Architecture) error {
	err := addNode(version, operatingSystem, arch)
	if err != nil {
		return err
	}

	err = useNode(version)
	return err
}

func listDownloadedNodes() {

}

func printCurrentNode() {
	output, err := exec.Command("node", "-v").Output()
	if err != nil {
		fmt.Println("There aren't any Node.js versions set as the current version.")
	} else {
		fmt.Print(string(output))
	}
}

func removeNode(version string) error {
	return nil
}

func useNode(version string) error {
	return nil
}
