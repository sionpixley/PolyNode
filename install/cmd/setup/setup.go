//go:build !windows

package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
)

func main() {
	operatingSystem := runtime.GOOS

	defer fmt.Println()

	currentBinaryLocation := "."
	if strings.Contains(os.Args[0], "/") {
		parts := strings.Split(os.Args[0], "/")
		currentBinaryLocation = strings.Join(parts[:len(parts)-1], "/")
	}

	var err error
	switch operatingSystem {
	case "aix":
		fallthrough
	case "darwin":
		fallthrough
	case "linux":
		home := os.Getenv("HOME")
		if oldVersionExists(home) {
			err = upgrade(currentBinaryLocation, home)
		} else {
			err = install(currentBinaryLocation, home)
		}
	default:
		err = errors.New("setup: unsupported operating system")
	}

	if err != nil {
		log.Fatalln(err.Error())
	} else {
		fmt.Println("The polyn command has been installed.")
		fmt.Println("Please close all open terminals.")
	}
}
