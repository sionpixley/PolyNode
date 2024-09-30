//go:build !windows && gui

package utilities

import (
	"os/exec"
)

func CopyUpgradableFiles(home string) error {
	err := exec.Command("cp", "./PolyNode/polyn", home+"/.PolyNode/polyn").Run()
	if err != nil {
		return err
	}

	err = exec.Command("cp", "./PolyNode/PolyNode", home+"/.PolyNode/PolyNode").Run()
	if err != nil {
		return err
	}

	err = exec.Command("cp", "./PolyNode/LICENSE", home+"/.PolyNode/LICENSE").Run()
	if err != nil {
		return err
	}

	err = exec.Command("cp", "./PolyNode/README.md", home+"/.PolyNode/README.md").Run()
	if err != nil {
		return err
	}

	err = exec.Command("cp", "./PolyNode/SECURITY.md", home+"/.PolyNode/SECURITY.md").Run()
	if err != nil {
		return err
	}

	err = exec.Command("cp", "-r", "./PolyNode/gui", home+"/.PolyNode/gui").Run()
	if err != nil {
		return err
	}

	return exec.Command("cp", "./PolyNode/uninstall/uninstall", home+"/.PolyNode/uninstall/uninstall").Run()
}
