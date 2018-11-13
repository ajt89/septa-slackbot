package septa

import (
	"strings"
	"testing"
)

var tests struct {
	trainNo string
}

func TestGetTrainViewGood(t *testing.T) {
	response := GetTrainView()
	if response == "Error" {
		t.Fail()
	}
}

func TestGetAllTrainNosGood(t *testing.T) {
	status := GetAllTrainNos()
	if status.Status == 1 {
		t.Error()
	}
	tests.trainNo = strings.Split(status.Data[0], " ")[0]
}

func TestGetTrainNoGood(t *testing.T) {
	status := GetTrainNo(tests.trainNo)
	if status.Status == 1 {
		t.Fail()
	}
}

func TestGetTrainNoBad(t *testing.T) {
	status := GetTrainNo("nope")
	if status.Status != 1 {
		t.Fail()
	}
}
