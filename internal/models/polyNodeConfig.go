package models

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/sionpixley/PolyNode/internal"
)

const (
	defaultAutoUpdate bool   = true
	defaultNodeMirror string = "https://nodejs.org/dist"
)

var defaultPolynrc PolyNodeConfig = PolyNodeConfig{
	AutoUpdate: defaultAutoUpdate,
	NodeMirror: defaultNodeMirror,
}

type PolyNodeConfig struct {
	AutoUpdate bool   `json:"autoUpdate"`
	NodeMirror string `json:"nodeMirror"`
}

func (config *PolyNodeConfig) UnmarshalJSON(b []byte) error {
	var temp map[string]any
	err := json.Unmarshal(b, &temp)
	if err != nil {
		return err
	}

	autoUpdate, exists := temp["autoUpdate"]
	if exists {
		config.AutoUpdate = autoUpdate.(bool)
	} else {
		config.AutoUpdate = defaultAutoUpdate
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
	configPath := internal.PolynHomeDir + internal.PathSeparator + "polynrc.json"
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		// Default config
		return defaultPolynrc
	} else if err != nil {
		// Default config
		return defaultPolynrc
	}

	content, err := os.ReadFile(configPath)
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
