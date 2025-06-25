package main

import (
	"log"
	"os"
	"strings"

	"golang.org/x/sys/windows/registry"
)

func main() {
	err := uninstall()
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func removePath(home string) error {
	key, err := registry.OpenKey(registry.CURRENT_USER, "Environment", registry.ALL_ACCESS)
	if err != nil {
		return err
	}
	defer key.Close()

	path, _, err := key.GetStringValue("Path")
	if err != nil {
		return err
	}

	var updatedPath string
	parts := strings.Split(path, ";")
	for _, part := range parts {
		if part != home+"\\PolyNode" && part != home+"\\PolyNode\\nodejs" {
			updatedPath += part + ";"
		}
	}
	updatedPath = strings.TrimSuffix(updatedPath, ";")

	return key.SetStringValue("Path", updatedPath)
}

func uninstall() error {
	home := os.Getenv("LOCALAPPDATA") + "\\Programs"
	err := removePath(home)
	if err != nil {
		return err
	}

	err = os.RemoveAll(home + "\\PolyNode")
	return err
}
