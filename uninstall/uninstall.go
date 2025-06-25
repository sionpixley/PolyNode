//go:build !windows

package main

import (
	"bufio"
	"errors"
	"log"
	"os"
	"runtime"
	"strings"
)

func main() {
	operatingSystem := runtime.GOOS

	var err error
	switch operatingSystem {
	case "aix":
		fallthrough
	case "darwin":
		fallthrough
	case "linux":
		err = uninstall()
	default:
		err = errors.New("uninstall: unsupported operating system")
	}

	if err != nil {
		log.Fatalln(err.Error())
	}
}

func removePath(home string, rcFile string) error {
	rc, err := os.Open(home + "/" + rcFile)
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(rc)
	scanner.Split(bufio.ScanLines)
	var content string
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
	switch {
	case strings.HasSuffix(shell, "/bash"):
		return removePath(home, ".bashrc")
	case strings.HasSuffix(shell, "/zsh"):
		return removePath(home, ".zshrc")
	case strings.HasSuffix(shell, "/ksh"):
		return removePath(home, ".kshrc")
	default:
		return errors.New("uninstall: unsupported shell")
	}
}

func uninstall() error {
	home := os.Getenv("HOME")
	err := removePathLinuxAndMac(home)
	if err != nil {
		return err
	}

	err = os.RemoveAll(home + "/.PolyNode")
	return err
}
