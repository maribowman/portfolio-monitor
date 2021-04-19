package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (controller *Controller) GetCrypto(context *gin.Context) {
	dto := context.Param("coinTicker")
	if !validateCoinTicker(dto) {
		context.JSON(http.StatusBadRequest, gin.H{"error": "failed to validate coin ticker"})
		return
	}
	price, err := controller.coinbaseService.ProcessAsset(dto)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, price)
}
