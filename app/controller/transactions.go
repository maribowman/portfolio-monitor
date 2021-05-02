package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

func (controller *Controller) PostTransaction(context *gin.Context) {
	dto := context.Param("type")
	if !strings.EqualFold("stocks", dto) || !strings.EqualFold("crypto", dto) {
		context.JSON(http.StatusBadRequest, gin.H{"error": "failed to validate coin ticker"})
		return
	}
	switch dto {
	case "stocks":
		log.Println("stocks")
	case "crypto":
		log.Println("crypto")
	}

	context.Status(http.StatusNoContent)
}
