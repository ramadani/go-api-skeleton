package handler

import (
	"encoding/json"
	"net/http"
)

// Response http
type Response struct{}

// JSON response
func (res *Response) JSON(w http.ResponseWriter, data interface{}, statusCode int) {
	result, err := json.Marshal(ResponseData{data})
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
	result, _ := json.Marshal(ResponseData{ResponseError{msg}})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(result)
}
