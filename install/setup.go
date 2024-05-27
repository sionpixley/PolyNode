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
		err = exec.Command("sudo", "cp", "-r", "PolyNode", "/opt").Run()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		err = exec.Command("sudo", "touch", "/etc/profile.d/polyn-path.sh").Run()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		err = exec.Command("bash", "-c", "echo export PATH=\\$PATH:/opt/PolyNode | sudo tee /etc/profile.d/polyn-path.sh").Run()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		err = exec.Command("sudo", "chmod", "+x", "/etc/profile.d/polyn-path.sh").Run()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		err = exec.Command("sudo", "/etc/profile.d/polyn-path.sh").Run()
	case "windows":
		err = exec.Command("cmd.exe", "/c", `cd "C:\\Program Files" && mkdir polyn`).Run()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		err = exec.Command("robocopy", ".\\polyn", "C:\\Program Files\\polyn").Run()
	default:
	}

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("The polyn command has been installed. Please close all open terminals.")
	}
}

func printOptionalLine(operatingSystem string) {
	if operatingSystem != "windows" {
		fmt.Println()
	}
}
