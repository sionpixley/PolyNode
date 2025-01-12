package main

import (
	"install/internal/constants"
	"os"
	"os/exec"

	"golang.org/x/sys/windows/registry"
)

func addToPath(home string) error {
	key, err := registry.OpenKey(registry.CURRENT_USER, "Environment", registry.ALL_ACCESS)
	if err != nil {
		return err
	}
	defer key.Close()

	path, _, err := key.GetStringValue("Path")
	if err != nil {
		return err
	}

	path += ";" + home + "\\PolyNode;" + home + "\\PolyNode\\nodejs"
	return key.SetStringValue("Path", path)
}

func copyUpgradableFiles(currentBinaryLocation string, home string) error {
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

func createPolynConfig(home string) error {
	return os.WriteFile(home+"\\PolyNode\\polynrc.json", []byte(constants.DEFAULT_POLYNRC), 0644)
}

func install(currentBinaryLocation string, home string) error {
	err := exec.Command("cmd", "/c", "xcopy", "/s", "/i", currentBinaryLocation+"\\PolyNode\\", home+"\\PolyNode\\").Run()
	if err != nil {
		return err
	}

	err = createPolynConfig(home)
	if err != nil {
		return err
	}

	return addToPath(home)
}

func oldVersionExists(home string) bool {
	if _, err := os.Stat(home + "\\PolyNode"); os.IsNotExist(err) {
		return false
	} else if err != nil {
		return false
	} else {
		return true
	}
}

func removeUpgradableFiles(home string) error {
	err := os.RemoveAll(home + "\\PolyNode\\polyn.exe")
	if err != nil {
		return err
	}

	err = os.RemoveAll(home + "\\PolyNode\\uninstall\\uninstall.exe")
	if err != nil {
		return err
	}

	err = os.RemoveAll(home + "\\PolyNode\\PolyNode.exe")
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

	err = os.RemoveAll(home + "\\PolyNode\\SECURITY.md")
	if err != nil {
		return err
	}

	return os.RemoveAll(home + "\\PolyNode\\gui")
}

func upgrade(currentBinaryLocation string, home string) error {
	err := removeUpgradableFiles(home)
	if err != nil {
		return err
	}

	return copyUpgradableFiles(currentBinaryLocation, home)
}
