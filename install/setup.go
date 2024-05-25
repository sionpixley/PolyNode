package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

func main() {
	operatingSystem := runtime.GOOS

	defer printOptionalLine(operatingSystem)

	var err error
	switch operatingSystem {
	case "darwin":
	case "linux":
	case "windows":
		err = exec.Command("cmd.exe", "/C", "cd C:\\Program Files && mkdir polyn").Run()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		err = exec.Command("robocopy", ".\\polyn", "C:\\Program Files\\polyn").Run()
	default:
	}

	if err != nil {
		fmt.Println(err.Error())
	}
}

func printOptionalLine(operatingSystem string) {
	if operatingSystem != "windows" {
		fmt.Println()
	}
}
