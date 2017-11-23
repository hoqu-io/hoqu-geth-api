package models

type RestResponse struct {
    Data  interface{} `json:"data"`
    Error RestError   `json:"error"`
}

type RestError struct {
    Code    string      `json:"code"`
    Message interface{} `json:"message"`
}
