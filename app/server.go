package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"maribowman/signal-transmitter/app/config"
	"maribowman/signal-transmitter/app/service"
	"net/http"
)

func InitServer() (*http.Server, error) {
	return &http.Server{
		Addr:    fmt.Sprintf(":%d", config.Config.Server.Port),
		Handler: setupRouter(),
	}, nil
}

func setupRouter() *gin.Engine {
	gin.SetMode(config.Config.Server.Mode)
	router := gin.Default()

	router.POST("/send", service.SendMessage)
	return router
}
