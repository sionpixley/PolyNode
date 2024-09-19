package services

import "os/exec"

func current() (string, error) {
	content, err := exec.Command("polyn", "current").Output()
	return string(content), err
}

func version() (string, error) {
	content, err := exec.Command("polyn", "version").Output()
	return string(content), err
}
