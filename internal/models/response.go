package models

import (
	"encoding/json"
	"net/http"
)

type Responder interface {
	WriteJSON(http.ResponseWriter, int) error
}

type SuccessResponse[T any] struct {
	StatusCode int    `json:"code"`
	Message    string `json:"message"`
	Data       T      `json:"data,omitempty"` // Optional data field
}

func (sr SuccessResponse[T]) WriteJSON(w http.ResponseWriter, code int) error {
	jsonData, err := json.Marshal(sr)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	_, err = w.Write(jsonData)
	return err
}

type ErrorResponse struct {
	StatusCode   int    `json:"code"`
	ErrorMessage string `json:"message"`
}

func (er ErrorResponse) WriteJSON(w http.ResponseWriter, code int) error {
	jsonData, err := json.Marshal(er)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	_, err = w.Write(jsonData)
	return err
}
