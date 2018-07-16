package septa

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

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
	trainNoStatus := TrainNoStatus{}
	trainNoData := TrainNoData{}
	requestStatus := GetHandler("http://www3.septa.org/hackathon/TrainView/")
	if requestStatus.Status == 1 {
		trainNoStatus.ErrorMsg = requestStatus.ErrorMsg
		trainNoStatus.Status = requestStatus.Status
		return trainNoStatus
	}

	decodeStatus := TrainViewDecoder(requestStatus.Data)
	if decodeStatus.Status == 1 {
		trainNoStatus.ErrorMsg = decodeStatus.ErrorMsg
		trainNoStatus.Status = requestStatus.Status
	}

	trainViewObject := decodeStatus.TrainViewData

	for i := range trainViewObject {
		if trainViewObject[i].Trainno == trainNo {
			trainNoData.NextStop = trainViewObject[i].Nextstop
			trainNoData.Late = strconv.Itoa(trainViewObject[i].Late)
			trainNoData.Dest = trainViewObject[i].Dest
			trainNoStatus.Data = trainNoData
			trainNoStatus.Status = 0
			break
		}
	}

	if trainNoStatus.Data == (TrainNoData{}) {
		trainNoStatus.ErrorMsg = fmt.Sprintf("Train %s was not found", trainNo)
		trainNoStatus.Status = 1
	}

	return trainNoStatus
}
