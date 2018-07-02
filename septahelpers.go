package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var netClient = &http.Client{
	Timeout: time.Second * 60,
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
