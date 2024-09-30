package models

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/sionpixley/PolyNode/internal"
)

const DEFAULT_NODE_MIRROR string = "https://nodejs.org/dist"

type PolyNodeConfig struct {
	NodeMirror string `json:"nodeMirror"`
}

func (config *PolyNodeConfig) UnmarshalJSON(b []byte) error {
	var temp map[string]string
	err := json.Unmarshal(b, &temp)
	if err != nil {
		return err
	}

	mirror, exists := temp["nodeMirror"]
	if exists {
		config.NodeMirror = strings.ToLower(strings.TrimSuffix(strings.TrimSpace(mirror), "/"))
	} else {
		config.NodeMirror = DEFAULT_NODE_MIRROR
	}

	return nil
}

func LoadPolyNodeConfig() PolyNodeConfig {
	if _, err := os.Stat(internal.PolynHomeDir + internal.PathSeparator + ".polynrc"); os.IsNotExist(err) {
		// Default config
		return PolyNodeConfig{NodeMirror: DEFAULT_NODE_MIRROR}
	} else {
		content, err := os.ReadFile(internal.PolynHomeDir + internal.PathSeparator + ".polynrc")
		if err != nil {
			// Default config
			return PolyNodeConfig{NodeMirror: DEFAULT_NODE_MIRROR}
		}

		config := PolyNodeConfig{}
		err = config.UnmarshalJSON(content)
		if err != nil {
			// Default config
			return PolyNodeConfig{NodeMirror: DEFAULT_NODE_MIRROR}
		}
		return config
	}
}
