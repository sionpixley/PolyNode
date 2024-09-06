package internal

import (
	"encoding/json"
	"reflect"
	"strings"
)

type Architecture int

type NodeVersion struct {
	Version string `json:"version"`
	Lts     bool   `json:"lts"`
}

func (nodeVersion *NodeVersion) UnmarshalJSON(b []byte) error {
	var temp map[string]interface{}
	err := json.Unmarshal(b, &temp)
	if err != nil {
		return err
	}

	nodeVersion.Version = temp["version"].(string)

	if reflect.TypeOf(temp["lts"]).String() == "bool" {
		nodeVersion.Lts = false
	} else {
		nodeVersion.Lts = true
	}

	return nil
}

type PolyNodeConfig struct {
	NodeMirror string `json:"nodeMirror"`
}

func (config *PolyNodeConfig) UnmarshalJSON(b []byte) error {
	var temp map[string]string
	err := json.Unmarshal(b, &temp)
	if err != nil {
		return err
	}

	var exists bool
	config.NodeMirror, exists = temp["nodeMirror"]
	if exists {
		config.NodeMirror = strings.ToLower(strings.TrimSuffix(strings.TrimSpace(config.NodeMirror), "/"))
	} else {
		config.NodeMirror = _DEFAULT_NODE_MIRROR
	}

	return nil
}

type OperatingSystem int

type command int
