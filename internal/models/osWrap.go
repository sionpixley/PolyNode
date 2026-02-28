package models

import "os"

type OSWrap struct{}

func (_ OSWrap) IsNotExist(err error) bool {
	return os.IsNotExist(err)
}

func (_ OSWrap) ReadFile(filename string) ([]byte, error) {
	return os.ReadFile(filename)
}

func (_ OSWrap) Stat(filename string) (os.FileInfo, error) {
	return os.Stat(filename)
}
