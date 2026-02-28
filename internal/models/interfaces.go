package models

import (
	"io"
	"net/http"
	"os"
)

type HTTPWrapper interface {
	Do(*http.Client, *http.Request) (*http.Response, error)
	NewClient() *http.Client
	NewRequest(string, string, io.Reader) (*http.Request, error)
}

type IOWrapper interface {
	Copy(io.Writer, io.Reader) (int64, error)
	ReadAll(io.Reader) ([]byte, error)
}

type OSWrapper interface {
	Create(string) (*os.File, error)
	Exit(int)
	IsNotExist(error) bool
	Link(string, string) error
	MkdirAll(string, os.FileMode) error
	Open(string) (*os.File, error)
	OpenFile(string, int, os.FileMode) (*os.File, error)
	ReadFile(string) ([]byte, error)
	ReadDir(string) ([]os.DirEntry, error)
	RemoveAll(string) error
	Stat(string) (os.FileInfo, error)
	Stderr() *os.File
	Stdout() *os.File
	Symlink(string, string) error
	WriteFile(string, []byte, os.FileMode) error
}
