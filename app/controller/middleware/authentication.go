package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"maribowman/portfolio-monitor/app/model"
	"net/http"
	"strconv"
	"time"
)

func authorized(context *gin.Context) (bool, model.Error) {
	key := context.GetHeader("key")
	signature := context.GetHeader("signature")
	timestamp := context.GetHeader("timestamp")
	if len(key) == 0 || len(signature) == 0 || len(timestamp) == 0 {
		return false, model.Error{
			StatusCode: http.StatusBadRequest,
			Message:    "authentication headers not complete",
		}
	}
	if secret, ok := getSecretForKey(key); ok {
		timestampInteger, err := strconv.ParseInt(timestamp, 10, 64)
		if err != nil {
			return false, model.Error{
				StatusCode: http.StatusForbidden,
				Message:    "invalid timestamp",
			}
		}
		// request older than 30"
		if time.Unix(timestampInteger, 0).Add(300000 * time.Second).Before(time.Now().UTC()) {
			return false, model.Error{
				StatusCode: http.StatusForbidden,
				Message:    "timestamp too old",
			}
		}
		sigHash := hmac.New(sha256.New, []byte(secret))
		sigHash.Write([]byte(timestamp + context.Request.Method + context.Request.URL.RawPath))
		return hmac.Equal([]byte(signature), []byte(hex.EncodeToString(sigHash.Sum(nil)))), model.Error{}
	}
	return false, model.Error{
		StatusCode: http.StatusUnauthorized,
		Message:    "invalid key",
	}
}

func getSecretForKey(key string) (string, bool) {
	if key != "test_key" {
		return "", false
	}
	return "test_secret", true
}
