package repository

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"maribowman/portfolio-monitor/app/model/mocks"
	"net/http"
	"testing"
)

type TestOuter struct {
	TestInner string `json:"inner"`
}

type TestStruct struct {
	TestString string    `json:"string"`
	TestInt    int       `json:"int"`
	TestFloat  float64   `json:"float"`
	TestBool   bool      `json:"bool"`
	TestOuter  TestOuter `json:"outer"`
	TestList   []string  `json:"list"`
}

func TestGetData(t *testing.T) {
	// given
	tables := []struct {
		requestDetails RequestDetails
		responseStatus int
		responseHeader map[string]string
		responseBody   string
		expected       TestStruct
		exception      error
	}{
		{
			requestDetails: RequestDetails{
				Protocol: "http",
				BaseUrl:  "localhost",
				Path:     "/api",
			},
			responseStatus: http.StatusInternalServerError,
			responseHeader: nil,
			responseBody:   "",
			expected:       TestStruct{},
			exception:      errors.New("failed request"),
		},
		{
			requestDetails: RequestDetails{
				Protocol: "https",
				BaseUrl:  "localhost",
				Path:     "/bacon",
				Query:    map[string]interface{}{"withEggs": true},
				Headers:  map[string]interface{}{"breakfast": true},
			},
			responseStatus: http.StatusOK,
			responseHeader: map[string]string{"coffee": "black"},
			responseBody:   `{"string":"string","int":1,"float":2.3456,"bool":false,"outer":{"inner":"inner"},"list":["item"]}`,
			expected: TestStruct{
				TestString: "string",
				TestInt:    1,
				TestFloat:  2.3456,
				TestBool:   false,
				TestOuter:  TestOuter{TestInner: "inner"},
				TestList:   []string{"item"},
			},
			exception: nil,
		},
	}

	for _, table := range tables {
		// and
		mocks.DoFunc = func(*http.Request) (*http.Response, error) {
			header := make(http.Header)
			for key, value := range table.responseHeader {
				header.Add(key, value)
			}
			return &http.Response{
				StatusCode: table.responseStatus,
				Header:     header,
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(table.responseBody))),
			}, table.exception
		}

		// when
		client := RestClient{
			client: new(mocks.MockClient),
		}
		var actual TestStruct
		reqErr := client.getData(table.requestDetails, &actual)

		// then
		assert.Equal(t, table.exception, reqErr)
		assert.Equal(t, table.expected, actual)

		// and
		assert.Equal(t, http.MethodGet, mocks.UsedMethod)
		assert.Equal(t, buildURL(table.requestDetails), mocks.UsedUrl.String())
		for key, value := range table.requestDetails.Headers {
			assert.Equal(t, mocks.UsedHeaders[key], fmt.Sprintf("%v", value))
		}
	}
}

func TestPostData(t *testing.T) {
	// given
	tables := []struct {
		requestDetails RequestDetails
		responseStatus int
		responseHeader map[string]string
		responseBody   string
		expected       TestStruct
		exception      error
	}{
		{
			requestDetails: RequestDetails{
				Protocol: "http",
				BaseUrl:  "localhost",
				Path:     "/api",
				Headers:  map[string]interface{}{"test": "exists"},
				Body: TestStruct{
					TestString: "string",
					TestInt:    1,
					TestFloat:  2.3456,
					TestBool:   false,
					TestOuter:  TestOuter{TestInner: "inner"},
					TestList:   []string{},
				},
			},
			responseStatus: http.StatusOK,
			responseHeader: nil,
			responseBody:   `{"string":"string","int":1,"float":2.3456,"bool":false,"outer":{"inner":"inner"},"list":["item"]}`,
			expected: TestStruct{
				TestString: "string",
				TestInt:    1,
				TestFloat:  2.3456,
				TestBool:   false,
				TestOuter:  TestOuter{TestInner: "inner"},
				TestList:   []string{"item"},
			},
			exception: nil,
		},
		{
			requestDetails: RequestDetails{
				Protocol: "http",
				BaseUrl:  "localhost",
				Path:     "/api",
				Headers:  map[string]interface{}{"test": "exists"},
				Body: TestStruct{
					TestString: "string",
					TestInt:    1,
					TestFloat:  2.3456,
					TestBool:   false,
					TestOuter:  TestOuter{TestInner: "inner"},
					TestList:   []string{},
				},
			},
			responseStatus: http.StatusInternalServerError,
			responseHeader: nil,
			responseBody:   "",
			expected:       TestStruct{},
			exception:      errors.New("failed request"),
		},
	}
	for _, table := range tables {
		// and
		mocks.DoFunc = func(*http.Request) (*http.Response, error) {
			header := make(http.Header)
			for key, value := range table.responseHeader {
				header.Add(key, fmt.Sprintf("%v", value))
			}
			return &http.Response{
				StatusCode: table.responseStatus,
				Header:     header,
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(table.responseBody))),
			}, table.exception
		}

		// when
		client := RestClient{
			client: new(mocks.MockClient),
		}
		var actual TestStruct
		reqErr := client.postData(table.requestDetails, &actual)

		// then
		assert.Equal(t, table.exception, reqErr)
		assert.Equal(t, table.expected, actual)

		// and
		assert.Equal(t, http.MethodPost, mocks.UsedMethod)
		assert.Equal(t, buildURL(table.requestDetails), mocks.UsedUrl.String())
		for key, value := range table.requestDetails.Headers {
			assert.Equal(t, mocks.UsedHeaders[key], fmt.Sprintf("%v", value))
		}
		body, err := json.Marshal(table.requestDetails.Body)
		if err != nil {
			assert.Fail(t, err.Error())
		}
		assert.Equal(t, body, mocks.UsedBody)
	}
}

func TestBuildUrl(t *testing.T) {
	// given
	tables := []struct {
		details  RequestDetails
		expected string
	}{
		{
			details:  RequestDetails{Protocol: "https", BaseUrl: "url", Path: "path/test/234", Query: map[string]interface{}{"key": "value"}},
			expected: "https://url/path/test/234?key=value",
		},
		{
			details:  RequestDetails{Protocol: "https", BaseUrl: "url", Path: "path/test/123", Query: nil},
			expected: "https://url/path/test/123",
		},
		{
			details:  RequestDetails{Protocol: "http", BaseUrl: "url", Path: "path", Query: map[string]interface{}{"key0": 0, "key1": "value", "key2": "value"}},
			expected: "http://url/path?key0=0&key1=value&key2=value",
		},
		{
			details:  RequestDetails{Protocol: "https", BaseUrl: "url", Path: "path", Query: map[string]interface{}{"key0": "value", "key1": false, "key2": 5, "key3": true, "key4": "value"}},
			expected: "https://url/path?key0=value&key1=false&key2=5&key3=true&key4=value",
		},
	}

	for _, table := range tables {
		// when
		actual := buildURL(table.details)

		// then
		assert.Equal(t, table.expected, actual)
	}
}
