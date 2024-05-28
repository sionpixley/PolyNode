package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

const c_FILE_CONTENT string = `#!/bin/bash

if [[ ":$PATH:" != *":/opt/PolyNode:"* ]]; then
    PATH=$PATH:/opt/PolyNode
fi`

func main() {
	operatingSystem := runtime.GOOS

	defer printOptionalLine(operatingSystem)

	var err error
	switch operatingSystem {
	case "darwin":
		err = installMac()
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
		fmt.Println("The polyn command has been installed. Please close all open terminals.")
	}
}

func installLinux() error {
	err := exec.Command("sudo", "cp", "-r", "PolyNode", "/opt").Run()
	if err != nil {
		return err
	}

	err = os.WriteFile("/etc/profile.d/polyn-path.sh", []byte(c_FILE_CONTENT), 0755)
	if err != nil {
		return err
	}

	err = exec.Command("bash", "-c", "source /etc/profile.d/polyn-path.sh").Run()
	return err
}

func installMac() error {
	return nil
}

func installWindows() error {
	return nil
}

func printOptionalLine(operatingSystem string) {
	if operatingSystem != "windows" {
		fmt.Println()
	}
}
