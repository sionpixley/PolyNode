//go:build !windows && gui

package utilities

import (
	"os/exec"
)

func CopyUpgradableFiles(currentBinaryLocation string, home string) error {
	err := exec.Command("cp", currentBinaryLocation+"/PolyNode/polyn", home+"/.PolyNode/polyn").Run()
	if err != nil {
		return err
	}

	err = exec.Command("cp", currentBinaryLocation+"/PolyNode/PolyNode", home+"/.PolyNode/PolyNode").Run()
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

	err = exec.Command("cp", "-r", currentBinaryLocation+"/PolyNode/gui", home+"/.PolyNode/gui").Run()
	if err != nil {
		return err
	}

	return exec.Command("cp", currentBinaryLocation+"/PolyNode/uninstall/uninstall", home+"/.PolyNode/uninstall/uninstall").Run()
}
