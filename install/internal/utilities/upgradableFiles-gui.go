//go:build !windows && gui

package utilities

import (
	"os"
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

func RemoveUpgradableFiles(home string) error {
	err := os.RemoveAll(home + "/.PolyNode/polyn")
	if err != nil {
		return err
	}

	err = os.RemoveAll(home + "/.PolyNode/PolyNode")
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

	err = os.RemoveAll(home + "/.PolyNode/gui")
	if err != nil {
		return err
	}

	return os.RemoveAll(home + "/.PolyNode/uninstall/uninstall")
}
