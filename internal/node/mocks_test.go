package node

import "github.com/sionpixley/PolyNode/internal/models"

var nodeVersions = []models.NodeVersion{
	{
		"v2.0.0",
		[]string{
			"aix-ppc64",
			"linux-arm64",
			"linux-ppc64le",
			"linux-s390x",
			"linux-x64",
			"osx-arm64-tar",
			"osx-x64-tar",
			"win-arm64-zip",
			"win-x64-zip",
		},
		false,
	},
	{
		"v1.3.5",
		[]string{
			"aix-ppc64",
			"linux-arm64",
			"linux-ppc64le",
			"linux-s390x",
			"linux-x64",
			"osx-arm64-tar",
			"osx-x64-tar",
			"win-arm64-zip",
			"win-x64-zip",
		},
		true,
	},
	{
		"v1.3.4",
		[]string{
			"aix-ppc64",
			"linux-arm64",
			"linux-ppc64le",
			"linux-s390x",
			"linux-x64",
			"osx-arm64-tar",
			"osx-x64-tar",
			"win-arm64-zip",
			"win-x64-zip",
		},
		true,
	},
}

func getAllNodeVersionsMock(_ models.OperatingSystem, _ models.Architecture, _ *models.PolyNodeConfig) ([]models.NodeVersion, error) {
	return nodeVersions, nil
}
