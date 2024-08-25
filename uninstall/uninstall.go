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

const linuxTemp string = `#!/bin/bash

rm -rf $HOME/.PolyNode
rm -f $HOME/polyn-uninstall-temp.sh`

const macTemp string = `#!/bin/zsh

rm -rf $HOME/.PolyNode
rm -f $HOME/polyn-uninstall-temp.zsh`

func main() {
	operatingSystem := runtime.GOOS

	defer fmt.Println()

	var err error
	switch operatingSystem {
	case "darwin":
		err = uninstallMac()
	case "linux":
		err = uninstallLinux()
	default:
		err = errors.New("unsupported operating system")
	}

	if err != nil {
		fmt.Println(err.Error())
	}
}

func removePath(home string, rcFile string) error {
	rc, err := os.Open(home + "/" + rcFile)
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(rc)
	scanner.Split(bufio.ScanLines)
	content := ""
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.Contains(line, "PATH=$PATH:"+home+"/.PolyNode:"+home+"/.PolyNode/nodejs/bin") {
			content += line + "\n"
		}
	}

	// Removing the extra new line character.
	if len(content) > 0 {
		content = content[:len(content)-1]
	}

	// Explicitly calling close instead of using defer.
	// Need to have more control before writing to the file.
	rc.Close()

	return os.WriteFile(home+"/"+rcFile, []byte(content), 0644)
}

func removePathLinuxAndMac(home string) error {
	shell := os.Getenv("SHELL")
	if strings.Contains(shell, "bash") {
		return removePath(home, ".bashrc")
	} else if strings.Contains(shell, "zsh") {
		return removePath(home, ".zshrc")
	} else {
		return errors.New("unsupported shell")
	}
}

func uninstallLinux() error {
	home := os.Getenv("HOME")
	err := removePathLinuxAndMac(home)
	if err != nil {
		return err
	}

	err = os.WriteFile(home+"/polyn-uninstall-temp.sh", []byte(linuxTemp), 0700)
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

	err = os.WriteFile(home+"/polyn-uninstall-temp.zsh", []byte(macTemp), 0700)
	if err != nil {
		return err
	}

	return exec.Command("/bin/zsh", "-c", home+"/polyn-uninstall-temp.zsh").Run()
}
