package models

import (
	"encoding/json"
)

type NodeVersion struct {
	Version string   `json:"version"`
	Files   []string `json:"files"`
	Lts     bool     `json:"lts"`
}

func (nodeVersion *NodeVersion) UnmarshalJSON(b []byte) error {
	var temp map[string]any
	err := json.Unmarshal(b, &temp)
	if err != nil {
		return err
	}

	nodeVersion.Version = temp["version"].(string)

	nodeVersion.Files = []string{}
	rawFiles := temp["files"].([]any)
	for _, rawFile := range rawFiles {
		nodeVersion.Files = append(nodeVersion.Files, rawFile.(string))
	}

	switch temp["lts"].(type) {
	case bool:
		nodeVersion.Lts = false
	default:
		nodeVersion.Lts = true
	}

	return nil
}
