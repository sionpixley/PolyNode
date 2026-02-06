package node

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/sionpixley/PolyNode/internal"
	"github.com/sionpixley/PolyNode/internal/constants"
	"github.com/sionpixley/PolyNode/internal/constants/opsys"
	"github.com/sionpixley/PolyNode/internal/models"
	"github.com/sionpixley/PolyNode/internal/utilities"
)

func add(version string, operatingSystem models.OperatingSystem, arch models.Architecture, config models.PolyNodeConfig) error {
	var err error

	if utilities.ValidVersionFormat(version) {
		version = utilities.ConvertToSemanticVersion(version)
	} else {
		version, err = convertPrefixToVersionDown(version, operatingSystem, arch, config)
		if err != nil {
			return err
		}
	}

	archiveName, err := getArchiveName(operatingSystem, arch)
	if err != nil {
		return err
	}

	fileName := "node-" + version + "-" + archiveName
	fmt.Printf("downloading %s...", fileName)

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
	defer func() { _ = response.Body.Close() }()

	nodePath := internal.PolynHomeDir + internal.PathSeparator + "node"
	err = os.MkdirAll(nodePath, os.ModePerm)
	if err != nil {
		return err
	}

	filePath := nodePath + internal.PathSeparator + fileName
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
		_ = file.Close()
		return err
	}
	// Calling file.Close() explicitly instead of with defer to prevent lock errors.
	_ = file.Close()

	folderPath := nodePath + internal.PathSeparator + version
	err = os.RemoveAll(folderPath)
	if err != nil {
		return err
	}

	fmt.Println("done")

	fmt.Printf("extracting %s...", fileName)
	err = utilities.ExtractFile(filePath, folderPath)
	if err != nil {
		return err
	}

	err = os.RemoveAll(filePath)
	if err != nil {
		return err
	}

	fmt.Println("done")
	fmt.Printf("adding Node.js %s...done\n", version)
	return nil
}

func configGet(configField string) {
	config := models.LoadPolyNodeConfig()
	if configField == "autoupdate" {
		fmt.Println(config.AutoUpdate)
	} else if configField == "nodemirror" {
		fmt.Println(config.NodeMirror)
	} else {
		err := fmt.Errorf(constants.InvalidConfigFieldError, configField)
		utilities.LogUserError(err)
	}
}

