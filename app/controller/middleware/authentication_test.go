package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/url"
	"testing"
)

func TestAuthorized(t *testing.T) {
	// todo write more test cases
	// given
	headers := make(http.Header)
	headers.Add("key", "test_key")
	headers.Add("signature", "8e439e3f7117fad3f67c3a2cd701bac2b8e57adea5649f97c2dd1a98ea06f1cf")
	headers.Add("timestamp", "1619175474")

	context := gin.Context{
		Request: &http.Request{
			Method:     http.MethodGet,
			URL:        &url.URL{RawPath: "/api?test=true"},
			Header:     headers,
			RemoteAddr: "",
			RequestURI: "",
		},
	}

	// when
	isAuthorized, err := authorized(&context)

	log.Println("MESSAGE: " + err.Message)

	// then
	assert.Equal(t, true, isAuthorized)
	assert.NotNil(t, err)
}
