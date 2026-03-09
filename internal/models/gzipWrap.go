package models

import (
	"compress/gzip"
	"io"
)

type GzipWrap struct{}

func (_ GzipWrap) Close(reader *gzip.Reader) error {
	return reader.Close()
}

func (_ GzipWrap) NewReader(reader io.Reader) (*gzip.Reader, error) {
	return gzip.NewReader(reader)
}
