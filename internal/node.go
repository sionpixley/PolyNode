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
func HandleNode(args []string, operatingSystem OperatingSystem, arch Architecture, config PolyNodeConfig) {
	if len(args) == 0 {
		fmt.Println(HELP)
		return
	}

	var err error
	command := convertToCommand(args[0])
	switch command {
	case _ADD:
		if len(args) > 1 {
			err = addNode(convertKeywordToNodeVersionStr(args[1], config), operatingSystem, arch, config)
		} else {
			fmt.Println(HELP)
		}
	case _CURRENT:
		printCurrentNode()
	case _INSTALL:
		if len(args) > 1 {
			err = installNode(convertKeywordToNodeVersionStr(args[1], config), operatingSystem, arch, config)
		} else {
			fmt.Println(HELP)
		}
	case _LIST:
		listDownloadedNodes()
	case _REMOVE:
		if len(args) > 1 {
			err = removeNode(args[1])
		} else {
			fmt.Println(HELP)
		}
	case _SEARCH:
		if len(args) > 1 {
			err = searchForSpecificNodeVersion(args[1], config)
		} else {
			err = searchAvailableNodeVersions(config)
		}
	case _USE:
		if len(args) > 1 {
			err = useNode(args[1], operatingSystem)
		} else {
			fmt.Println(HELP)
		}
	default:
		fmt.Println(HELP)
	}

	if err != nil {
		fmt.Println(err.Error())
	}
}