func configSet(configField string, value string) error {
	config := models.LoadPolyNodeConfig()
	if configField == "autoupdate" {
		if value == "true" {
			config.AutoUpdate = true
		} else if value == "false" {
			config.AutoUpdate = false
		} else {
			err := fmt.Errorf("invalid value: '%s' is not a valid bool value", value)
			utilities.LogUserError(err)
		}

		return config.Save()
	} else if configField == "nodemirror" {
		config.NodeMirror = value
		return config.Save()
	}

	err := fmt.Errorf(constants.InvalidConfigFieldError, configField)
	utilities.LogUserError(err)
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

func def(version string, operatingSystem models.OperatingSystem) error {
	var err error

	if utilities.ValidVersionFormat(version) {
		version = utilities.ConvertToSemanticVersion(version)
	} else {
		version, err = convertPrefixToVersionLocalDesc(version)
		// We don't want to do anything when the error's value is 'skip'.
		// If the error is 'skip' then that means the node directory doesn't exist.
		// We don't treat it like an error in that case.
		// This is unique to this function (asc and desc versions) throughout the system (so far at least). 2025-02-25
		if err != nil && err.Error() != "skip" {
			return err
		}
	}

	fmt.Printf("switching to Node.js %s...", version)

	nodejsPath := internal.PolynHomeDir + internal.PathSeparator + "nodejs"
	err = os.RemoveAll(nodejsPath)
	if err != nil {
		return err
	}

	if operatingSystem == opsys.Windows {
		err = exec.Command("cmd", "/c", "mklink", "/j", nodejsPath, internal.PolynHomeDir+"\\node\\"+version).Run()
		if err != nil {
			return err
		}
	} else {
		err = os.Symlink(internal.PolynHomeDir+"/node/"+version, nodejsPath)
		if err != nil {
			return err
		}
	}

	fmt.Println("done")
	return nil
}

func install(version string, operatingSystem models.OperatingSystem, arch models.Architecture, config models.PolyNodeConfig) error {
	err := add(version, operatingSystem, arch, config)
	if err != nil {
		return err
	}

	return def(version, operatingSystem)
}

func list() {
	dir, err := os.ReadDir(internal.PolynHomeDir + internal.PathSeparator + "node")
	if err != nil {
		// This means that the node folder doesn't exist. So, there are no Node.js versions downloaded.
		fmt.Println(constants.NoDownloadedNodejsMessage)
		return
	}

	var current string
	output, err := exec.Command("node", "-v").Output()
	if err == nil {
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
	var err error

	if utilities.ValidVersionFormat(version) {
		version = utilities.ConvertToSemanticVersion(version)
	} else {
		version, err = convertPrefixToVersionLocalAsc(version)
		// We don't want to do anything when the error's value is 'skip'.
		// If the error is 'skip' then that means the node directory doesn't exist.
		// We don't treat it like an error in that case.
		// This is unique to this function (asc and desc versions) throughout the system (so far at least). 2025-02-25
		if err != nil && err.Error() != "skip" {
			return err
		}
	}

	fmt.Printf("removing Node.js %s...", version)

	folderName := internal.PolynHomeDir + internal.PathSeparator + "node" + internal.PathSeparator + version
	err = os.RemoveAll(folderName)
	if err != nil {
		return err
	}

	var current string
	output, err := exec.Command("node", "-v").Output()
	if err == nil {
		current = strings.TrimSpace(string(output))
	}

	if current == version {
		folderName := internal.PolynHomeDir + internal.PathSeparator + "nodejs"
		err = os.RemoveAll(folderName)
		if err != nil {
			return err
		}
	}

	fmt.Println("done")
	return nil
}

func search(prefix string, operatingSystem models.OperatingSystem, arch models.Architecture, config models.PolyNodeConfig) error {
	prefix = utilities.ConvertToSemanticVersion(prefix)

	allVersions, err := getAllNodeVersionsForOSAndArch(operatingSystem, arch, config)
	if err != nil {
		return err
	}

	var builder strings.Builder
	builder.WriteString("Node.js\n--------------------------")
	for _, nodeVersion := range allVersions {
		if nodeVersion.Lts && strings.HasPrefix(nodeVersion.Version, prefix) {
			builder.WriteString("\n")
			builder.WriteString(nodeVersion.Version)
			builder.WriteString(" (lts)")
		} else if strings.HasPrefix(nodeVersion.Version, prefix) {
			builder.WriteString("\n")
			builder.WriteString(nodeVersion.Version)
		}
	}

	fmt.Println(builder.String())
	return nil
}

func searchDefault(operatingSystem models.OperatingSystem, arch models.Architecture, config models.PolyNodeConfig) error {
	nodeVersions, err := getAllNodeVersionsForOSAndArch(operatingSystem, arch, config)
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

	output := "Latest stable versions of Node.js\n---------------------------------"
	for _, stableVersion := range stableVersions {
		output += "\n" + stableVersion
	}

	output += "\n\nLatest LTS versions of Node.js\n---------------------------------"
	for _, ltsVersion := range ltsVersions {
		output += "\n" + ltsVersion + " (lts)"
	}

	fmt.Println(output)
	return nil
}

func use(version string, operatingSystem models.OperatingSystem) error {
	var err error

	if utilities.ValidVersionFormat(version) {
		version = utilities.ConvertToSemanticVersion(version)
	} else {
		version, err = convertPrefixToVersionLocalDesc(version)
		// We don't want to do anything when the error's value is 'skip'.
		// If the error is 'skip' then that means the node directory doesn't exist.
		// We don't treat it like an error in that case.
		// This is unique to this function (asc and desc versions) throughout the system (so far at least). 2025-02-25
		if err != nil && err.Error() != "skip" {
			return err
		}
	}

	if operatingSystem == opsys.Windows {
		nodeVersionPath := internal.PolynHomeDir + "\\node\\" + version
		fmt.Println("If using Command Prompt, run this command:")
		fmt.Printf("\n  set PATH=%s%s\n", nodeVersionPath, ";%PATH%")
		fmt.Println("\nIf using PowerShell, run this command:")
		fmt.Printf("\n  $env:Path = \"%s%s\n", nodeVersionPath, ";\" + $env:Path")
	} else {
		fmt.Printf("export PATH=%s", internal.PolynHomeDir+"/node/"+version+"/bin:$PATH")
	}

	return nil
}
