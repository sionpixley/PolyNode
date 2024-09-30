//go:build !gui

package utilities

import (
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
