package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

const c_LINUX_TEMP string = `#!/bin/bash

rm -rf $HOME/.PolyNode
rm -f $HOME/polyn-uninstall-temp.sh`

const c_MAC_TEMP string = `#!/bin/zsh

rm -rf $HOME/.PolyNode
rm -f $HOME/polyn-uninstall-temp.zsh`

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

func removePathLinuxAndMac(home string) error {
	shell := os.Getenv("SHELL")
	if strings.Contains(shell, "bash") {
		return removePathFromBashrc(home)
	} else if strings.Contains(shell, "zsh") {
		return removePathFromZshrc(home)
	} else {
		return errors.New("unsupported shell")
	}
}

func removePathFromBashrc(home string) error {
	bashrc, err := os.Open(home + "/.bashrc")
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(bashrc)
	scanner.Split(bufio.ScanLines)
	content := ""
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.Contains(line, "PATH=$PATH:"+home+"/.PolyNode:"+home+"/.PolyNode/nodejs/bin") {
			content += line + "\n"
		}
	}

	// Explicitly calling close instead of using defer.
	// Need to have more control before writing to the file.
	bashrc.Close()

	err = os.WriteFile(home+"/.bashrc", []byte(content), 0644)
	return err
}

func removePathFromZshrc(home string) error {
	zshrc, err := os.Open(home + "/.zshrc")
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(zshrc)
	scanner.Split(bufio.ScanLines)
	content := ""
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.Contains(line, "PATH=$PATH:"+home+"/.PolyNode:"+home+"/.PolyNode/nodejs/bin") {
			content += line + "\n"
		}
	}

	// Explicitly calling close instead of using defer.
	// Need to have more control before writing to the file.
	zshrc.Close()

	err = os.WriteFile(home+"/.zshrc", []byte(content), 0644)
	return err
}

func uninstallLinux() error {
	home := os.Getenv("HOME")
	err := removePathLinuxAndMac(home)
	if err != nil {
		return err
	}

	err = os.WriteFile(home+"/polyn-uninstall-temp.sh", []byte(c_LINUX_TEMP), 0700)
	if err != nil {
		return err
	}

	return exec.Command("/bin/bash", "-c", home+"/polyn-uninstall-temp.sh").Run()
}

func uninstallMac() error {
	home := os.Getenv("HOME")
	err := removePathLinuxAndMac(home)
	if err != nil {
		return err
	}

	err = os.WriteFile(home+"/polyn-uninstall-temp.zsh", []byte(c_MAC_TEMP), 0700)
	if err != nil {
		return err
	}

	return exec.Command("/bin/zsh", "-c", home+"/polyn-uninstall-temp.zsh").Run()
}

func uninstallWindows() error {
	err := os.WriteFile("C:\\polyn-uninstall-temp.cmd", []byte(c_WIN_TEMP), 0700)
	if err != nil {
		return err
	}

	err = exec.Command("C:\\polyn-uninstall-temp").Run()
	return err
}
