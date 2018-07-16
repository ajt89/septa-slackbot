package septa

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

// force timeout on requests
var netClient = &http.Client{
	Timeout: time.Second * 60,
}

// GetHandler returns the response body of a get request
func GetHandler(url string) RequestStatus {
	requestStatus := RequestStatus{}
	response, requestError := netClient.Get(url)
	if requestError != nil {
		requestStatus.ErrorMsg = "Error requesting data"
		requestStatus.Status = 1
		return requestStatus
	}

	defer response.Body.Close()
	responseBody, responseError := ioutil.ReadAll(response.Body)

	if responseError != nil {
		requestStatus.ErrorMsg = "Error reading response body"
		requestStatus.Status = 1
		return requestStatus
	}

	requestStatus.Status = 0
	requestStatus.Data = responseBody

	return requestStatus
}

// TrainViewDecoder decodes the response body to the TrainViewObject type
func TrainViewDecoder(responseBody []byte) TrainViewDecodeStatus {
	trainViewDecodeStatus := TrainViewDecodeStatus{}
	var trainViewObject []TrainViewObject
	jsonDecodeError := json.Unmarshal(responseBody, &trainViewObject)

	if jsonDecodeError != nil {
		trainViewDecodeStatus.ErrorMsg = "Error decoding response body"
		trainViewDecodeStatus.Status = 1
		trainViewDecodeStatus.TrainViewData = trainViewObject
		return trainViewDecodeStatus
	}

	trainViewDecodeStatus.ErrorMsg = ""
	trainViewDecodeStatus.Status = 0
	trainViewDecodeStatus.TrainViewData = trainViewObject

	return trainViewDecodeStatus
}
