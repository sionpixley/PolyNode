package node

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/sionpixley/PolyNode/internal/constants"
	"github.com/sionpixley/PolyNode/internal/models"
	"github.com/sionpixley/PolyNode/pkg/polynrc"
)

func convertKeywordToVersion(keyword string, config polynrc.PolyNodeConfig) string {
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

func getAllNodeVersions(config polynrc.PolyNodeConfig) ([]models.NodeVersion, error) {
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
	archiveName := ""
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
