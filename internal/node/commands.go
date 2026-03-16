package node

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/sionpixley/PolyNode/internal"
	"github.com/sionpixley/PolyNode/internal/constants"
	"github.com/sionpixley/PolyNode/internal/constants/opsys"
	"github.com/sionpixley/PolyNode/internal/models"
	"github.com/sionpixley/PolyNode/internal/utilities"
)

func add(version string, operatingSystem models.OperatingSystem, arch models.Architecture, config *models.PolyNodeConfig) error {
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

	client := &http.Client{Timeout: time.Duration(config.TimeoutInSeconds) * time.Second}
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

func configGet(config *models.PolyNodeConfig, configField string) {
	switch {
	case strings.EqualFold(configField, "autoUpdate"):
		fmt.Println(config.AutoUpdate)
	case strings.EqualFold(configField, "nodeMirror"):
		fmt.Println(config.NodeMirror)
	case strings.EqualFold(configField, "timeoutInSeconds"):
		fmt.Println(config.TimeoutInSeconds)
	default:
		err := fmt.Errorf(constants.InvalidConfigFieldError, configField)
		utilities.LogUserError(err)
	}
}

func configGetAll(config *models.PolyNodeConfig) {
	s := `{
  "autoUpdate": %t,
  "nodeMirror": "%s",
  "timeoutInSeconds": %d
}
`
	fmt.Printf(s, config.AutoUpdate, config.NodeMirror, config.TimeoutInSeconds)
}

func configSet(config *models.PolyNodeConfig, configField string, value string) error {
	switch {
	case strings.EqualFold(configField, "autoUpdate"):
		v, err := strconv.ParseBool(value)
		if err != nil {
			err = fmt.Errorf("invalid value: '%s' is not a valid bool value", value)
			utilities.LogUserError(err)
		}

		config.AutoUpdate = v
		return config.Save()
	case strings.EqualFold(configField, "nodeMirror"):
		config.NodeMirror = value
		return config.Save()
	case strings.EqualFold(configField, "timeoutInSeconds"):
		v, e := strconv.Atoi(value)
		if e != nil {
			e = fmt.Errorf("invalid value: '%s' is not a valid int value", value)
			utilities.LogUserError(e)
		}

		config.TimeoutInSeconds = v
		return config.Save()
	default:
		e2 := fmt.Errorf(constants.InvalidConfigFieldError, configField)
		utilities.LogUserError(e2)
		return nil
	}
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
		} else if err != nil {
			return nil
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

func install(version string, operatingSystem models.OperatingSystem, arch models.Architecture, config *models.PolyNodeConfig) error {
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

	var c string
	output, err := exec.Command("node", "-v").Output()
	if err == nil {
		c = strings.TrimSpace(string(output))
	}

	for _, item := range dir {
		if item.IsDir() && c == item.Name() {
			fmt.Printf("Node.js - %s (current)\n", item.Name())
		} else if item.IsDir() {
			fmt.Printf("Node.js - %s\n", item.Name())
		}
	}
}

func migrate(from string, to string, operatingSystem models.OperatingSystem, arch models.Architecture, config *models.PolyNodeConfig) error {
	var err error
	from, err = convertPrefixToVersionLocalDesc(from)
	// We don't want to do anything when the error's value is 'skip'.
	// If the error is 'skip' then that means the node directory doesn't exist.
	// We don't treat it like an error in that case.
	// This is unique to this function (asc and desc versions) throughout the system (so far at least). 2025-02-25
	if err != nil && err.Error() != "skip" {
		return err
	} else if err != nil {
		return nil
	}

	to, err = convertPrefixToVersionDown(to, operatingSystem, arch, config)
	if err != nil {
		return err
	}

	if from == to {
		majorVersion := strings.Split(from, ".")[0]
		fmt.Printf("%s is already the most recent %s\n", from, majorVersion)
		return nil
	}

	err = add(to, operatingSystem, arch, config)
	if err != nil {
		return err
	}

	err = def(from, operatingSystem)
	if err != nil {
		return err
	}

	data, err := exec.Command("npm", "ls", "-g", "--depth=0", "--json").Output()
	if err != nil {
		return err
	}

	var npmList models.NPMList
	err = json.Unmarshal(data, &npmList)
	if err != nil {
		return err
	}

	exclusions := map[string]struct{}{
		"corepack": {},
		"npm":      {},
	}

	dependencies := []string{"install", "-g"}
	for name, dependency := range npmList.Dependencies {
		if _, exists := exclusions[name]; !exists {
			dependencies = append(dependencies, name+"@"+dependency.Version)
		}
	}

	err = def(to, operatingSystem)
	if err != nil {
		return err
	}

	fmt.Print("migrating global npm packages (this might take a while)...")

	if len(dependencies) > 2 {
		data, err = exec.Command("npm", dependencies...).Output()
		if err != nil {
			return err
		}
	}

	fmt.Println("done")
	fmt.Print(string(data))
	return nil
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
		} else if err != nil {
			return nil
		}
	}

	fmt.Printf("removing Node.js %s...", version)

	folderName := internal.PolynHomeDir + internal.PathSeparator + "node" + internal.PathSeparator + version
	err = os.RemoveAll(folderName)
	if err != nil {
		return err
	}

	var c string
	output, err := exec.Command("node", "-v").Output()
	if err == nil {
		c = strings.TrimSpace(string(output))
	}

	if c == version {
		folderName = internal.PolynHomeDir + internal.PathSeparator + "nodejs"
		err = os.RemoveAll(folderName)
		if err != nil {
			return err
		}
	}

	fmt.Println("done")
	return nil
}

