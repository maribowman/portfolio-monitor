package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"maribowman/portfolio-monitor/app/config"
	"maribowman/portfolio-monitor/app/controller"
	"maribowman/portfolio-monitor/app/repository"
	"maribowman/portfolio-monitor/app/service"
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

	controller.NewController(&controller.Wiring{
		Router: router,
		CoinbaseService: service.NewCoinbaseService(&service.Wiring{
			FinanceClient: repository.NewCoinbaseClient(),
			Messenger:     nil,
		}),
	})

	return router
}
