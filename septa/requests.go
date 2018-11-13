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
	if err == nil {
		return bodyString
	} else {
		return "Error"
	}
}

// GetTrainNo retrieves data from septa based on train number
func GetTrainNo(trainNo string) TrainNoStatus {
	status := TrainNoStatus{}
	trainNoData := TrainNoData{}
	requestStatus := GetHandler("http://www3.septa.org/hackathon/TrainView/")
	if requestStatus.Status == 1 {
		status.ErrorMsg = requestStatus.ErrorMsg
		status.Status = requestStatus.Status
		return status
	}

	decodeStatus := TrainViewDecoder(requestStatus.Data)
	if decodeStatus.Status == 1 {
		status.ErrorMsg = decodeStatus.ErrorMsg
		status.Status = requestStatus.Status
		return status
	}

	trainViewArray := decodeStatus.TrainViewData

	for i := range trainViewArray {
		if trainViewArray[i].Trainno == trainNo {
			trainNoData.NextStop = trainViewArray[i].Nextstop
			trainNoData.Late = strconv.Itoa(trainViewArray[i].Late)
			trainNoData.Dest = trainViewArray[i].Dest
			status.Data = trainNoData
			status.Status = 0
			break
		}
	}

	if status.Data == (TrainNoData{}) {
		status.ErrorMsg = fmt.Sprintf("Train %s was not found", trainNo)
		status.Status = 1
	}

	return status
}

// GetAllTrainNos returns all train numbers and their destinations
func GetAllTrainNos() GetAllTrainNoStatus {
	status := GetAllTrainNoStatus{}
	requestStatus := GetHandler("http://www3.septa.org/hackathon/TrainView/")
	if requestStatus.Status == 1 {
		status.ErrorMsg = requestStatus.ErrorMsg
		status.Status = requestStatus.Status
		return status
	}

	decodeStatus := TrainViewDecoder(requestStatus.Data)
	if decodeStatus.Status == 1 {
		status.ErrorMsg = decodeStatus.ErrorMsg
		status.Status = requestStatus.Status
		return status
	}

	var trainNumberArray []string
	trainViewArray := decodeStatus.TrainViewData
	for i := range trainViewArray {
		indexValue := fmt.Sprintf("%s (%s)", trainViewArray[i].Trainno, trainViewArray[i].Dest)
		trainNumberArray = append(trainNumberArray, indexValue)
	}

	status.Data = trainNumberArray
	return status
}

// GetAllTrainsNextToArrive returns all trains numbers and destinations next to arrive at a station
func GetAllTrainsNextToArrive(stationName string) GetAllTrainNoStatus {
	status := GetAllTrainNoStatus{}
	requestStatus := GetHandler("http://www3.septa.org/hackathon/TrainView/")
	if requestStatus.Status == 1 {
		status.ErrorMsg = requestStatus.ErrorMsg
		status.Status = requestStatus.Status
		return status
	}

	decodeStatus := TrainViewDecoder(requestStatus.Data)
	if decodeStatus.Status == 1 {
		status.ErrorMsg = decodeStatus.ErrorMsg
		status.Status = requestStatus.Status
		return status
	}

	var trainNumberArray []string
	trainViewArray := decodeStatus.TrainViewData
	for i := range trainViewArray {
		if trainViewArray[i].Nextstop == stationName {
			indexValue := fmt.Sprintf("%s (%s)", trainViewArray[i].Trainno, trainViewArray[i].Dest)
			trainNumberArray = append(trainNumberArray, indexValue)
		}

	}

	status.Data = trainNumberArray
	return status
}
