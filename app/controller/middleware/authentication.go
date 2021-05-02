package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

func AuthorizationMiddleware() gin.HandlerFunc {
	return authorized
}

func authorized(context *gin.Context) {
	key := context.GetHeader("key")
	signature := context.GetHeader("signature")
	timestamp := context.GetHeader("timestamp")
	if len(key) == 0 || len(signature) == 0 || len(timestamp) == 0 {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "authentication headers not complete"})
		return
	}
	if secret, ok := getSecretForKey(key); ok {
		timestampInteger, err := strconv.ParseInt(timestamp, 10, 64)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "invalid timestamp"})
			return
		}
		// request older than 30"
		log.Println(time.Unix(timestampInteger, 0).String())
		if time.Unix(timestampInteger, 0).Add(30 * time.Second).Before(time.Now().UTC()) {
			context.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "timestamp too old"})
			return
		} else if time.Unix(timestampInteger, 0).After(time.Now().UTC()) {
			context.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "timestamp in the future"})
			return
		}
		sigHash := hmac.New(sha256.New, []byte(secret))
		sigHash.Write([]byte(timestamp + context.Request.Method + context.Request.URL.Path))
		if hmac.Equal([]byte(signature), []byte(hex.EncodeToString(sigHash.Sum(nil)))) {
			context.Next()
		} else {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid signature"})
			return
		}
	} else {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid key"})
		return
	}
}

func getSecretForKey(key string) (string, bool) {
	if key != "test_key" {
		return "", false
	}
	return "test_secret", true
}
