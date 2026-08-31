package models

import (
	"encoding/json"
	"fmt"

	"github.com/sionpixley/PolyNode/internal/constants"
)

type NodeVersion struct {
	Version string   `json:"version"`
	Files   []string `json:"files"`
	LTS     bool     `json:"lts"`
}

func (nodeVersion *NodeVersion) UnmarshalJSON(b []byte) error {
	var temp map[string]any
	err := json.Unmarshal(b, &temp)
	if err != nil {
		return err
	}

	if version, ok := temp["version"].(string); ok {
		nodeVersion.Version = version
	} else {
		return fmt.Errorf(constants.InvalidJSONResponseError, "nodeVersion")
	}

	if rawFiles, ok := temp["files"].([]any); ok {
		nodeVersion.Files = make([]string, len(rawFiles))
		for i, rawFile := range rawFiles {
			if file, ok := rawFile.(string); ok {
				nodeVersion.Files[i] = file
			} else {
				return fmt.Errorf(constants.InvalidJSONResponseError, "nodeVersion")
			}
		}
	} else {
		return fmt.Errorf(constants.InvalidJSONResponseError, "nodeVersion")
	}

	switch temp["lts"].(type) {
	case bool:
		nodeVersion.LTS = false
	default:
		nodeVersion.LTS = true
	}

	return nil
}