func addNode(version string, operatingSystem OperatingSystem, arch Architecture, config PolyNodeConfig) error {
	if !isValidVersionFormat(version) {
		return errors.New(_INVALID_VERSION_FORMAT_ERROR)
	}

	version = convertToSemanticVersion(version)

	archiveName, err := getNodeTargetArchiveName(operatingSystem, arch)
	if err != nil {
		return err
	}

	fileName := "node-" + version + "-" + archiveName
	fmt.Printf("Downloading %s...", fileName)

	url := config.NodeMirror + "/" + version + "/" + fileName

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
	// Calling file.Close() explicitly instead of with defer to prevent lock errors.
	file.Close()

	folderPath := polynHomeDir + pathSeparator + "node" + pathSeparator + version
	err = os.RemoveAll(folderPath)
	if err != nil {
		return err
	}

	fmt.Println("Done.")

	fmt.Printf("Extracting %s...", fileName)
	err = extractFile(filePath, folderPath)
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

func convertKeywordToNodeVersionStr(keyword string, config PolyNodeConfig) string {
	if keyword == "lts" {
		nodeVersions, err := getAllNodeVersions(config)
		if err != nil {
			return keyword
		}

		for _, nodeVersion := range nodeVersions {
			if nodeVersion.Lts {
				return nodeVersion.Version
			}
		}
		return keyword
	} else if keyword == "latest" {
		nodeVersions, err := getAllNodeVersions(config)
		if err != nil {
			return keyword
		}

		return nodeVersions[0].Version
	} else {
		return keyword
	}
}

func getAllNodeVersions(config PolyNodeConfig) ([]NodeVersion, error) {
	url := config.NodeMirror + "/index.json"

	client := new(http.Client)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var nodeVersions []NodeVersion
	err = json.NewDecoder(response.Body).Decode(&nodeVersions)
	return nodeVersions, err
}

func getNodeTargetArchiveName(operatingSystem OperatingSystem, arch Architecture) (string, error) {
	archiveName := ""
	switch operatingSystem {
	case LINUX:
		if arch == ARM64 {
			archiveName = "linux-arm64.tar.xz"
		} else if arch == X64 {
			archiveName = "linux-x64.tar.xz"
		}
	case MAC:
		if arch == ARM64 {
			archiveName = "darwin-arm64.tar.gz"
		} else if arch == X64 {
			archiveName = "darwin-x64.tar.gz"
		}
	case WINDOWS:
		if arch == ARM64 {
			archiveName = "win-arm64.zip"
		} else if arch == X64 {
			archiveName = "win-x64.zip"
		}
	default:
		return "", errors.New(UNSUPPORTED_OS_ERROR)
	}

	return archiveName, nil
}

func installNode(version string, operatingSystem OperatingSystem, arch Architecture, config PolyNodeConfig) error {
	err := addNode(version, operatingSystem, arch, config)
	if err != nil {
		return err
	}

	return useNode(version, operatingSystem)
}

func listDownloadedNodes() {
	dir, err := os.ReadDir(polynHomeDir + "/node")
	if err != nil {
		// This means that the node folder doesn't exist. So, there are no Node.js versions downloaded.
		fmt.Println("There are no Node.js versions downloaded.")
		fmt.Println("To download a Node.js version, use the 'add' or 'install' command.")
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
	if !isValidVersionFormat(version) {
		return errors.New(_INVALID_VERSION_FORMAT_ERROR)
	}

	version = convertToSemanticVersion(version)

	fmt.Printf("Removing Node.js %s...", version)

	folderName := polynHomeDir + "/node/" + version
	err := os.RemoveAll(folderName)
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

	if current == version {
		err = os.RemoveAll(folderName)
		if err != nil {
			return err
		}
	}

	fmt.Println("Done.")
	return nil
}

func searchAvailableNodeVersions(config PolyNodeConfig) error {
	nodeVersions, err := getAllNodeVersions(config)
	if err != nil {
		return err
	}

	maxEntries := 7

	majorVersions := map[string]struct{}{}
	stableVersions := []string{}
	ltsVersions := []string{}

	for _, nodeVersion := range nodeVersions {
		if len(stableVersions) == maxEntries && len(ltsVersions) == maxEntries {
			break
		}

		majorVersion := strings.Split(nodeVersion.Version, ".")[0]
		if _, exists := majorVersions[majorVersion]; !exists {
			majorVersions[majorVersion] = struct{}{}
			if nodeVersion.Lts && len(ltsVersions) < maxEntries {
				ltsVersions = append(ltsVersions, nodeVersion.Version)
			} else if !nodeVersion.Lts && len(stableVersions) < maxEntries {
				stableVersions = append(stableVersions, nodeVersion.Version)
			}
		}
	}

	output := "\nLatest stable versions of Node.js\n---------------------------------"
	for _, stableVersion := range stableVersions {
		output += "\n" + stableVersion
	}

	output += "\n\nLatest LTS versions of Node.js\n---------------------------------"
	for _, ltsVersion := range ltsVersions {
		output += "\n" + ltsVersion
	}

	fmt.Println(output)
	return nil
}

func searchForSpecificNodeVersion(prefix string, config PolyNodeConfig) error {
	prefix = convertToSemanticVersion(prefix)

	allVersions, err := getAllNodeVersions(config)
	if err != nil {
		return err
	}

	output := "\nNode.js\n--------------------------"
	for _, nodeVersion := range allVersions {
		if strings.HasPrefix(nodeVersion.Version, prefix) {
			output += "\n" + nodeVersion.Version
		}
	}

	fmt.Println(output)
	return nil
}

func useNode(version string, operatingSystem OperatingSystem) error {
	if !isValidVersionFormat(version) {
		return errors.New(_INVALID_VERSION_FORMAT_ERROR)
	}

	version = convertToSemanticVersion(version)

	fmt.Printf("Switching to Node.js %s...", version)

	err := os.RemoveAll(polynHomeDir + pathSeparator + "nodejs")
	if err != nil {
		return err
	}

	if operatingSystem == WINDOWS {
		err = exec.Command("cmd", "/c", "mklink", "/j", polynHomeDir+"\\nodejs", polynHomeDir+"\\node\\"+version).Run()
		if err != nil {
			return err
		}
	} else {
		err = os.Symlink(polynHomeDir+"/node/"+version, polynHomeDir+"/nodejs")
		if err != nil {
			return err
		}
	}

	fmt.Println("Done.")
	return nil
}
