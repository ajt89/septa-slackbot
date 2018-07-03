package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var netClient = &http.Client{
	Timeout: time.Second * 60,
}

type Dict struct {
	Value string
}

func getJson(url string, target interface{}) error {
	r, err := netClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
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

func GetTrainNo() string {
	data := Dict{}
	getJson("http://www3.septa.org/hackathon/TrainView/", &data)
	return string(data.Value)
}
