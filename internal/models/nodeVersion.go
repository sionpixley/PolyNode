package models

import (
	"encoding/json"
	"reflect"
)

type NodeVersion struct {
	Version string `json:"version"`
	Lts     bool   `json:"lts"`
}

func (nodeVersion *NodeVersion) UnmarshalJSON(b []byte) error {
	var temp map[string]any
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
