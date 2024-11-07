package models

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/sionpixley/PolyNode/internal"
)

const (
	DEFAULT_NODE_MIRROR string = "https://nodejs.org/dist"
)

var _DEFAULT_POLYNRC PolyNodeConfig = PolyNodeConfig{
	NodeMirror: DEFAULT_NODE_MIRROR,
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
		config.NodeMirror = DEFAULT_NODE_MIRROR
	}

	return nil
}

func LoadPolyNodeConfig() PolyNodeConfig {
	if _, err := os.Stat(internal.PolynHomeDir + internal.PathSeparator + "polynrc.json"); os.IsNotExist(err) {
		// Default config
		return _DEFAULT_POLYNRC
	} else if err != nil {
		// Default config
		return _DEFAULT_POLYNRC
	} else {
		content, err := os.ReadFile(internal.PolynHomeDir + internal.PathSeparator + "polynrc.json")
		if err != nil {
			// Default config
			return _DEFAULT_POLYNRC
		}

		config := PolyNodeConfig{}
		err = config.UnmarshalJSON(content)
		if err != nil {
			// Default config
			return _DEFAULT_POLYNRC
		}
		return config
	}
}
