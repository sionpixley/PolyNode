package main

import (
	"io"
	"net/http"
	"os"
)

type httpMock struct{}
type ioMock struct{}
type osMockExist struct{}
type osMockNotExist struct{}

func (_ httpMock) Do(_ *http.Request) (*http.Response, error) {
	resp := &http.Response{StatusCode: http.StatusOK}
	return resp, nil
}

func (_ httpMock) NewClient() *http.Client {
	return nil
}

func (_ httpMock) NewRequest(_ string, _ string, _ io.Reader) (*http.Request, error) {
	return nil, nil
}

func (_ ioMock) Copy(_ io.Writer, _ io.Reader) (int64, error) {
	return 0, nil
}

func (_ osMockExist) Create(_ string) (*os.File, error) {
	return &os.File{}, nil
}

func (_ osMockNotExist) Create(_ string) (*os.File, error) {
	return nil, nil
}

func (_ osMockExist) Exit(_ int) {}

func (_ osMockNotExist) Exit(_ int) {}

func (_ osMockExist) IsNotExist(_ error) bool {
	return false
}

func (_ osMockNotExist) IsNotExist(_ error) bool {
	return true
}

func (_ osMockExist) Link(_ string, _ string) error {
	return nil
}

func (_ osMockNotExist) Link(_ string, _ string) error {
	return nil
}

func (_ osMockExist) MkdirAll(_ string, _ os.FileMode) error {
	return nil
}

func (_ osMockNotExist) MkdirAll(_ string, _ os.FileMode) error {
	return nil
}

func (_ osMockExist) Open(_ string) (*os.File, error) {
	return nil, nil
}

func (_ osMockNotExist) Open(_ string) (*os.File, error) {
	return nil, nil
}

func (_ osMockExist) OpenFile(_ string, _ int, _ os.FileMode) (*os.File, error) {
	return nil, nil
}

func (_ osMockNotExist) OpenFile(_ string, _ int, _ os.FileMode) (*os.File, error) {
	return nil, nil
}

func (_ osMockExist) ReadFile(_ string) ([]byte, error) {
	return []byte("2025-02-26T12:23:11.723Z"), nil
}

func (_ osMockNotExist) ReadFile(_ string) ([]byte, error) {
	return []byte{}, nil
}

func (_ osMockExist) ReadDir(_ string) ([]os.DirEntry, error) {
	return []os.DirEntry{}, nil
}

func (_ osMockNotExist) ReadDir(_ string) ([]os.DirEntry, error) {
	return []os.DirEntry{}, nil
}

func (_ osMockExist) RemoveAll(_ string) error {
	return nil
}

func (_ osMockNotExist) RemoveAll(_ string) error {
	return nil
}

func (_ osMockExist) Stat(_ string) (os.FileInfo, error) {
	return nil, nil
}

func (_ osMockNotExist) Stat(_ string) (os.FileInfo, error) {
	return nil, nil
}

func (_ osMockExist) Stderr() *os.File {
	return nil
}

func (_ osMockNotExist) Stderr() *os.File {
	return nil
}

func (_ osMockExist) Stdout() *os.File {
	return nil
}

func (_ osMockNotExist) Stdout() *os.File {
	return nil
}

func (_ osMockExist) Symlink(_ string, _ string) error {
	return nil
}

func (_ osMockNotExist) Symlink(_ string, _ string) error {
	return nil
}

func (_ osMockExist) WriteFile(_ string, _ []byte, _ os.FileMode) error {
	return nil
}

func (_ osMockNotExist) WriteFile(_ string, _ []byte, _ os.FileMode) error {
	return nil
}
