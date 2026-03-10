package models

import (
	"encoding/json"
	"strings"

	"github.com/sionpixley/PolyNode/internal"
)

const (
	defaultAutoUpdate       = true
	defaultNodeMirror       = "https://nodejs.org/dist"
	defaultTimeoutInSeconds = 180
)

var defaultPolynrc = PolyNodeConfig{
	AutoUpdate:       defaultAutoUpdate,
	NodeMirror:       defaultNodeMirror,
	TimeoutInSeconds: defaultTimeoutInSeconds,
}

type PolyNodeConfig struct {
	NodeMirror       string `json:"nodeMirror"`
	AutoUpdate       bool   `json:"autoUpdate"`
	TimeoutInSeconds int    `json:"timeoutInSeconds"`
}

func (config *PolyNodeConfig) Save(osWrapper OSWrapper) error {
	configPath := internal.PolynHomeDir + internal.PathSeparator + "polynrc.json"
	jsonBytes, err := json.Marshal(config)
	if err != nil {
		return err
	}

	return osWrapper.WriteFile(configPath, jsonBytes, 0644)
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

	timeout, exists := temp["timeout"]
	if exists {
		val := timeout.(int)
		if val < 0 {
			config.TimeoutInSeconds = defaultTimeoutInSeconds
		} else {
			config.TimeoutInSeconds = val
		}
	} else {
		config.TimeoutInSeconds = defaultTimeoutInSeconds
	}

	return nil
}

func NewPolyNodeConfig(osWrapper OSWrapper) *PolyNodeConfig {
	configPath := internal.PolynHomeDir + internal.PathSeparator + "polynrc.json"
	if _, err := osWrapper.Stat(configPath); osWrapper.IsNotExist(err) {
		// Default config
		return &defaultPolynrc
	} else if err != nil {
		// Default config
		return &defaultPolynrc
	}

	content, err := osWrapper.ReadFile(configPath)
	if err != nil {
		// Default config
		return &defaultPolynrc
	}

	config := new(PolyNodeConfig)
	err = config.UnmarshalJSON(content)
	if err != nil {
		// Default config
		return &defaultPolynrc
	}

	return config
}
