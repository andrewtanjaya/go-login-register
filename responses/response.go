package responses

import (
	"encoding/json"
	"net/http"
)

type BaseValueResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type BaseResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func Response(w http.ResponseWriter, code int, message string, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	var response any

	if payload != nil {
		response = &BaseValueResponse{
			Code:    code,
			Message: message,
			Data:    payload,
		}
	} else {
		response = &BaseResponse{
			Code:    code,
			Message: message,
		}
	}

	res, _ := json.Marshal(response)
	w.Write(res)
}
