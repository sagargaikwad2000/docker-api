package models

type Response struct {
	StatusCode int         `json:"statusCode"`
	Msg        string      `json:"msg,omitempty"`
	Err        string      `json:"err,omitempty"`
	Result     interface{} `json:"result,omitempty"`
}
