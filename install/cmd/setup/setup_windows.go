package main

import (
	"fmt"
	"install/internal/constants"
	"install/internal/utilities"
	"log"
	"os"
	"os/exec"
	"strings"

	"golang.org/x/sys/windows/registry"
)

func main() {
	defer fmt.Println()

	currentBinaryLocation := "."
	if strings.Contains(os.Args[0], "\\") {
		parts := strings.Split(os.Args[0], "\\")
		currentBinaryLocation = strings.Join(parts[:len(parts)-1], "\\")
	}

	home := os.Getenv("LOCALAPPDATA") + "\\Programs"

	var err error
	if oldVersionExists(home) {
		err = upgrade(currentBinaryLocation, home)
	} else {
		err = install(currentBinaryLocation, home)
	}

	if err != nil {
		log.Fatal(err.Error())
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

func upgrade(currentBinaryLocation string, home string) error {
	err := utilities.RemoveUpgradableFiles(home)
	if err != nil {
		return err
	}

	return utilities.CopyUpgradableFiles(currentBinaryLocation, home)
}
