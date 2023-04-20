package schema

import (
	"encoding/json"
	"log"
	"net/http"
)

type ApiOkResponse struct {
	Status string `json:"status"`
	Data   any    `json:"data"`
}

type APIErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func APIOk(w http.ResponseWriter, code int, data any) error {
	resp := &ApiOkResponse{
		Status: http.StatusText(code),
		Data:   data,
	}
	respBytes, err := json.Marshal(resp)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	_, err = w.Write(respBytes)
	return err
}

func APIError(w http.ResponseWriter, status int, err error) error {
	log.Printf("service error: %v", err)
	resp := &APIErrorResponse{
		Code:    http.StatusText(status),
		Message: err.Error(),
	}
	respBytes, err := json.Marshal(resp)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)

	_, err = w.Write(respBytes)

	return err
}
