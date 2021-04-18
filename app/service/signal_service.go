package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"maribowman/signal-transmitter/app/model"
	"net/http"
)

func SendMessage(context *gin.Context) {
	var messageDTA model.Message
	if err := context.ShouldBindWith(&messageDTA, binding.Query); err != nil {
		var message string
		for index, e := range err.(validator.ValidationErrors) {
			if index != 0 {
				message += "\n"
			}
			message += fmt.Sprintf("failed to validate '%s' parameter", e.Field())
		}
		context.String(http.StatusBadRequest, message)
		return
	}


	context.Status(http.StatusNoContent)
}
