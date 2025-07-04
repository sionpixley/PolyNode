package main

import (
	"log"
	"os"
	"os/exec"
	"strings"

	"golang.org/x/sys/windows/registry"
)

const uninstallBatch string = `@echo off
timeout /t 1 /nobreak > nul
cd %LOCALAPPDATA%
del %LOCALAPPDATA%\Programs\PolyNode /s /f /q > nul
rmdir %LOCALAPPDATA%\Programs\PolyNode /s /q
if exist %LOCALAPPDATA%\Programs\PolyNode\ (
  del %LOCALAPPDATA%\Programs\PolyNode /s /q > nul
  rmdir %LOCALAPPDATA%\Programs\PolyNode /s /q
)
(goto) 2>nul & del "%~f0"`

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

	err = os.WriteFile(home+"\\polyn-uninstall-temp.cmd", []byte(uninstallBatch), 0744)
	if err != nil {
		return err
	}

	return exec.Command("cmd", "/c", "start", "/b", home+"\\polyn-uninstall-temp").Run()
}
