package controller

import (
	"github.com/gin-gonic/gin"
	"maribowman/portfolio-monitor/app/controller/middleware"
	"maribowman/portfolio-monitor/app/model"
)

type Controller struct {
	coinbaseService model.FinanceService
}

type Wiring struct {
	Router          *gin.Engine
	CoinbaseService model.FinanceService
}

func NewController(wiring *Wiring) {
	controller := &Controller{
		coinbaseService: wiring.CoinbaseService,
	}

	wiring.Router.Use(middleware.AuthorizationMiddleware())
	wiring.Router.GET("/crypto/:coinTicker", controller.GetCrypto)
}
