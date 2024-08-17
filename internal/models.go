package internal

type Architecture int

type NodeVersion struct {
	Version  string   `json:"version"`
	Date     string   `json:"date"`
	Files    []string `json:"files"`
	Npm      string   `json:"npm"`
	V8       string   `json:"v8"`
	Uv       string   `json:"uv"`
	Zlib     string   `json:"zlib"`
	Openssl  string   `json:"openssl"`
	Modules  string   `json:"modules"`
	Lts      string   `json:"lts"`
	Security bool     `json:"security"`
}

type OperatingSystem int

type command int
