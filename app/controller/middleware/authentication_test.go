package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestAuthorized(t *testing.T) {
	// given
	tables := []struct {
		headers  map[string]interface{}
		expected struct {
			code    int
			message string
		}
	}{
		{
			headers: map[string]interface{}{
				"key":       "test_key",
				"signature": "",
				"timestamp": "1619175474",
			},
			expected: struct {
				code    int
				message string
			}{code: http.StatusBadRequest, message: `{"error":"authentication headers not complete"}`},
		},
		{
			headers: map[string]interface{}{
				"key":       "test_key",
				"signature": "dummy",
				"timestamp": time.Now().UTC(),
			},
			expected: struct {
				code    int
				message string
			}{code: http.StatusForbidden, message: `{"error":"invalid timestamp"}`},
		},
		{
			headers: map[string]interface{}{
				"key":       "test_key",
				"signature": "dummy",
				"timestamp": time.Now().Add(-24 * time.Hour).UTC().Unix(),
			},
			expected: struct {
				code    int
				message string
			}{code: http.StatusForbidden, message: `{"error":"timestamp too old"}`},
		},
		{
			headers: map[string]interface{}{
				"key":       "test_key",
				"signature": "dummy",
				"timestamp": time.Now().Add(24 * time.Hour).UTC().Unix(),
			},
			expected: struct {
				code    int
				message string
			}{code: http.StatusForbidden, message: `{"error":"timestamp in the future"}`},
		},
		// todo add successful test here
		{
			headers: map[string]interface{}{
				"key":       "invalid",
				"signature": "dummy",
				"timestamp": time.Now().UTC().Unix(),
			},
			expected: struct {
				code    int
				message string
			}{code: http.StatusUnauthorized, message: `{"error":"invalid key"}`},
		},
	}

	// and
	router := gin.New()
	router.Use(AuthorizationMiddleware())
	router.GET("/", func(context *gin.Context) {
		context.Status(http.StatusNoContent)
		return
	})

	for _, table := range tables {
		// when
		w := performRequest(router, http.MethodGet, "/", table.headers)

		// then
		assert.Equal(t, table.expected.code, w.Code)
		assert.Equal(t, table.expected.message, w.Body.String())
	}

}

func performRequest(r http.Handler, method, path string, headers map[string]interface{}) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, nil)
	for key, value := range headers {
		req.Header.Add(key, fmt.Sprintf("%v", value))
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
