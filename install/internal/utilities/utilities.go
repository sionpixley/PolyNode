//go:build !windows

package utilities

import (
	"os"
	"os/exec"
)

func CopyUpgradableFiles(currentBinaryLocation string, home string) error {
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

	return exec.Command("cp", currentBinaryLocation+"/PolyNode/uninstall/uninstall", home+"/.PolyNode/uninstall/uninstall").Run()
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
