package internal

import (
	"encoding/json"
	"reflect"
)

type Architecture int

type NodeVersion struct {
	Version string `json:"version"`
	Date    string `json:"date"`
	Lts     bool   `json:"lts"`
}

func (nodeVersion *NodeVersion) UnmarshalJSON(b []byte) error {
	var temp map[string]interface{}
	err := json.Unmarshal(b, &temp)
	if err != nil {
		return err
	}

	nodeVersion.Version = temp["version"].(string)
	nodeVersion.Date = temp["date"].(string)

	if reflect.TypeOf(temp["lts"]).String() == "bool" {
		nodeVersion.Lts = false
	} else {
		nodeVersion.Lts = true
	}

	return nil
}

type OperatingSystem int

type command int
