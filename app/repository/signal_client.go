package repository

import "maribowman/portfolio-monitor/app/model"

type SignalClient struct {
	restClient *RestClient
}

func NewSignalClient() model.MessengerClient {
	return &SignalClient{
		restClient: NewRestClient(),
	}
}

func (client *SignalClient) Push(holding model.Holding, recipient string) error {
	return nil
}
