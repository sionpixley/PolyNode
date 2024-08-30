package main

import (
	"fmt"
	"os"
	"os/exec"

	"golang.org/x/sys/windows/registry"
)

func main() {
	err := install()
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

}

func checkForOldVersion(home string) error {
	if _, err := os.Stat(home + "\\PolyNode"); os.IsNotExist(err) {
		return nil
	} else {
		return exec.Command(home + "\\PolyNode\\uninstall\\uninstall.exe").Run()
	}
}

func install() error {
	home := os.Getenv("LOCALAPPDATA") + "\\Programs"

	err := checkForOldVersion(home)
	if err != nil {
		return err
	}

	err = exec.Command()
}
