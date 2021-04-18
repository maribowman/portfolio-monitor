package model

import "github.com/gin-gonic/gin"

type Error struct {
	StatusCode int
	Message    string `json:"message"`
}

func returnError(context *gin.Context, err Error) {
	context.JSON(err.StatusCode, err.Message)
	return
}
