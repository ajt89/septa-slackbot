package septa

// TrainViewObject represents data sent back by SEPTA train view
type TrainViewObject struct {
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

// RequestStatus provides a return status of requests
type RequestStatus struct {
	ErrorMsg string
	Status   int
	Data     []byte
}

// TrainViewDecodeStatus provides a return status of decoding response bodies to json
type TrainViewDecodeStatus struct {
	ErrorMsg      string
	Status        int
	TrainViewData []TrainViewObject
}

// TrainNoStatus response format for GetTrainNo
type TrainNoStatus struct {
	ErrorMsg string
	Status   int
	Data     TrainNoData
}

// TrainNoData provides data return format
type TrainNoData struct {
	NextStop string
	Late     string
	Dest     string
}

// GetAllTrainNoStatus response format for GetAllTrainNos
type GetAllTrainNoStatus struct {
	ErrorMsg string
	Status   int
	Data     []string
}
