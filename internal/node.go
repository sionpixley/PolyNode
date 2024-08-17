package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
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
		err = listDownloadedNodes()
	case c_REMOVE:
		if len(args) > 1 {
			err = removeNode(args[1])
		} else {
			PrintHelp()
		}
	case c_SEARCH:
		err = searchAvailableNodeVersions()
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
	fmt.Printf("Downloading %s...", fileName)

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

	err = os.MkdirAll(polynHomeDir+pathSeparator+"node", os.ModePerm)
	if err != nil {
		return err
	}

	filePath := polynHomeDir + pathSeparator + "node" + pathSeparator + fileName
	err = os.RemoveAll(filePath)
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

	folderPath := polynHomeDir + pathSeparator + "node" + pathSeparator + version
	err = os.RemoveAll(folderPath)
	if err != nil {
		return err
	}

	fmt.Println("Done.")

	fmt.Printf("Extracting %s...", fileName)
	err = extractFile(filePath, folderPath, operatingSystem)
	if err != nil {
		return err
	}

	err = os.RemoveAll(filePath)
	if err != nil {
		return err
	}

	fmt.Println("Done.")
	fmt.Printf("Adding Node.js %s...Done.\n", version)
	return nil
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

func listDownloadedNodes() error {
	dir, err := os.ReadDir(polynHomeDir + pathSeparator + "node")
	if err != nil {
		return err
	}

	current := ""
	output, err := exec.Command("node", "-v").Output()
	if err != nil {
		// Do nothing. This just means that there isn't a current version set.
	} else {
		current = strings.TrimSpace(string(output))
	}

	for _, item := range dir {
		if item.IsDir() && current == item.Name() {
			fmt.Printf("Node.js - %s (current)\n", item.Name())
		} else if item.IsDir() {
			fmt.Printf("Node.js - %s\n", item.Name())
		}
	}

	return nil
}

func printAvailableNodeVersions(nodeVersions []NodeVersion) {
	majorVersions := map[string]bool{}
	stableVersions := []string{}
	ltsVersions := []string{}

	for _, nodeVersion := range nodeVersions {
		if len(majorVersions) == 6 {
			break
		}

		majorVersion := strings.Split(nodeVersion.Version, ".")[0]
		_, exists := majorVersions[majorVersion]
		if exists {
			continue
		} else {
			majorVersions[majorVersion] = true
			if nodeVersion.Lts != "false" {
				ltsVersions = append(ltsVersions, nodeVersion.Version)
			} else {
				stableVersions = append(stableVersions, nodeVersion.Version)
			}
		}
	}

	output := `----------------
| Stable | LTS |
----------------`

	for i := 0; i < len(stableVersions); i += 1 {
		output += "| " + stableVersions[i] + " | " + ltsVersions[i] + " |\n----------------"
	}

	fmt.Println(output)
}

func printCurrentNode() {
	output, err := exec.Command("node", "-v").Output()
	if err != nil {
		fmt.Println("There aren't any Node.js versions set as the current version.")
	} else {
		fmt.Printf("Node.js - %s", string(output))
	}
}

func removeNode(version string) error {
	version, err := convertToSemanticVersion(version)
	if err != nil {
		return err
	}

	folderName := polynHomeDir + pathSeparator + "node" + pathSeparator + version
	err = os.RemoveAll(folderName)
	if err != nil {
		return err
	}

	fmt.Printf("Deleted Node.js %s.\n", version)
	return nil
}

func searchAvailableNodeVersions() error {
	url := "https://nodejs.org/dist/index.json"

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

	nodeVersions := []NodeVersion{}
	err = json.NewDecoder(response.Body).Decode(&nodeVersions)
	if err != nil {
		return err
	}

	printAvailableNodeVersions(nodeVersions)
	return nil
}

func useNode(version string) error {
	version, err := convertToSemanticVersion(version)
	if err != nil {
		return err
	}

	fmt.Printf("Switching to Node.js %s...", version)

	err = os.RemoveAll(polynHomeDir + pathSeparator + "nodejs")
	if err != nil {
		return err
	}

	err = os.Symlink(polynHomeDir+pathSeparator+"node"+pathSeparator+version, polynHomeDir+pathSeparator+"nodejs")
	if err != nil {
		return err
	}

	fmt.Println("Done.")
	return nil
}
