package models

import "archive/zip"

type ZipWrap struct{}

func (_ ZipWrap) Close(r *zip.ReadCloser) error {
	return r.Close()
}

func (_ ZipWrap) File(r *zip.ReadCloser) []*zip.File {
	return r.File
}

func (_ ZipWrap) OpenReader(path string) (*zip.ReadCloser, error) {
	return zip.OpenReader(path)
}
