package models

type NodeVersion struct {
	Version string `json:"version"`
	Lts     bool   `json:"lts"`
}
