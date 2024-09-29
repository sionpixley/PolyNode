package node

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/sionpixley/PolyNode/internal/constants"
	"github.com/sionpixley/PolyNode/internal/models"
	"github.com/sionpixley/PolyNode/pkg/polynconfig"
)

func convertKeywordToVersion(keyword string, config polynconfig.PolyNodeConfig) string {
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

func getAllNodeVersions(config polynconfig.PolyNodeConfig) ([]models.NodeVersion, error) {
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
	case constants.LINUX:
		if arch == constants.ARM64 {
			archiveName = "linux-arm64.tar.xz"
		} else if arch == constants.X64 {
			archiveName = "linux-x64.tar.xz"
		}
	case constants.MAC:
		if arch == constants.ARM64 {
			archiveName = "darwin-arm64.tar.gz"
		} else if arch == constants.X64 {
			archiveName = "darwin-x64.tar.gz"
		}
	case constants.WINDOWS:
		if arch == constants.ARM64 {
			archiveName = "win-arm64.zip"
		} else if arch == constants.X64 {
			archiveName = "win-x64.zip"
		}
	default:
		return "", errors.New(constants.UNSUPPORTED_OS_ERROR)
	}

	return archiveName, nil
}
