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

func (_ HTTPWrap) NewClient(config *PolyNodeConfig) *http.Client {
	return &http.Client{Timeout: time.Second * time.Duration(config.TimeoutInSeconds)}
}

func (_ HTTPWrap) NewRequest(method string, url string, body io.Reader) (*http.Request, error) {
	return http.NewRequest(method, url, body)
}
