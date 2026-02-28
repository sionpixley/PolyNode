package models

import "os"

type OSWrap struct{}

func (_ OSWrap) Create(name string) (*os.File, error) {
	return os.Create(name)
}

func (_ OSWrap) Exit(code int) {
	os.Exit(code)
}

func (_ OSWrap) IsNotExist(err error) bool {
	return os.IsNotExist(err)
}

func (_ OSWrap) Link(oldname string, newname string) error {
	return os.Link(oldname, newname)
}

func (_ OSWrap) MkdirAll(name string, perm os.FileMode) error {
	return os.MkdirAll(name, perm)
}

func (_ OSWrap) Open(name string) (*os.File, error) {
	return os.Open(name)
}

func (_ OSWrap) OpenFile(name string, flag int, perm os.FileMode) (*os.File, error) {
	return os.OpenFile(name, flag, perm)
}

func (_ OSWrap) ReadFile(filename string) ([]byte, error) {
	return os.ReadFile(filename)
}

func (_ OSWrap) ReadDir(dirname string) ([]os.DirEntry, error) {
	return os.ReadDir(dirname)
}

func (_ OSWrap) RemoveAll(name string) error {
	return os.RemoveAll(name)
}

func (_ OSWrap) Stat(filename string) (os.FileInfo, error) {
	return os.Stat(filename)
}

func (_ OSWrap) Stderr() *os.File {
	return os.Stderr
}

func (_ OSWrap) Stdout() *os.File {
	return os.Stdout
}

func (_ OSWrap) Symlink(oldname string, newname string) error {
	return os.Symlink(oldname, newname)
}

func (_ OSWrap) WriteFile(filename string, data []byte, perm os.FileMode) error {
	return os.WriteFile(filename, data, perm)
}
