package models

import (
	"archive/tar"
	"io"
)

type TarWrap struct{}

func (_ TarWrap) NewReader(reader io.Reader) *tar.Reader {
	return tar.NewReader(reader)
}

func (_ TarWrap) Next(reader *tar.Reader) (*tar.Header, error) {
	return reader.Next()
}
