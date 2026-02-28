package models

import "io"

type IOWrap struct{}

func (_ IOWrap) Copy(dst io.Writer, src io.Reader) (written int64, err error) {
	return io.Copy(dst, src)
}

func (_ IOWrap) ReadAll(r io.Reader) ([]byte, error) {
	return io.ReadAll(r)
}
