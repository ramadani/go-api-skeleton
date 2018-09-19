package handler

import (
	"encoding/json"
	"net/http"
)

// Respose http
type Response struct{}

type responseError struct {
	Message string `json:"message"`
}

// JSON response
func (res *Response) JSON(w http.ResponseWriter, data interface{}, statusCode int) {
	result, err := json.Marshal(data)
	if err != nil {
		res.jsonError(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(result)
}

func (res *Response) jsonError(w http.ResponseWriter, statusCode int, err error) {
	result, _ := json.Marshal(responseError{err.Error()})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(result)
}
