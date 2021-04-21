package service

import (
	"maribowman/portfolio-monitor/app/repository"
	"testing"
)

func TestProcessAsset(t *testing.T) {
	service := NewCoinbaseService(&Wiring{
		FinanceClient: repository.NewCoinbaseClient(),
		Messenger:     nil,
	})

	service.ProcessAsset("")
}
