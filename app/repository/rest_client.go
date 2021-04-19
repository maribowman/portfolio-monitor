package repository

import (
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

func NewRestClient() *RestClient {
	return &RestClient{
		client: &http.Client{
			Timeout: config.Config.Client.Timeout * time.Second,
		},
	}
}

func (client *RestClient) getData(baseURL, apiPath, pathParams string, headers, queryParams map[string]interface{}, model interface{}) error {
	requestUrl := buildURL(baseURL, apiPath, pathParams, queryParams)
	request, reqErr := http.NewRequest(http.MethodGet, requestUrl, nil)
	if reqErr != nil {
		log.Println("create request: ", reqErr.Error())
		return reqErr
	}
	for key, value := range headers {
		request.Header.Add(key, fmt.Sprintf("%v", value))
	}
	response, clientErr := client.client.Do(request)
	if clientErr != nil {
		log.Println("rest request: ", clientErr)
		return clientErr
	}
	defer response.Body.Close()
	return readAndConvertData(response.Body, model)
}

func readAndConvertData(body io.ReadCloser, model interface{}) error {
	bytes, readErr := ioutil.ReadAll(body)
	if readErr != nil {
		log.Println("convert data: ", readErr)
		return readErr
	}
	if jsonErr := json.Unmarshal(bytes, model); jsonErr != nil {
		log.Println("unmarshal response: ", jsonErr)
		return jsonErr
	}
	return nil
}

func buildURL(baseURL, apiPath, pathParams string, queryParams map[string]interface{}) string {
	query := make(url.Values)
	for key, value := range queryParams {
		query.Add(key, fmt.Sprintf("%v", value))
	}
	requestUrl := url.URL{
		Scheme:   "https",
		Host:     baseURL,
		Path:     apiPath + pathParams,
		RawQuery: query.Encode(),
	}
	return requestUrl.String()
}
