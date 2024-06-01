package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	operatingSystem := runtime.GOOS

	defer printOptionalLine(operatingSystem)

	var err error
	switch operatingSystem {
	case "darwin":
		err = installMac()
	case "linux":
		err = installLinux()
	case "windows":
		err = installWindows()
	default:
		err = errors.New("unsupported operating system")
	}

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("The polyn command has been installed.")
		fmt.Println("Please close all open terminals.")
	}
}

func installLinux() error {
	home := os.Getenv("HOME")
	err := exec.Command("cp", "-r", "PolyNode", home+"/.PolyNode").Run()
	if err != nil {
		return err
	}

	contentData, err := os.ReadFile(home + "/.bashrc")
	if err != nil {
		return err
	}

	content := string(contentData)
	content += "\nPATH=$PATH:" + home + "/.PolyNode:" + home + "/.PolyNode/nodejs/bin"

	err = os.WriteFile(home+"/.bashrc", []byte(content), 0644)
	return err
}

func installMac() error {
	err := exec.Command("sudo", "cp", "-r", "PolyNode", "/opt").Run()
	if err != nil {
		return err
	}
	return err
}

func installWindows() error {
	return nil
}

func printOptionalLine(operatingSystem string) {
	if operatingSystem != "windows" {
		fmt.Println()
	}
}
