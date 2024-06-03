//go:build !windows

// This file only exists to prevent build errors when building for Linux and macOS.

package main

import "errors"

func addToWindowsPath(home string) error {
	return errors.New("control flow path should never occur")
}
