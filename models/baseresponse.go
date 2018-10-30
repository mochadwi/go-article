package models

// BaseResponse views
type BaseResponse struct {
	RequestID string      `json:"requestID"`
	Now       int64       `json:"now"`
	Code      int       `json:"code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
}