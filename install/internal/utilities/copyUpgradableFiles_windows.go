package utilities

import (
	"os/exec"
)

func CopyUpgradableFiles(currentBinaryLocation string, home string) error {
	err := exec.Command("cmd", "/c", "copy", currentBinaryLocation+"\\PolyNode\\polyn.exe", home+"\\PolyNode\\polyn.exe").Run()
	if err != nil {
		return err
	}

	err = exec.Command("cmd", "/c", "copy", currentBinaryLocation+"\\PolyNode\\LICENSE", home+"\\PolyNode\\LICENSE").Run()
	if err != nil {
		return err
	}

	err = exec.Command("cmd", "/c", "copy", currentBinaryLocation+"\\PolyNode\\README.md", home+"\\PolyNode\\README.md").Run()
	if err != nil {
		return err
	}

	err = exec.Command("cmd", "/c", "copy", currentBinaryLocation+"\\PolyNode\\SECURITY.md", home+"\\PolyNode\\SECURITY.md").Run()
	if err != nil {
		return err
	}

	return exec.Command("cmd", "/c", "copy", currentBinaryLocation+"\\PolyNode\\uninstall\\uninstall.exe", home+"\\PolyNode\\uninstall\\uninstall.exe").Run()
}
