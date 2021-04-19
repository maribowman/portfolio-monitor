package repository

import (
	"bytes"
	"errors"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"maribowman/portfolio-monitor/app/model/mocks"
	"net/http"
	"testing"
)

func TestGetData(t *testing.T) {
	// given
	tables := []struct {
		status       int
		header       map[string]string
		responseJSON string
		expected     interface{}
		exception    error
	}{
		{
			http.StatusInternalServerError,
			nil,
			`{"string":"","int":1,"float":1.0,"bool":false}`,
			struct {
				TestString string  `json:"string"`
				TestInt    int     `json:"int"`
				TestFloat  float64 `json:"float"`
				TestBool   bool    `json:"bool"`
			}{
				TestString: "",
				TestInt:    1,
				TestFloat:  1.0,
				TestBool:   false,
			},
			errors.New("failed request"),
		},
		{
			http.StatusOK,
			map[string]string{
				"testHeaderKey": "testHeaderValue",
			},
			`{"string":"bacon","int":21515,"float":12.1561,"bool":true}`,
			struct {
				TestString string  `json:"string"`
				TestInt    int     `json:"int"`
				TestFloat  float64 `json:"float"`
				TestBool   bool    `json:"bool"`
			}{
				TestString: "bacon",
				TestInt:    21515,
				TestFloat:  12.1561,
				TestBool:   true,
			},
			nil,
		},
	}
	for _, table := range tables {
		// and
		mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
			header := make(http.Header)
			for key, value := range table.header {
				header.Add(key, value)
			}
			return &http.Response{
				StatusCode: table.status,
				Header:     header,
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(table.responseJSON))),
			}, table.exception
		}
		// and
		var actual struct {
			TestString string  `json:"string"`
			TestInt    int     `json:"int"`
			TestFloat  float64 `json:"float"`
			TestBool   bool    `json:"bool"`
		}

		// when
		client := RestClient{
			client: &mocks.MockClient{},
		}
		reqErr := client.getData("", "", "", nil, nil, &actual)

		// then
		if table.exception == nil {
			assert.Nil(t, reqErr)
			assert.Equal(t, table.expected, actual)
		} else {
			assert.Equal(t, table.exception, reqErr)
			assert.NotEqual(t, table.expected, actual)
		}
	}
}

func TestBuildUrl(t *testing.T) {
	// given
	tables := []struct {
		baseURL     string
		apiPath     string
		pathParams  string
		queryParams map[string]interface{}
		expected    string
	}{
		{"url", "path", "/test/234", map[string]interface{}{"key": "value"}, "https://url/path/test/234?key=value"},
		{"url", "path", "/test/123", nil, "https://url/path/test/123"},
		{"url", "path", "", map[string]interface{}{"key0": 0, "key1": "value", "key2": "value"}, "https://url/path?key0=0&key1=value&key2=value"},
		{"url", "path", "", map[string]interface{}{"key0": "value", "key1": false, "key2": 5, "key3": "value", "key4": "value"}, "https://url/path?key0=value&key1=false&key2=5&key3=value&key4=value"},
	}

	for _, table := range tables {
		// when
		actual := buildURL(table.baseURL, table.apiPath, table.pathParams, table.queryParams)
		// then
		assert.Equal(t, table.expected, actual)
	}
}
