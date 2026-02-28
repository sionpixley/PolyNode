package main

import "os"

type osMockExist struct{}
type osMockNotExist struct{}

func (_ osMockExist) IsNotExist(_ error) bool {
	return false
}

func (_ osMockNotExist) IsNotExist(_ error) bool {
	return true
}

func (_ osMockExist) ReadFile(_ string) ([]byte, error) {
	return []byte("2025-02-26T12:23:11.723Z"), nil
}

func (_ osMockNotExist) ReadFile(_ string) ([]byte, error) {
	return []byte{}, nil
}

func (_ osMockExist) Stat(_ string) (os.FileInfo, error) {
	return nil, nil
}

func (_ osMockNotExist) Stat(_ string) (os.FileInfo, error) {
	return nil, nil
}
