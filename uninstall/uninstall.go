package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

const c_LINUX_TEMP string = `#!/bin/bash

sudo rm -f /opt/nodejs
sudo rm -rf /opt/PolyNode
sudo rm -f /opt/polyn-uninstall-temp.sh`

const c_MAC_TEMP string = `#!/bin/zsh

sudo rm -f /opt/nodejs
sudo rm -rf /opt/PolyNode
sudo rm -f /opt/polyn-uninstall-temp.zsh`

const c_WIN_TEMP string = `del C:\\Program Files\\PolyNode /s /f /q > nul
rmdir C:\\Program Files\\PolyNode /s /q
del C:\\polyn-uninstall-temp.cmd`

func main() {
	operatingSystem := runtime.GOOS

	defer printOptionalLine(operatingSystem)

	var err error
	switch operatingSystem {
	case "darwin":
		err = uninstallMac()
	case "linux":
		err = uninstallLinux()
	case "windows":
		err = uninstallWindows()
	default:
		err = errors.New("unsupported operating system")
	}

	if err != nil {
		fmt.Println(err.Error())
	}
}

func printOptionalLine(operatingSystem string) {
	if operatingSystem != "windows" {
		fmt.Println()
	}
}

func uninstallLinux() error {
	err := os.Remove("/etc/profile.d/polyn-path.sh")
	if err != nil {
		return err
	}

	err = os.WriteFile("/opt/polyn-uninstall-temp.sh", []byte(c_LINUX_TEMP), 0700)
	if err != nil {
		return err
	}

	err = exec.Command("sudo", "/opt/polyn-uninstall-temp.sh").Run()
	return err
}

func uninstallMac() error {
	err := os.Remove("/etc/profile.d/polyn-path.zsh")
	if err != nil {
		return err
	}

	err = os.WriteFile("/opt/polyn-uninstall-temp.zsh", []byte(c_MAC_TEMP), 0700)
	if err != nil {
		return err
	}

	err = exec.Command("sudo", "/opt/polyn-uninstall-temp.zsh").Run()
	return err
}

func uninstallWindows() error {
	err := os.WriteFile("C:\\polyn-uninstall-temp.cmd", []byte(c_WIN_TEMP), 0700)
	if err != nil {
		return err
	}

	err = exec.Command("C:\\polyn-uninstall-temp").Run()
	return err
}
