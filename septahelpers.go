package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

// force timeout on requests
var netClient = &http.Client{
	Timeout: time.Second * 60,
}

// SeptaObject represents data sent back by SEPTA
type SeptaObject struct {
	Lat         string  `json:"lat"`
	Lon         string  `json:"lon"`
	Trainno     string  `json:"trainno"`
	Service     string  `json:"service"`
	Dest        string  `json:"dest"`
	Nextstop    string  `json:"nextstop"`
	Line        string  `json:"line"`
	Consist     string  `json:"consist"`
	Heading     float32 `json:"heading"`
	Late        int     `json:"late"`
	SOURCE      string  `json:"SOURCE"`
	TRACK       string  `json:"TRACK"`
	TRACKCHANGE string  `json:"TRACK_CHANGE"`
}

// TrainNoStatus allows functions to return status with error or data
type TrainNoStatus struct {
	ErrorMsg string
	Data     TrainNoData
	Status   int
}

// TrainNoData provides data return format
type TrainNoData struct {
	NextStop string
	Late     string
	Dest     string
}

// GetTrainView retrieves train view data from septa
func GetTrainView() string {
	response, err := netClient.Get("http://www3.septa.org/hackathon/TrainView/")
	defer response.Body.Close()
	bodyBytes, err := ioutil.ReadAll(response.Body)
	bodyString := string(bodyBytes)
	fmt.Printf(bodyString)
	if err == nil {
		return bodyString
	} else {
		return "Error"
	}
}

// GetTrainNo retrieves data from septa based on train number
func GetTrainNo(trainNo string) TrainNoStatus {
	funcStatus := TrainNoStatus{}
	funcData := TrainNoData{}
	response, requestError := netClient.Get("http://www3.septa.org/hackathon/TrainView/")

	if requestError != nil {
		funcStatus.ErrorMsg = "Error requesting data"
		funcStatus.Data = funcData
		funcStatus = 1
		return funcStatus
	}

	defer response.Body.Close()
	responseBody, responseError := ioutil.ReadAll(response.Body)

	if responseError != nil {
		funcStatus.ErrorMsg = "Error reading response body"
		funcStatus.Data = funcData
		funcStatus.Status = 1
		return funcStatus
	}

	var septaObjects []SeptaObject
	jsonDecodeError := json.Unmarshal(responseBody, &septaObjects)

	if jsonDecodeError != nil {
		funcStatus.ErrorMsg = "Error decoding response body"
	}

	for i := range septaObjects {
		if septaObjects[i].Trainno == trainNo {
			funcData.NextStop = septaObjects[i].Nextstop
			funcData.Late = strconv.Itoa(septaObjects[i].Late)
			funcData.Dest = septaObjects[i].Dest
			funcStatus.Data = funcData
			funcStatus.Status = 0
			break
		}
	}

	if funcStatus.Data == (TrainNoData{}) {
		funcStatus.ErrorMsg = fmt.Sprintf("Train %s was not found", trainNo)
		funcStatus.Data = funcData
		funcStatus.Status = 1
	}

	return funcStatus
}
