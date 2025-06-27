package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	currentBinaryLocation := "."
	if strings.Contains(os.Args[0], "\\") {
		parts := strings.Split(os.Args[0], "\\")
		currentBinaryLocation = strings.Join(parts[:len(parts)-1], "\\")
	}

	home := os.Getenv("LOCALAPPDATA") + "\\Programs"

	var err error
	if oldVersionExists(home) {
		err = update(currentBinaryLocation, home)
	} else {
		err = install(currentBinaryLocation, home)
	}

	if err != nil {
		log.Fatalln(err.Error())
	} else {
		fmt.Println("The polyn command has been installed.")
		fmt.Println("Please close all open terminals.")
	}
}
