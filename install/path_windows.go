package main

import "golang.org/x/sys/windows/registry"

func addToWindowsPath(home string) error {
	key, err := registry.OpenKey(registry.CURRENT_USER, "Environment", registry.ALL_ACCESS)
	if err != nil {
		return err
	}
	defer key.Close()

	path, _, err := key.GetStringValue("Path")
	if err != nil {
		return err
	}
	path += ";" + home + "\\PolyNode;" + home + "\\PolyNode\\nodejs"

	err = key.SetStringValue("Path", path)
	return err
}
