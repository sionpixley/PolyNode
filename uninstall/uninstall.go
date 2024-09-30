//go:build !windows

package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

const _UNINSTALL_SCRIPT string = `#!/bin/sh

rm -rf $HOME/.PolyNode
rm $HOME/polyn-uninstall-temp`

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
		err = uninstall()
	default:
		err = errors.New("unsupported operating system")
	}

	if err != nil {
		log.Fatal(err.Error())
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
		if !strings.Contains(line, "export PATH=$PATH:"+home+"/.PolyNode:"+home+"/.PolyNode/nodejs/bin") {
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
	if strings.HasSuffix(shell, "/bash") {
		return removePath(home, ".bashrc")
	} else if strings.HasSuffix(shell, "/zsh") {
		return removePath(home, ".zshrc")
	} else if strings.HasSuffix(shell, "/ksh") {
		return removePath(home, ".kshrc")
	} else {
		return errors.New("unsupported shell")
	}
}

func uninstall() error {
	home := os.Getenv("HOME")
	err := removePathLinuxAndMac(home)
	if err != nil {
		return err
	}

	err = os.WriteFile(home+"/polyn-uninstall-temp", []byte(_UNINSTALL_SCRIPT), 0700)
	if err != nil {
		return err
	}

	return exec.Command(home + "/polyn-uninstall-temp").Run()
}
