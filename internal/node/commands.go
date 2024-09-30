package node

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
	"github.com/sionpixley/PolyNode/internal/utilities"
)

func add(version string, operatingSystem models.OperatingSystem, arch models.Architecture, config models.PolyNodeConfig) error {
	if !utilities.IsValidVersionFormat(version) {
		return errors.New(constants.INVALID_VERSION_FORMAT_ERROR)
	}

	version = utilities.ConvertToSemanticVersion(version)

	archiveName, err := getArchiveName(operatingSystem, arch)
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

	err = os.MkdirAll(internal.PolynHomeDir+internal.PathSeparator+"node", os.ModePerm)
	if err != nil {
		return err
	}

	filePath := internal.PolynHomeDir + internal.PathSeparator + "node" + internal.PathSeparator + fileName
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

	folderPath := internal.PolynHomeDir + internal.PathSeparator + "node" + internal.PathSeparator + version
	err = os.RemoveAll(folderPath)
	if err != nil {
		return err
	}

	fmt.Println("Done.")

	fmt.Printf("Extracting %s...", fileName)
	err = utilities.ExtractFile(filePath, folderPath)
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

func current() {
	output, err := exec.Command("node", "-v").Output()
	if err != nil {
		fmt.Println("There aren't any Node.js versions set as the current version.")
	} else {
		fmt.Printf("Node.js - %s", string(output))
	}
}

func install(version string, operatingSystem models.OperatingSystem, arch models.Architecture, config models.PolyNodeConfig) error {
	err := add(version, operatingSystem, arch, config)
	if err != nil {
		return err
	}

	return use(version, operatingSystem)
}

func list() {
	dir, err := os.ReadDir(internal.PolynHomeDir + "/node")
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

func remove(version string) error {
	if !utilities.IsValidVersionFormat(version) {
		return errors.New(constants.INVALID_VERSION_FORMAT_ERROR)
	}

	version = utilities.ConvertToSemanticVersion(version)

	fmt.Printf("Removing Node.js %s...", version)

	folderName := internal.PolynHomeDir + "/node/" + version
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

func search(prefix string, config models.PolyNodeConfig) error {
	prefix = utilities.ConvertToSemanticVersion(prefix)

	allVersions, err := getAllNodeVersions(config)
	if err != nil {
		return err
	}

	output := "\nNode.js\n--------------------------"
	for _, nodeVersion := range allVersions {
		if nodeVersion.Lts && strings.HasPrefix(nodeVersion.Version, prefix) {
			output += "\n" + nodeVersion.Version + " (lts)"
		} else if strings.HasPrefix(nodeVersion.Version, prefix) {
			output += "\n" + nodeVersion.Version
		}
	}

	fmt.Println(output)
	return nil
}

func searchDefault(config models.PolyNodeConfig) error {
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

func use(version string, operatingSystem models.OperatingSystem) error {
	if !utilities.IsValidVersionFormat(version) {
		return errors.New(constants.INVALID_VERSION_FORMAT_ERROR)
	}

	version = utilities.ConvertToSemanticVersion(version)

	fmt.Printf("Switching to Node.js %s...", version)

	err := os.RemoveAll(internal.PolynHomeDir + internal.PathSeparator + "nodejs")
	if err != nil {
		return err
	}

	if operatingSystem == constants.WINDOWS {
		err = exec.Command("cmd", "/c", "mklink", "/j", internal.PolynHomeDir+"\\nodejs", internal.PolynHomeDir+"\\node\\"+version).Run()
		if err != nil {
			return err
		}
	} else {
		err = os.Symlink(internal.PolynHomeDir+"/node/"+version, internal.PolynHomeDir+"/nodejs")
		if err != nil {
			return err
		}
	}

	fmt.Println("Done.")
	return nil
}
