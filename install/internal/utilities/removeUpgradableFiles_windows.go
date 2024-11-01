package utilities

import "os"

func RemoveUpgradableFiles(home string) error {
	err := os.RemoveAll(home + "\\PolyNode\\polyn.exe")
	if err != nil {
		return err
	}

	err = os.RemoveAll(home + "\\PolyNode\\uninstall\\uninstall.exe")
	if err != nil {
		return err
	}

	err = os.RemoveAll(home + "\\PolyNode\\PolyNode.exe")
	if err != nil {
		return err
	}

	err = os.RemoveAll(home + "\\PolyNode\\LICENSE")
	if err != nil {
		return err
	}

	err = os.RemoveAll(home + "\\PolyNode\\README.md")
	if err != nil {
		return err
	}

	err = os.RemoveAll(home + "\\PolyNode\\SECURITY.md")
	if err != nil {
		return err
	}

	err = os.RemoveAll(home + "\\PolyNode\\temp")
	if err != nil {
		return err
	}

	return os.RemoveAll(home + "\\PolyNode\\gui")
}
