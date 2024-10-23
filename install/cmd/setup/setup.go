//go:build !windows

package main

import (
	"errors"
	"fmt"
	"install/internal/constants"
	"install/internal/utilities"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func main() {
	operatingSystem := runtime.GOOS

	defer fmt.Println()

	var err error
	switch operatingSystem {
	case "aix":
		fallthrough
	case "darwin":
		fallthrough
	case "linux":
		home := os.Getenv("HOME")
		if oldVersionExists(home) {
			err = upgrade(home)
		} else {
			err = install(home)
		}
	default:
		err = errors.New("unsupported operating system")
	}

	if err != nil {
		log.Fatal(err.Error())
	} else {
		fmt.Println("The polyn command has been installed.")
		fmt.Println("Please close all open terminals.")
	}
}

func addToPath(home string, rcFile string) error {
	// Creating the file if it doesn't exist.
	f, err := os.OpenFile(home+"/"+rcFile, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	// Calling close directly instead of with defer. Will be reopening the file soon.
	f.Close()

	contentData, err := os.ReadFile(home + "/" + rcFile)
	if err != nil {
		return err
	}

	content := string(contentData)
	content += "\nexport PATH=$PATH:" + home + "/.PolyNode:" + home + "/.PolyNode/nodejs/bin"

	return os.WriteFile(home+"/"+rcFile, []byte(content), 0644)
}

func createPolynConfig(home string) error {
	return os.WriteFile(home+"/.PolyNode/polynrc.json", []byte(constants.DEFAULT_POLYNRC), 0644)
}

func install(home string) error {
	err := exec.Command("cp", "-r", "PolyNode", home+"/.PolyNode").Run()
	if err != nil {
		return err
	}

	err = createPolynConfig(home)
	if err != nil {
		return err
	}

	shell := os.Getenv("SHELL")
	if strings.HasSuffix(shell, "/bash") {
		return addToPath(home, ".bashrc")
	} else if strings.HasSuffix(shell, "/zsh") {
		return addToPath(home, ".zshrc")
	} else if strings.HasSuffix(shell, "/ksh") {
		return addToPath(home, ".kshrc")
	} else {
		return errors.New("unsupported shell")
	}
}

func oldVersionExists(home string) bool {
	if _, err := os.Stat(home + "/.PolyNode"); os.IsNotExist(err) {
		return false
	} else if err != nil {
		return false
	} else {
		return true
	}
}

func upgrade(home string) error {
	err := utilities.RemoveUpgradableFiles(home)
	if err != nil {
		return err
	}

	return utilities.CopyUpgradableFiles(home)
}
