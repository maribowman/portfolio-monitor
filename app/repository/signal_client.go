package repository

import (
	"maribowman/portfolio-monitor/app/config"
	"maribowman/portfolio-monitor/app/model"
)

type SignalClient struct {
	restClient       *RestClient
	dispatcherServer string
	linkedDevice     string
}

func NewSignalClient() model.MessengerClient {
	return &SignalClient{
		restClient:       NewRestClient(),
		dispatcherServer: config.Config.Signal.DispatcherServer,
		linkedDevice:     config.Config.Signal.LinkedDevice,
	}
}

func (client *SignalClient) Push(holding model.Holding, message model.Message) error {
	var response string
	details := RequestDetails{
		Protocol: "http",
		BaseUrl:  client.dispatcherServer,
		Path:     "/v2/send",
		Query:    nil,
		Headers:  nil,
		Body:     message,
	}
	client.restClient.postData(details, &response)
	return nil
}
