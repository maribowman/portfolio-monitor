package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"maribowman/portfolio-monitor/app/config"
	"net/http"
	"net/url"
	"time"
)

type HttpClient interface {
	Do(request *http.Request) (*http.Response, error)
}

type RestClient struct {
	client HttpClient
}

type RequestDetails struct {
	Protocol string
	BaseUrl  string
	Path     string
	Query    map[string]interface{}
	Headers  map[string]interface{}
	Body     interface{}
}

func NewRestClient() *RestClient {
	return &RestClient{
		client: &http.Client{
			Timeout: config.Config.Client.Timeout * time.Second,
		},
	}
}

func (client *RestClient) getData(details RequestDetails, model interface{}) error {
	requestUrl := buildURL(details)
	request, reqErr := http.NewRequest(http.MethodGet, requestUrl, nil)
	if reqErr != nil {
		log.Println("create request: ", reqErr.Error())
		return reqErr
	}
	for key, value := range details.Headers {
		request.Header.Add(key, fmt.Sprintf("%v", value))
	}
	response, clientErr := client.client.Do(request)
	if clientErr != nil {
		log.Println("rest request: ", clientErr)
		return clientErr
	}
	defer response.Body.Close()
	return json.NewDecoder(response.Body).Decode(model)
}

func (client *RestClient) postData(details RequestDetails, model interface{}) error {
	requestUrl := buildURL(details)
	bodyBytes, marshalErr := json.Marshal(details.Body)
	if marshalErr != nil {
		log.Println("marshal body: ", marshalErr.Error())
		return marshalErr
	}
	request, reqErr := http.NewRequest(http.MethodPost, requestUrl, bytes.NewBuffer(bodyBytes))
	if reqErr != nil {
		log.Println("create request: ", reqErr.Error())
		return reqErr
	}
	for key, value := range details.Headers {
		request.Header.Add(key, fmt.Sprintf("%v", value))
	}
	response, clientErr := client.client.Do(request)
	if clientErr != nil {
		log.Println("rest request: ", clientErr)
		return clientErr
	}
	defer response.Body.Close()
	return json.NewDecoder(response.Body).Decode(model)
}

func readAndConvertData(body io.ReadCloser, model interface{}) error {
	responseBytes, readErr := ioutil.ReadAll(body)
	if readErr != nil {
		log.Println("convert data: ", readErr)
		return readErr
	}
	if jsonErr := json.Unmarshal(responseBytes, model); jsonErr != nil {
		log.Println("unmarshal response: ", jsonErr)
		return jsonErr
	}
	return nil
}

func buildURL(details RequestDetails) string {
	query := make(url.Values)
	for key, value := range details.Query {
		query.Add(key, fmt.Sprintf("%v", value))
	}
	requestUrl := url.URL{
		Scheme:   details.Protocol,
		Host:     details.BaseUrl,
		Path:     details.Path,
		RawQuery: query.Encode(),
	}
	return requestUrl.String()
}
