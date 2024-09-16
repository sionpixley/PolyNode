package internal

import "os/exec"

func current() (string, error) {
	content, err := exec.Command("polyn", "current").Output()
	return string(content), err
}
