package node

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"sort"
	"strings"

	"github.com/sionpixley/PolyNode/internal"
	"github.com/sionpixley/PolyNode/internal/constants"
	"github.com/sionpixley/PolyNode/internal/models"
	"github.com/sionpixley/PolyNode/internal/utilities"
)

func convertKeywordToVersion(keyword string, config models.PolyNodeConfig) string {
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

func convertOsAndArchToNodeVersionFile(operatingSystem models.OperatingSystem, arch models.Architecture) (string, error) {
	switch operatingSystem {
	case constants.AIX:
		return "aix-ppc64", nil
	case constants.LINUX:
		switch arch {
		case constants.ARM64:
			return "linux-arm64", nil
		case constants.PPC64LE:
			return "linux-ppc64le", nil
		case constants.S390X:
			return "linux-s390x", nil
		case constants.X64:
			return "linux-x64", nil
		default:
			return "", errors.New(constants.UNSUPPORTED_ARCH_ERROR)
		}
	case constants.MAC:
		if arch == constants.ARM64 {
			return "osx-arm64-tar", nil
		} else if arch == constants.X64 {
			return "osx-x64-tar", nil
		} else {
			return "", errors.New(constants.UNSUPPORTED_ARCH_ERROR)
		}
	case constants.WINDOWS:
		if arch == constants.ARM64 {
			return "win-arm64-zip", nil
		} else if arch == constants.X64 {
			return "win-x64-zip", nil
		} else {
			return "", errors.New(constants.UNSUPPORTED_ARCH_ERROR)
		}
	default:
		return "", errors.New(constants.UNSUPPORTED_OS_ERROR)
	}
}

func convertPrefixToVersionDown(prefix string, config models.PolyNodeConfig) (string, error) {
	nodeVersions, err := getAllNodeVersions(config)
	if err != nil {
		return "", err
	}

	prefix = utilities.ConvertToSemanticVersion(prefix)
	for _, nodeVersion := range nodeVersions {
		if strings.HasPrefix(nodeVersion.Version, prefix) {
			return nodeVersion.Version, nil
		}
	}

	return "", errors.New("polyn error: no Node.js versions match the prefix '" + prefix + "'")
}

func convertPrefixToVersionLocalAsc(prefix string) (string, error) {
	dir, err := os.ReadDir(internal.PolynHomeDir + internal.PathSeparator + "node")
	if err != nil {
		// The node directory doesn't exist.
		// Passing a 'skip' to explicitly not treat this code path as an error.
		fmt.Println(constants.NO_DOWNLOADED_NODEJS_MESSAGE)
		return "", errors.New("skip")
	}

	prefix = utilities.ConvertToSemanticVersion(prefix)
	for _, item := range dir {
		if strings.HasPrefix(item.Name(), prefix) && item.IsDir() {
			return item.Name(), nil
		}
	}

	return "", errors.New("polyn error: no downloaded Node.js versions match the prefix '" + prefix + "'")
}

func convertPrefixToVersionLocalDesc(prefix string) (string, error) {
	dir, err := os.ReadDir(internal.PolynHomeDir + internal.PathSeparator + "node")
	if err != nil {
		// The node directory doesn't exist.
		// Passing a 'skip' to explicitly not treat this code path as an error.
		fmt.Println(constants.NO_DOWNLOADED_NODEJS_MESSAGE)
		return "", errors.New("skip")
	}

	prefix = utilities.ConvertToSemanticVersion(prefix)
	sort.Slice(dir, func(i int, j int) bool {
		return dir[i].Name() > dir[j].Name()
	})
	for _, item := range dir {
		if strings.HasPrefix(item.Name(), prefix) && item.IsDir() {
			return item.Name(), nil
		}
	}

	return "", errors.New("polyn error: no downloaded Node.js versions match the prefix '" + prefix + "'")
}

func getAllNodeVersions(config models.PolyNodeConfig) ([]models.NodeVersion, error) {
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

	var nodeVersions []models.NodeVersion
	err = json.NewDecoder(response.Body).Decode(&nodeVersions)
	return nodeVersions, err
}

func getArchiveName(operatingSystem models.OperatingSystem, arch models.Architecture) (string, error) {
	var archiveName string
	switch operatingSystem {
	case constants.AIX:
		archiveName = "aix-ppc64.tar.gz"
	case constants.LINUX:
		switch arch {
		case constants.ARM64:
			archiveName = "linux-arm64.tar.xz"
		case constants.PPC64LE:
			archiveName = "linux-ppc64le.tar.xz"
		case constants.S390X:
			archiveName = "linux-s390x.tar.xz"
		case constants.X64:
			archiveName = "linux-x64.tar.xz"
		default:
			return "", errors.New(constants.UNSUPPORTED_ARCH_ERROR)
		}
	case constants.MAC:
		if arch == constants.ARM64 {
			archiveName = "darwin-arm64.tar.gz"
		} else if arch == constants.X64 {
			archiveName = "darwin-x64.tar.gz"
		} else {
			return "", errors.New(constants.UNSUPPORTED_ARCH_ERROR)
		}
	case constants.WINDOWS:
		if arch == constants.ARM64 {
			archiveName = "win-arm64.zip"
		} else if arch == constants.X64 {
			archiveName = "win-x64.zip"
		} else {
			return "", errors.New(constants.UNSUPPORTED_ARCH_ERROR)
		}
	default:
		return "", errors.New(constants.UNSUPPORTED_OS_ERROR)
	}

	return archiveName, nil
}
