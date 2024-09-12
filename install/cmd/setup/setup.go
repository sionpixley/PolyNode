//go:build !windows

package main

import (
	"errors"
	"fmt"
	"install/internal"
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
		fmt.Println(err.Error())
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
	return os.WriteFile(home+"/.PolyNode/.polynrc", []byte(internal.DEFAULT_POLYNRC), 0644)
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
	if strings.Contains(shell, "bash") {
		return addToPath(home, ".bashrc")
	} else if strings.Contains(shell, "zsh") {
		return addToPath(home, ".zshrc")
	} else {
		return errors.New("unsupported shell")
	}
}

func oldVersionExists(home string) bool {
	_, err := os.Stat(home + "/.PolyNode")
	return !os.IsNotExist(err)
}

func upgrade(home string) error {
	err := os.RemoveAll(home + "/.PolyNode/polyn")
	if err != nil {
		return err
	}

	err = os.RemoveAll(home + "/.PolyNode/uninstall/uninstall")
	if err != nil {
		return err
	}

	err = exec.Command("cp", "./PolyNode/polyn", home+"/.PolyNode/polyn").Run()
	if err != nil {
		return err
	}

	return exec.Command("cp", "./PolyNode/uninstall/uninstall", home+"/.PolyNode/uninstall/uninstall").Run()
}
