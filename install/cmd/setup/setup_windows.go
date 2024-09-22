package main

import (
	"fmt"
	"install/internal"
	"os"
	"os/exec"

	"golang.org/x/sys/windows/registry"
)

func main() {
	home := os.Getenv("LOCALAPPDATA") + "\\Programs"

	var err error
	if oldVersionExists(home) {
		err = upgrade(home)
	} else {
		err = install(home)
	}

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("The polyn command has been installed.")
		fmt.Println("Please close all open terminals.")
	}
}

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

func createPolynConfig(home string) error {
	return os.WriteFile(home+"\\PolyNode\\.polynrc", []byte(internal.DEFAULT_POLYNRC), 0644)
}

func install(home string) error {
	err := exec.Command("xcopy", "/s", "/i", ".\\PolyNode\\", home+"\\PolyNode\\").Run()
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
	_, err := os.Stat(home + "\\PolyNode")
	return !os.IsNotExist(err)
}

func upgrade(home string) error {
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

	err = os.RemoveAll(home + "\\PolyNode\\gui")
	if err != nil {
		return err
	}

	err = exec.Command("cmd", "/c", "copy", ".\\PolyNode\\polyn.exe", home+"\\PolyNode\\polyn.exe").Run()
	if err != nil {
		return err
	}

	err = exec.Command("cmd", "/c", "copy", ".\\PolyNode\\PolyNode.exe", home+"\\PolyNode\\PolyNode.exe").Run()
	if err != nil {
		return err
	}

	err = exec.Command("cmd", "/c", "xcopy", "/s", "/i", ".\\PolyNode\\gui\\", home+"\\PolyNode\\gui\\").Run()
	if err != nil {
		return err
	}

	return exec.Command("cmd", "/c", "copy", ".\\PolyNode\\uninstall\\uninstall.exe", home+"\\PolyNode\\uninstall\\uninstall.exe").Run()
}
