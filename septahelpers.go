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

type SeptaObject struct {
	Lat          string
	Lon          string
	Trainno      string
	Service      string
	Dest         string
	Nextstop     string
	Line         string
	Consist      string
	Heading      float32
	Late         int
	SOURCE       string
	TRACK        string
	TRACK_CHANGE string
}

type Status struct {
	errorMsg string
	data     [3]string
	status   int
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

// GetTrainNo retrieves data from septa based on train number
func GetTrainNo(trainNo string) Status {
	fmt.Println(trainNo)
	funcStatus := Status{}
	response, err := netClient.Get("http://www3.septa.org/hackathon/TrainView/")
	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)

	if err != nil {
		funcStatus.errorMsg = "Error decoding response content"
		funcStatus.status = 1
		return funcStatus
	}

	var septaObjects []SeptaObject
	fmt.Println(string(responseBody))
	json.Unmarshal(responseBody, &septaObjects)
	fmt.Println(septaObjects)

	var nextStop string
	var late string
	var dest string
	for i := range septaObjects {
		if septaObjects[i].Trainno == trainNo {
			nextStop = septaObjects[i].Nextstop
			late = strconv.Itoa(septaObjects[i].Late)
			fmt.Println(late)
			dest = septaObjects[i].Dest
			funcStatus.data = [3]string{nextStop, late, dest}
			funcStatus.status = 0
			break
		}
	}

	return funcStatus
}
