package common

//HTTPResponse common data response for handling REST data and status
type DataResponse struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Length  uint64      `json:"length,omitempty"`
}
