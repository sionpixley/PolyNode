package models

type NPMPackage struct {
	Version    string `json:"version"`
	Overridden bool   `json:"overridden"`
}

type NPMList struct {
	Dependencies map[string]NPMPackage `json:"dependencies"`
	Name         string                `json:"name"`
}
