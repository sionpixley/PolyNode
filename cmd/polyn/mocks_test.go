package main

import "os"

func isNotExistMockFalse(error) bool {
	return false
}

func isNotExistMockTrue(error) bool {
	return true
}

func readFileMockEmptyFile(string) ([]byte, error) {
	return []byte{}, nil
}

func readFileMockFakeData(string) ([]byte, error) {
	return []byte("2026-02-26T12:23:11.723Z"), nil
}

func statMock(string) (os.FileInfo, error) {
	return nil, nil
}
