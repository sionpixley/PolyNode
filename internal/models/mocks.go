package models

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

type ExecMock struct {
	TimesOutputCalled int
	TimesRunCalled    int
}

type GzipMock struct {
	TimesCloseCalled     int
	TimesNewReaderCalled int
}

type HTTPMock struct {
	TimesDoCalled         int
	TimesNewClientCalled  int
	TimesNewRequestCalled int
}

type IOMock struct {
	TimesCopyCalled    int
	TimesReadAllCalled int
}

type OSMockExist struct {
	TimesCreateCalled     int
	TimesExitCalled       int
	TimesIsNotExistCalled int
	TimesLinkCalled       int
	TimesMkdirAllCalled   int
	TimesOpenCalled       int
	TimesOpenFileCalled   int
	TimesReadFileCalled   int
	TimesReadDirCalled    int
	TimesRemoveAllCalled  int
	TimesStatCalled       int
	TimesStderrCalled     int
	TimesStdoutCalled     int
	TimesSymlinkCalled    int
	TimesWriteFileCalled  int
}

type OSMockNotExist struct {
	TimesCreateCalled     int
	TimesExitCalled       int
	TimesIsNotExistCalled int
	TimesLinkCalled       int
	TimesMkdirAllCalled   int
	TimesOpenCalled       int
	TimesOpenFileCalled   int
	TimesReadFileCalled   int
	TimesReadDirCalled    int
	TimesRemoveAllCalled  int
	TimesStatCalled       int
	TimesStderrCalled     int
	TimesStdoutCalled     int
	TimesSymlinkCalled    int
	TimesWriteFileCalled  int
}

type TarMock struct {
	TimesNewReaderCalled int
	TimesNextCalled      int
}

func (execWrapper *ExecMock) Output(_ *exec.Cmd) ([]byte, error) {
	execWrapper.TimesOutputCalled += 1
	return []byte{}, nil
}

func (execWrapper *ExecMock) Run(_ *exec.Cmd) error {
	execWrapper.TimesRunCalled += 1
	return nil
}

func (gzipWrapper *GzipMock) Close(_ *gzip.Reader) error {
	gzipWrapper.TimesCloseCalled += 1
	return nil
}

func (gzipWrapper *GzipMock) NewReader(_ io.Reader) (*gzip.Reader, error) {
	gzipWrapper.TimesNewReaderCalled += 1
	return nil, nil
}

func (httpWrapper *HTTPMock) Do(_ *http.Client, _ *http.Request) (*http.Response, error) {
	httpWrapper.TimesDoCalled += 1
	message := `{ "message": "ok" }`
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(strings.NewReader(message)),
	}
	return resp, nil
}

func (httpWrapper *HTTPMock) NewClient() *http.Client {
	httpWrapper.TimesNewClientCalled += 1
	return nil
}

func (httpWrapper *HTTPMock) NewRequest(_ string, _ string, _ io.Reader) (*http.Request, error) {
	httpWrapper.TimesNewRequestCalled += 1
	return nil, nil
}

func (ioWrapper *IOMock) Copy(_ io.Writer, _ io.Reader) (int64, error) {
	ioWrapper.TimesCopyCalled += 1
	return 0, nil
}

func (ioWrapper *IOMock) ReadAll(_ io.Reader) ([]byte, error) {
	ioWrapper.TimesReadAllCalled += 1
	return []byte{}, nil
}

func (osWrapper *OSMockExist) Create(_ string) (*os.File, error) {
	osWrapper.TimesCreateCalled += 1
	return &os.File{}, nil
}

func (osWrapper *OSMockNotExist) Create(_ string) (*os.File, error) {
	osWrapper.TimesCreateCalled += 1
	return nil, nil
}

func (osWrapper *OSMockExist) Exit(_ int) {
	osWrapper.TimesExitCalled += 1
}

func (osWrapper *OSMockNotExist) Exit(_ int) {
	osWrapper.TimesExitCalled += 1
}

func (osWrapper *OSMockExist) IsNotExist(_ error) bool {
	osWrapper.TimesIsNotExistCalled += 1
	return false
}

func (osWrapper *OSMockNotExist) IsNotExist(_ error) bool {
	osWrapper.TimesIsNotExistCalled += 1
	return true
}