func search(prefix string, operatingSystem models.OperatingSystem, arch models.Architecture, config *models.PolyNodeConfig) error {
	prefix = utilities.ConvertToSemanticVersion(prefix)

	allVersions, err := getAllNodeVersionsForOSAndArch(operatingSystem, arch, config)
	if err != nil {
		return err
	}

	var builder strings.Builder
	builder.WriteString("Node.js\n--------------------------")
	for _, nodeVersion := range allVersions {
		if nodeVersion.LTS && strings.HasPrefix(nodeVersion.Version, prefix) {
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

func searchDefault(operatingSystem models.OperatingSystem, arch models.Architecture, config *models.PolyNodeConfig) error {
	nodeVersions, err := getAllNodeVersionsForOSAndArch(operatingSystem, arch, config)
	if err != nil {
		return err
	}

	maxEntries := 7

	majorVersions := map[string]struct{}{}
	var stableVersions []string
	var ltsVersions []string

	for _, nodeVersion := range nodeVersions {
		if len(stableVersions) == maxEntries && len(ltsVersions) == maxEntries {
			break
		}

		majorVersion := strings.Split(nodeVersion.Version, ".")[0]
		if _, exists := majorVersions[majorVersion]; !exists {
			majorVersions[majorVersion] = struct{}{}
			if nodeVersion.LTS && len(ltsVersions) < maxEntries {
				ltsVersions = append(ltsVersions, nodeVersion.Version)
			} else if !nodeVersion.LTS && len(stableVersions) < maxEntries {
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
		} else if err != nil {
			return nil
		}
	}

	if operatingSystem == opsys.Windows {
		if ranInCmd, err := runningInCmd(); err != nil {
			return err
		} else if ranInCmd {
			fmt.Printf("set PATH=%s\\node\\%s;%s\n", internal.PolynHomeDir, version, "%PATH%")
		} else {
			fmt.Printf("$env:Path = \"%s\\node\\%s;\" + $env:Path\n", internal.PolynHomeDir, version)
		}
	} else {
		fmt.Printf("export PATH=%s", internal.PolynHomeDir+"/node/"+version+"/bin:$PATH")
	}

	return nil
}
