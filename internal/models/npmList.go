package models

type NPMPackage struct {
	Overridden bool   `json:"overridden"`
	Version    string `json:"version"`
}

type NPMList struct {
	Name         string                `json:"name"`
	Dependencies map[string]NPMPackage `json:"dependencies"`
}
