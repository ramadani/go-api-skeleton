package handler

import (
	"encoding/json"
	"net/http"
)

// Response http
type Response struct{}

type responseError struct {
	Message string `json:"message"`
}

// JSON response
func (res *Response) JSON(w http.ResponseWriter, data interface{}, statusCode int) {
	result, err := json.Marshal(data)
	if err != nil {
		res.Fail(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(result)
}

// Fail response with json format
func (res *Response) Fail(w http.ResponseWriter, msg string, statusCode int) {
	result, _ := json.Marshal(responseError{msg})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(result)
}
