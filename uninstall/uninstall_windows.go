package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"golang.org/x/sys/windows/registry"
)

const windowsUninstallScript string = `@echo off

del %LOCALAPPDATA%\Programs\PolyNode /s /f /q > nul
rmdir %LOCALAPPDATA%\Programs\PolyNode /s /q
del %LOCALAPPDATA%\Programs\polyn-uninstall-temp.cmd`

func main() {
	err := uninstall()
	if err != nil {
		fmt.Println(err.Error())
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

	updatedPath := ""
	parts := strings.Split(path, ";")
	for _, part := range parts {
		if part == home+"\\PolyNode" || part == home+"\\PolyNode\\nodejs" {
			continue
		} else {
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

	err = os.WriteFile(home+"\\polyn-uninstall-temp.cmd", []byte(windowsUninstallScript), 0700)
	if err != nil {
		return err
	}

	return exec.Command("cmd", "/c", home+"\\polyn-uninstall-temp.cmd").Run()
}
