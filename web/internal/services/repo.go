package services

import "os/exec"

func add(version string) (string, error) {
	content, err := exec.Command("polyn", "add", version).Output()
	return string(content), err
}

func list() (string, error) {
	content, err := exec.Command("polyn", "list").Output()
	return string(content), err
}

func remove(version string) (string, error) {
	content, err := exec.Command("polyn", "remove", version).Output()
	return string(content), err
}

func search() (string, error) {
	content, err := exec.Command("polyn", "search").Output()
	return string(content), err
}

func use(version string) (string, error) {
	content, err := exec.Command("polyn", "use", version).Output()
	return string(content), err
}

func version() (string, error) {
	content, err := exec.Command("polyn", "version").Output()
	return string(content), err
}
