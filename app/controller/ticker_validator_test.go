package controller

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"testing"
)

var validate *validator.Validate

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("coinTicker", CoinTickerValidator)
		validate = v
	}
}

func TestCoinTickerValidator(t *testing.T) {
	//// given
	//tables := []struct {
	//	start       string
	//	destination string
	//	isValid     bool
	//}{
	//	{"48.705705,9.102373", "48.749866,9.177561", true},
	//}
	//
	//for _, table := range tables {
	//	// when
	//	err := validate.Struct(model.CoordinateRequestParameters{
	//		Start: table.start,
	//		Dest:  table.destination,
	//	})
	//
	//	// then
	//	if table.isValid {
	//		assert.NoError(t, err)
	//	} else {
	//		assert.Error(t, err)
	//	}
	//}
}