func (osWrapper *OSMockExist) Link(_ string, _ string) error {
	osWrapper.TimesLinkCalled += 1
	return nil
}

func (osWrapper *OSMockNotExist) Link(_ string, _ string) error {
	osWrapper.TimesLinkCalled += 1
	return nil
}

func (osWrapper *OSMockExist) MkdirAll(_ string, _ os.FileMode) error {
	osWrapper.TimesMkdirAllCalled += 1
	return nil
}

func (osWrapper *OSMockNotExist) MkdirAll(_ string, _ os.FileMode) error {
	osWrapper.TimesMkdirAllCalled += 1
	return nil
}

func (osWrapper *OSMockExist) Open(_ string) (*os.File, error) {
	osWrapper.TimesOpenCalled += 1
	return &os.File{}, nil
}

func (osWrapper *OSMockNotExist) Open(_ string) (*os.File, error) {
	osWrapper.TimesOpenCalled += 1
	return nil, nil
}

func (osWrapper *OSMockExist) OpenFile(_ string, _ int, _ os.FileMode) (*os.File, error) {
	osWrapper.TimesOpenFileCalled += 1
	return nil, nil
}

func (osWrapper *OSMockNotExist) OpenFile(_ string, _ int, _ os.FileMode) (*os.File, error) {
	osWrapper.TimesOpenFileCalled += 1
	return nil, nil
}

func (osWrapper *OSMockExist) ReadFile(_ string) ([]byte, error) {
	osWrapper.TimesReadFileCalled += 1
	return []byte("2025-02-26T12:23:11.723Z"), nil
}

func (osWrapper *OSMockNotExist) ReadFile(_ string) ([]byte, error) {
	osWrapper.TimesReadFileCalled += 1
	return []byte{}, nil
}

func (osWrapper *OSMockExist) ReadDir(_ string) ([]os.DirEntry, error) {
	osWrapper.TimesReadDirCalled += 1
	return []os.DirEntry{}, nil
}

func (osWrapper *OSMockNotExist) ReadDir(_ string) ([]os.DirEntry, error) {
	osWrapper.TimesReadDirCalled += 1
	return []os.DirEntry{}, nil
}

func (osWrapper *OSMockExist) RemoveAll(_ string) error {
	osWrapper.TimesRemoveAllCalled += 1
	return nil
}

func (osWrapper *OSMockNotExist) RemoveAll(_ string) error {
	osWrapper.TimesRemoveAllCalled += 1
	return nil
}

func (osWrapper *OSMockExist) Stat(_ string) (os.FileInfo, error) {
	osWrapper.TimesStatCalled += 1
	return nil, nil
}

func (osWrapper *OSMockNotExist) Stat(_ string) (os.FileInfo, error) {
	osWrapper.TimesStatCalled += 1
	return nil, nil
}

func (osWrapper *OSMockExist) Stderr() *os.File {
	osWrapper.TimesStderrCalled += 1
	return nil
}

func (osWrapper *OSMockNotExist) Stderr() *os.File {
	osWrapper.TimesStderrCalled += 1
	return nil
}

func (osWrapper *OSMockExist) Stdout() *os.File {
	osWrapper.TimesStdoutCalled += 1
	return nil
}

func (osWrapper *OSMockNotExist) Stdout() *os.File {
	osWrapper.TimesStdoutCalled += 1
	return nil
}

func (osWrapper *OSMockExist) Symlink(_ string, _ string) error {
	osWrapper.TimesSymlinkCalled += 1
	return nil
}

func (osWrapper *OSMockNotExist) Symlink(_ string, _ string) error {
	osWrapper.TimesSymlinkCalled += 1
	return nil
}

func (osWrapper *OSMockExist) WriteFile(_ string, _ []byte, _ os.FileMode) error {
	osWrapper.TimesWriteFileCalled += 1
	return nil
}

func (osWrapper *OSMockNotExist) WriteFile(_ string, _ []byte, _ os.FileMode) error {
	osWrapper.TimesWriteFileCalled += 1
	return nil
}

func (tarWrapper *TarMock) NewReader(reader io.Reader) *tar.Reader {
	return tar.NewReader(reader)
}

func (tarWrapper *TarMock) Next(_ *tar.Reader) (*tar.Header, error) {
	tarWrapper.TimesNextCalled += 1
	return nil, io.EOF
}
