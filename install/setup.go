package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func main() {
	operatingSystem := runtime.GOOS

	defer printOptionalLine(operatingSystem)

	var err error
	switch operatingSystem {
	case "darwin":
		fallthrough
	case "linux":
		err = installLinux()
	case "windows":
		err = installWindows()
	default:
		err = errors.New("unsupported operating system")
	}

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("The polyn command has been installed.")
		fmt.Println("Please close all open terminals.")
	}
}

func addToBashPath(home string) error {
	contentData, err := os.ReadFile(home + "/.bashrc")
	if err != nil {
		return err
	}

	content := string(contentData)
	content += "\nPATH=$PATH:" + home + "/.PolyNode:" + home + "/.PolyNode/nodejs/bin"

	err = os.WriteFile(home+"/.bashrc", []byte(content), 0644)
	return err
}

func addToWindowsPath(home string) error {
}

func addToZshPath(home string) error {
	contentData, err := os.ReadFile(home + "/.zshrc")
	if err != nil {
		return err
	}

	content := string(contentData)
	content += "\nPATH=$PATH:" + home + "/.PolyNode:" + home + "/.PolyNode/nodejs/bin"

	err = os.WriteFile(home+"/.zshrc", []byte(content), 0644)
	return err
}

func installLinux() error {
	home := os.Getenv("HOME")
	err := exec.Command("cp", "-r", "PolyNode", home+"/.PolyNode").Run()
	if err != nil {
		return err
	}

	shell := os.Getenv("SHELL")
	if strings.Contains(shell, "bash") {
		return addToBashPath(home)
	} else if strings.Contains(shell, "zsh") {
		return addToZshPath(home)
	} else {
		return errors.New("unsupported shell")
	}
}

func installWindows() error {
	home := os.Getenv("PROGRAMFILES")
	err := exec.Command("xcopy", "/s", "PolyNode", home+"\\PolyNode").Run()
	if err != nil {
		return err
	}

	return addToWindowsPath(home)
}

func printOptionalLine(operatingSystem string) {
	if operatingSystem != "windows" {
		fmt.Println()
	}
}
