package node

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/sionpixley/PolyNode/internal"
	"github.com/sionpixley/PolyNode/internal/constants"
	"github.com/sionpixley/PolyNode/internal/models"
	"github.com/sionpixley/PolyNode/internal/utilities"
)

func convertKeywordToVersion(keyword string, operatingSystem models.OperatingSystem, arch models.Architecture, config models.PolyNodeConfig) string {
	if keyword == "lts" {
		nodeVersions, err := getAllNodeVersionsForOsAndArch(operatingSystem, arch, config)
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
		nodeVersions, err := getAllNodeVersionsForOsAndArch(operatingSystem, arch, config)
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
	case constants.Aix:
		return "aix-ppc64", nil
	case constants.Linux:
		switch arch {
		case constants.Arm64:
			return "linux-arm64", nil
		case constants.Ppc64Le:
			return "linux-ppc64le", nil
		case constants.S390x:
			return "linux-s390x", nil
		case constants.X64:
			return "linux-x64", nil
		default:
			return "", errors.New(constants.UnsupportedArchError)
		}
	case constants.Mac:
		if arch == constants.Arm64 {
			return "osx-arm64-tar", nil
		} else if arch == constants.X64 {
			return "osx-x64-tar", nil
		} else {
			return "", errors.New(constants.UnsupportedArchError)
		}
	case constants.Windows:
		if arch == constants.Arm64 {
			return "win-arm64-zip", nil
		} else if arch == constants.X64 {
			return "win-x64-zip", nil
		} else {
			return "", errors.New(constants.UnsupportedArchError)
		}
	default:
		return "", errors.New(constants.UnsupportedOSError)
	}
}

func convertPrefixToVersionDown(prefix string, operatingSystem models.OperatingSystem, arch models.Architecture, config models.PolyNodeConfig) (string, error) {
	nodeVersions, err := getAllNodeVersionsForOsAndArch(operatingSystem, arch, config)
	if err != nil {
		return "", err
	}

	prefix = utilities.ConvertToSemanticVersion(prefix)
	for _, nodeVersion := range nodeVersions {
		if strings.HasPrefix(nodeVersion.Version, prefix) {
			return nodeVersion.Version, nil
		}
	}

	return "", fmt.Errorf("polyn error: no Node.js versions match the prefix '%s'", prefix)
}

func convertPrefixToVersionLocalAsc(prefix string) (string, error) {
	dir, err := os.ReadDir(internal.PolynHomeDir + internal.PathSeparator + "node")
	if err != nil {
		// The node directory doesn't exist.
		// Passing a 'skip' to explicitly not treat this code path as an error.
		fmt.Println(constants.NoDownloadedNodejsMessage)
		return "", errors.New("skip")
	}

	prefix = utilities.ConvertToSemanticVersion(prefix)
	sort.Slice(dir, func(i int, j int) bool {
		left := dir[i].Name()
		right := dir[j].Name()

		leftParts := strings.Split(left[1:], ".")
		rightParts := strings.Split(right[1:], ".")

		for k := range len(leftParts) {
			leftVal, err := strconv.Atoi(leftParts[k])
			if err != nil {
				panic(err)
			}
			rightVal, err := strconv.Atoi(rightParts[k])
			if err != nil {
				panic(err)
			}

			if leftVal != rightVal {
				return leftVal < rightVal
			}
		}

		return true
	})
	for _, item := range dir {
		if strings.HasPrefix(item.Name(), prefix) && item.IsDir() {
			return item.Name(), nil
		}
	}

	return "", fmt.Errorf("polyn error: no downloaded Node.js versions match the prefix '%s'", prefix)
}

func convertPrefixToVersionLocalDesc(prefix string) (string, error) {
	dir, err := os.ReadDir(internal.PolynHomeDir + internal.PathSeparator + "node")
	if err != nil {
		// The node directory doesn't exist.
		// Passing a 'skip' to explicitly not treat this code path as an error.
		fmt.Println(constants.NoDownloadedNodejsMessage)
		return "", errors.New("skip")
	}

	prefix = utilities.ConvertToSemanticVersion(prefix)
	sort.Slice(dir, func(i int, j int) bool {
		left := dir[i].Name()
		right := dir[j].Name()

		leftParts := strings.Split(left[1:], ".")
		rightParts := strings.Split(right[1:], ".")

		for k := range len(leftParts) {
			leftVal, err := strconv.Atoi(leftParts[k])
			if err != nil {
				panic(err)
			}
			rightVal, err := strconv.Atoi(rightParts[k])
			if err != nil {
				panic(err)
			}

			if leftVal != rightVal {
				return leftVal > rightVal
			}
		}

		return true
	})
	for _, item := range dir {
		if strings.HasPrefix(item.Name(), prefix) && item.IsDir() {
			return item.Name(), nil
		}
	}

	return "", fmt.Errorf("polyn error: no downloaded Node.js versions match the prefix '%s'", prefix)
}

func getAllNodeVersionsForOsAndArch(operatingSystem models.OperatingSystem, arch models.Architecture, config models.PolyNodeConfig) ([]models.NodeVersion, error) {
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
	if err != nil {
		return nil, err
	}

	nodeVersionFile, err := convertOsAndArchToNodeVersionFile(operatingSystem, arch)
	if err != nil {
		return nil, err
	}

	compatibleNodeVersions := []models.NodeVersion{}
	for _, nodeVersion := range nodeVersions {
		if slices.Contains(nodeVersion.Files, nodeVersionFile) {
			compatibleNodeVersions = append(compatibleNodeVersions, nodeVersion)
		}
	}

	return compatibleNodeVersions, err
}

func getArchiveName(operatingSystem models.OperatingSystem, arch models.Architecture) (string, error) {
	var archiveName string
	switch operatingSystem {
	case constants.Aix:
		archiveName = "aix-ppc64.tar.gz"
	case constants.Linux:
		switch arch {
		case constants.Arm64:
			archiveName = "linux-arm64.tar.gz"
		case constants.Ppc64Le:
			archiveName = "linux-ppc64le.tar.gz"
		case constants.S390x:
			archiveName = "linux-s390x.tar.gz"
		case constants.X64:
			archiveName = "linux-x64.tar.gz"
		default:
			return "", errors.New(constants.UnsupportedArchError)
		}
	case constants.Mac:
		if arch == constants.Arm64 {
			archiveName = "darwin-arm64.tar.gz"
		} else if arch == constants.X64 {
			archiveName = "darwin-x64.tar.gz"
		} else {
			return "", errors.New(constants.UnsupportedArchError)
		}
	case constants.Windows:
		if arch == constants.Arm64 {
			archiveName = "win-arm64.zip"
		} else if arch == constants.X64 {
			archiveName = "win-x64.zip"
		} else {
			return "", errors.New(constants.UnsupportedArchError)
		}
	default:
		return "", errors.New(constants.UnsupportedOSError)
	}

	return archiveName, nil
}
