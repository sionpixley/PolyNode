//go:build !windows

package utilities

import "os"

func RemoveUpgradableFiles(home string) error {
	err := os.RemoveAll(home + "/.PolyNode/polyn")
	if err != nil {
		return err
	}

	err = os.RemoveAll(home + "/.PolyNode/PolyNode")
	if err != nil {
		return err
	}

	err = os.RemoveAll(home + "/.PolyNode/LICENSE")
	if err != nil {
		return err
	}

	err = os.RemoveAll(home + "/.PolyNode/README.md")
	if err != nil {
		return err
	}

	err = os.RemoveAll(home + "/.PolyNode/SECURITY.md")
	if err != nil {
		return err
	}

	err = os.RemoveAll(home + "/.PolyNode/gui")
	if err != nil {
		return err
	}

	return os.RemoveAll(home + "/.PolyNode/uninstall/uninstall")
}
