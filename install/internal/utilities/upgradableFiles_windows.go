//go:build !gui

package utilities

import (
	"os"
	"os/exec"
)

func CopyUpgradableFiles(home string) error {
	err := exec.Command("cmd", "/c", "copy", ".\\PolyNode\\polyn.exe", home+"\\PolyNode\\polyn.exe").Run()
	if err != nil {
		return err
	}

	err = exec.Command("cmd", "/c", "copy", ".\\PolyNode\\LICENSE", home+"\\PolyNode\\LICENSE").Run()
	if err != nil {
		return err
	}

	err = exec.Command("cmd", "/c", "copy", ".\\PolyNode\\README.md", home+"\\PolyNode\\README.md").Run()
	if err != nil {
		return err
	}

	err = exec.Command("cmd", "/c", "copy", ".\\PolyNode\\SECURITY.md", home+"\\PolyNode\\SECURITY.md").Run()
	if err != nil {
		return err
	}

	return exec.Command("cmd", "/c", "copy", ".\\PolyNode\\uninstall\\uninstall.exe", home+"\\PolyNode\\uninstall\\uninstall.exe").Run()
}

func RemoveUpgradableFiles(home string) error {
	err := os.RemoveAll(home + "\\PolyNode\\polyn.exe")
	if err != nil {
		return err
	}

	err = os.RemoveAll(home + "\\PolyNode\\uninstall\\uninstall.exe")
	if err != nil {
		return err
	}

	err = os.RemoveAll(home + "\\PolyNode\\LICENSE")
	if err != nil {
		return err
	}

	err = os.RemoveAll(home + "\\PolyNode\\README.md")
	if err != nil {
		return err
	}

	return os.RemoveAll(home + "\\PolyNode\\SECURITY.md")
}
