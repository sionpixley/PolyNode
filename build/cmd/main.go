package main

import (
	"fmt"
	"os/exec"
)

func main() {

}

func printError(err error) {
	fmt.Println(err.Error())
}

func setGoArch(arch string) error {
	output, err := exec.Command("go", "env", "-w", "GOARCH="+arch).Output()
	if err != nil {
		return err
	}

	fmt.Print(string(output))
	return nil
}

func setGoOs(operatingSystem string) error {
	output, err := exec.Command("go", "env", "-w", "GOOS="+operatingSystem).Output()
	if err != nil {
		return err
	}

	fmt.Print(string(output))
	return nil
}
