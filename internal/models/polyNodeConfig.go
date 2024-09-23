package models

import (
	"encoding/json"
	"strings"

	"github.com/sionpixley/PolyNode/internal"
)

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
		config.NodeMirror = internal.DEFAULT_NODE_MIRROR
	}

	return nil
}
