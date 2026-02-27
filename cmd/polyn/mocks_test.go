package main

import "os"

func isNotExistMockFalse(_ error) bool {
	return false
}

func isNotExistMockTrue(_ error) bool {
	return true
}

func readFileMockEmptyFile(_ string) ([]byte, error) {
	return []byte{}, nil
}

func readFileMockFakeData(_ string) ([]byte, error) {
	return []byte("2026-02-26T12:23:11.723Z"), nil
}

func statMock(_ string) (os.FileInfo, error) {
	return nil, nil
}
