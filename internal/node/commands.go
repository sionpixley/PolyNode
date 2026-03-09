package node

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/sionpixley/PolyNode/internal"
	"github.com/sionpixley/PolyNode/internal/constants"
	"github.com/sionpixley/PolyNode/internal/constants/opsys"
	"github.com/sionpixley/PolyNode/internal/models"
	"github.com/sionpixley/PolyNode/internal/utilities"
)

func add(version string, operatingSystem models.OperatingSystem, arch models.Architecture, config *models.PolyNodeConfig, httpWrapper models.HTTPWrapper, gzipWrapper models.GzipWrapper, ioWrapper models.IOWrapper, osWrapper models.OSWrapper, tarWrapper models.TarWrapper) error {
	var err error

	if utilities.ValidVersionFormat(version) {
		version = utilities.ConvertToSemanticVersion(version)
	} else {
		version, err = convertPrefixToVersionDown(version, operatingSystem, arch, config, httpWrapper)
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

	client := httpWrapper.NewClient()
	request, err := httpWrapper.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	response, err := httpWrapper.Do(client, request)
	if err != nil {
		return err
	}
	defer func() { _ = response.Body.Close() }()

	nodePath := internal.PolynHomeDir + internal.PathSeparator + "node"
	err = osWrapper.MkdirAll(nodePath, os.ModePerm)
	if err != nil {
		return err
	}

	filePath := nodePath + internal.PathSeparator + fileName
	err = osWrapper.RemoveAll(filePath)
	if err != nil {
		return err
	}

	file, err := osWrapper.Create(filePath)
	if err != nil {
		return err
	}

	_, err = ioWrapper.Copy(file, response.Body)
	if err != nil {
		_ = file.Close()
		return err
	}
	// Calling file.Close() explicitly instead of with defer to prevent lock errors.
	_ = file.Close()

	folderPath := nodePath + internal.PathSeparator + version
	err = osWrapper.RemoveAll(folderPath)
	if err != nil {
		return err
	}

	fmt.Println("done")

	fmt.Printf("extracting %s...", fileName)
	err = utilities.ExtractFile(filePath, folderPath, gzipWrapper, ioWrapper, osWrapper, tarWrapper)
	if err != nil {
		return err
	}

	err = osWrapper.RemoveAll(filePath)
	if err != nil {
		return err
	}

	fmt.Println("done")
	fmt.Printf("adding Node.js %s...done\n", version)
	return nil
}

func configGet(config *models.PolyNodeConfig, configField string, osWrapper models.OSWrapper) {
	if configField == "autoupdate" {
		fmt.Println(config.AutoUpdate)
	} else if configField == "nodemirror" {
		fmt.Println(config.NodeMirror)
	} else {
		err := fmt.Errorf(constants.InvalidConfigFieldError, configField)
		utilities.LogUserError(err, osWrapper)
	}
}

func configGetAll(config *models.PolyNodeConfig) {
	s := `{
  "autoUpdate": %t,
  "nodeMirror": "%s"
}
`
	fmt.Printf(s, config.AutoUpdate, config.NodeMirror)
}

func configSet(config *models.PolyNodeConfig, configField string, value string, osWrapper models.OSWrapper) error {
	if configField == "autoupdate" {
		v, err := strconv.ParseBool(value)
		if err != nil {
			err = fmt.Errorf("invalid value: '%s' is not a valid bool value", value)
			utilities.LogUserError(err, osWrapper)
		}

		config.AutoUpdate = v
		return config.Save(osWrapper)
	} else if configField == "nodemirror" {
		config.NodeMirror = value
		return config.Save(osWrapper)
	}

	e := fmt.Errorf(constants.InvalidConfigFieldError, configField)
	utilities.LogUserError(e, osWrapper)
	return nil
}

func current(execWrapper models.ExecWrapper) {
	output, err := execWrapper.Output(exec.Command("node", "-v"))
	if err != nil {
		fmt.Println("There aren't any Node.js versions set as the current version.")
	} else {
		fmt.Printf("Node.js - %s", string(output))
	}
}

func def(version string, operatingSystem models.OperatingSystem, execWrapper models.ExecWrapper, osWrapper models.OSWrapper) error {
	var err error

	if utilities.ValidVersionFormat(version) {
		version = utilities.ConvertToSemanticVersion(version)
	} else {
		version, err = convertPrefixToVersionLocalDesc(version, osWrapper)
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
	err = osWrapper.RemoveAll(nodejsPath)
	if err != nil {
		return err
	}

	if operatingSystem == opsys.Windows {
		err = execWrapper.Run(exec.Command("cmd", "/c", "mklink", "/j", nodejsPath, internal.PolynHomeDir+"\\node\\"+version))
		if err != nil {
			return err
		}
	} else {
		err = osWrapper.Symlink(internal.PolynHomeDir+"/node/"+version, nodejsPath)
		if err != nil {
			return err
		}
	}

	fmt.Println("done")
	return nil
}

func install(version string, operatingSystem models.OperatingSystem, arch models.Architecture, config *models.PolyNodeConfig, execWrapper models.ExecWrapper, gzipWrapper models.GzipWrapper, httpWrapper models.HTTPWrapper, ioWrapper models.IOWrapper, osWrapper models.OSWrapper, tarWrapper models.TarWrapper) error {
	err := add(version, operatingSystem, arch, config, httpWrapper, gzipWrapper, ioWrapper, osWrapper, tarWrapper)
	if err != nil {
		return err
	}

	return def(version, operatingSystem, execWrapper, osWrapper)
}

func list(execWrapper models.ExecWrapper, osWrapper models.OSWrapper) {
	dir, err := osWrapper.ReadDir(internal.PolynHomeDir + internal.PathSeparator + "node")
	if err != nil {
		// This means that the node folder doesn't exist. So, there are no Node.js versions downloaded.
		fmt.Println(constants.NoDownloadedNodejsMessage)
		return
	}

	var current string
	output, err := execWrapper.Output(exec.Command("node", "-v"))
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

func remove(version string, execWrapper models.ExecWrapper, osWrapper models.OSWrapper) error {
	var err error

	if utilities.ValidVersionFormat(version) {
		version = utilities.ConvertToSemanticVersion(version)
	} else {
		version, err = convertPrefixToVersionLocalAsc(version, osWrapper)
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
	err = osWrapper.RemoveAll(folderName)
	if err != nil {
		return err
	}

	var current string
	output, err := execWrapper.Output(exec.Command("node", "-v"))
	if err == nil {
		current = strings.TrimSpace(string(output))
	}

	if current == version {
		folderName := internal.PolynHomeDir + internal.PathSeparator + "nodejs"
		err = osWrapper.RemoveAll(folderName)
		if err != nil {
			return err
		}
	}

	fmt.Println("done")
	return nil
}

func search(prefix string, operatingSystem models.OperatingSystem, arch models.Architecture, config *models.PolyNodeConfig, httpWrapper models.HTTPWrapper) error {
	prefix = utilities.ConvertToSemanticVersion(prefix)

	allVersions, err := getAllNodeVersionsForOSAndArch(operatingSystem, arch, config, httpWrapper)
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

func searchDefault(operatingSystem models.OperatingSystem, arch models.Architecture, config *models.PolyNodeConfig, httpWrapper models.HTTPWrapper) error {
	nodeVersions, err := getAllNodeVersionsForOSAndArch(operatingSystem, arch, config, httpWrapper)
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

func use(version string, operatingSystem models.OperatingSystem, execWrapper models.ExecWrapper, osWrapper models.OSWrapper) error {
	var err error

	if utilities.ValidVersionFormat(version) {
		version = utilities.ConvertToSemanticVersion(version)
	} else {
		version, err = convertPrefixToVersionLocalDesc(version, osWrapper)
		// We don't want to do anything when the error's value is 'skip'.
		// If the error is 'skip' then that means the node directory doesn't exist.
		// We don't treat it like an error in that case.
		// This is unique to this function (asc and desc versions) throughout the system (so far at least). 2025-02-25
		if err != nil && err.Error() != "skip" {
			return err
		}
	}

	if operatingSystem == opsys.Windows {
		if ranInCmd, err := runningInCmd(execWrapper); err != nil {
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
