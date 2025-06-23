package models

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/sionpixley/PolyNode/internal"
)

const (
	defaultNodeMirror string = "https://nodejs.org/dist"
)

var defaultPolynrc PolyNodeConfig = PolyNodeConfig{
	NodeMirror: defaultNodeMirror,
}

type PolyNodeConfig struct {
	NodeMirror string `json:"nodeMirror"`
}

func (config *PolyNodeConfig) UnmarshalJSON(b []byte) error {
	var temp map[string]any
	err := json.Unmarshal(b, &temp)
	if err != nil {
		return err
	}

	mirror, exists := temp["nodeMirror"]
	if exists {
		config.NodeMirror = strings.ToLower(strings.TrimSuffix(strings.TrimSpace(mirror.(string)), "/"))
	} else {
		config.NodeMirror = defaultNodeMirror
	}

	return nil
}

func LoadPolyNodeConfig() PolyNodeConfig {
	if _, err := os.Stat(internal.PolynHomeDir + internal.PathSeparator + "polynrc.json"); os.IsNotExist(err) {
		// Default config
		return defaultPolynrc
	} else if err != nil {
		// Default config
		return defaultPolynrc
	} else {
		content, err := os.ReadFile(internal.PolynHomeDir + internal.PathSeparator + "polynrc.json")
		if err != nil {
			// Default config
			return defaultPolynrc
		}

		var config PolyNodeConfig
		err = config.UnmarshalJSON(content)
		if err != nil {
			// Default config
			return defaultPolynrc
		}
		return config
	}
}
