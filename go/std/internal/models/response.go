package models

type RespStatus string

const (
	StatusSuccess RespStatus = "success"
	StatusFail    RespStatus = "fail"
	StatusError   RespStatus = "error"
	StatusUnknown RespStatus = "unknown"
)

type Response struct {
	Status  RespStatus `json:"status"`
	Data    any        `json:"data,omitempty,omitzero"`
	Code    int        `json:"code,omitempty"`
	Message string     `json:"message,omitempty"`
}
