package models

import "os"

type OSWrapper interface {
	IsNotExist(error) bool
	ReadFile(string) ([]byte, error)
	Stat(string) (os.FileInfo, error)
}
