package main

import (
	"fmt"
	"os"

	"golang.org/x/sys/windows/registry"
)

func main() {
	err := uninstall()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func removePath(home string) error {
	key, err := registry.OpenKey(registry.CURRENT_USER, "Environment", registry.ALL_ACCESS)
	if err != nil {
		return err
	}
	defer key.Close()
}

func uninstall() error {
	home := os.Getenv("LOCALAPPDATA") + "\\Programs"
	err := removePath(home)
	if err != nil {
		return err
	}

}
