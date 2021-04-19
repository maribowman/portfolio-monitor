package controller

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

const (
	coinTickerRegexString = "^[a-zA-Z]{3}-(?i:EUR|USD)$"
)

var (
	coinTickerRegex = regexp.MustCompile(coinTickerRegexString)
)

func CoinTickerValidator(fl validator.FieldLevel) bool {
	return validateCoinTicker(fl.Field().String())
}

func validateCoinTicker(dto string) bool {
	return coinTickerRegex.MatchString(dto)
}
