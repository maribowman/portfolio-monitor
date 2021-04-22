package mocks

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type MockClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	UsedMethod = req.Method
	UsedUrl = req.URL
	UsedHeaders = make(map[string]string)
	for key, _ := range req.Header {
		UsedHeaders[strings.ToLower(key)] = fmt.Sprintf("%v", req.Header.Get(key))
	}
	if req.Body != nil {
		var err error
		if UsedBody, err = ioutil.ReadAll(req.Body); err != nil {
			return nil, err
		}
	}
	return DoFunc(req)
}

var (
	UsedMethod  string
	UsedUrl     *url.URL
	UsedHeaders map[string]string
	UsedBody    []byte
	DoFunc      func(req *http.Request) (*http.Response, error)
)
