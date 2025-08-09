//go:build !windows

package main

import (
	"errors"
	"install/internal/constants"
	"os"
	"os/exec"
	"strings"
	"time"
)

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

func copyUpdatableFiles(currentBinaryLocation string, home string) error {
	err := exec.Command("cp", currentBinaryLocation+"/PolyNode/polyn", home+"/.PolyNode/polyn").Run()
	if err != nil {
		return err
	}

	err = exec.Command("cp", currentBinaryLocation+"/PolyNode/LICENSE", home+"/.PolyNode/LICENSE").Run()
	if err != nil {
		return err
	}

	err = exec.Command("cp", currentBinaryLocation+"/PolyNode/README.md", home+"/.PolyNode/README.md").Run()
	if err != nil {
		return err
	}

	err = exec.Command("cp", currentBinaryLocation+"/PolyNode/SECURITY.md", home+"/.PolyNode/SECURITY.md").Run()
	if err != nil {
		return err
	}

	err = exec.Command("cp", currentBinaryLocation+"/PolyNode/uninstall/uninstall", home+"/.PolyNode/uninstall/uninstall").Run()
	if err != nil {
		return err
	}

	return createLastUpdateFile(home)
}

func createLastUpdateFile(home string) error {
	now := time.Now().UTC()
	return os.WriteFile(home+"/.PolyNode/last-update.txt", []byte(now.Format(constants.ISODateTimeFormat)), 0644)
}

func createPolynConfig(home string) error {
	return os.WriteFile(home+"/.PolyNode/polynrc.json", []byte(constants.DefaultPolynrc), 0644)
}

func install(currentBinaryLocation string, home string) error {
	err := exec.Command("cp", "-r", currentBinaryLocation+"/PolyNode", home+"/.PolyNode").Run()
	if err != nil {
		return err
	}

	err = createPolynConfig(home)
	if err != nil {
		return err
	}

	err = createLastUpdateFile(home)
	if err != nil {
		return err
	}

	shell := os.Getenv("SHELL")
	switch {
	case strings.HasSuffix(shell, "/bash"):
		return addToPath(home, ".bashrc")
	case strings.HasSuffix(shell, "/zsh"):
		return addToPath(home, ".zshrc")
	case strings.HasSuffix(shell, "/ksh"):
		return addToPath(home, ".kshrc")
	default:
		return errors.New("setup: unsupported shell")
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

func removeUpdatableFiles(home string) error {
	err := os.RemoveAll(home + "/.PolyNode/polyn")
	if err != nil {
		return err
	}

	err = os.RemoveAll(home + "/.PolyNode/LICENSE")
	if err != nil {
		return err
	}

	err = os.RemoveAll(home + "/.PolyNode/README.md")
	if err != nil {
		return err
	}

	err = os.RemoveAll(home + "/.PolyNode/SECURITY.md")
	if err != nil {
		return err
	}

	err = os.RemoveAll(home + "/.PolyNode/lastAutoUpdate.txt")
	if err != nil {
		return err
	}

	return os.RemoveAll(home + "/.PolyNode/uninstall/uninstall")
}

func update(currentBinaryLocation string, home string) error {
	err := removeUpdatableFiles(home)
	if err != nil {
		return err
	}

	return copyUpdatableFiles(currentBinaryLocation, home)
}
