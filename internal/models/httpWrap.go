package models

import (
	"io"
	"net/http"
	"time"
)

type HTTPWrap struct{}

func (_ HTTPWrap) Do(client *http.Client, request *http.Request) (*http.Response, error) {
	return client.Do(request)
}

func (_ HTTPWrap) NewClient() *http.Client {
	return &http.Client{Timeout: time.Second * 5}
}

func (_ HTTPWrap) NewRequest(method string, url string, body io.Reader) (*http.Request, error) {
	return http.NewRequest(method, url, body)
}
