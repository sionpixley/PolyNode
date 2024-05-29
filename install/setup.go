package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

const c_LINUX_PROFILE_D string = `#!/bin/bash

if [[ ":$PATH:" != *":/opt/PolyNode:"* ]]; then
    PATH=$PATH:/opt/PolyNode
fi`

const c_MAC_PROFILE_D string = `#!/bin/zsh

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
		fmt.Println("The polyn command has been installed.")
		fmt.Println("Please logout and log back in to apply your changes.")
	}
}

func installLinux() error {
	err := exec.Command("sudo", "cp", "-r", "PolyNode", "/opt").Run()
	if err != nil {
		return err
	}

	err = os.WriteFile("/etc/profile.d/polyn-path.sh", []byte(c_LINUX_PROFILE_D), 0755)
	return err
}

func installMac() error {
	err := exec.Command("sudo", "cp", "-r", "PolyNode", "/opt").Run()
	if err != nil {
		return err
	}

	err = os.WriteFile("/etc/profile.d/polyn-path.zsh", []byte(c_MAC_PROFILE_D), 0755)
	return err
}

func installWindows() error {
	return nil
}

func printOptionalLine(operatingSystem string) {
	if operatingSystem != "windows" {
		fmt.Println()
	}
}
