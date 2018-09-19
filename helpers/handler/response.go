package handler

import (
	"encoding/json"
	"net/http"
)

// Respose http
type Respose struct{}

type responseError struct {
	Message string `json:"message"`
}

// JSON response
func (res *Respose) JSON(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")

	result, err := json.Marshal(data)
	if err != nil {
		resErr, _ := json.Marshal(responseError{err.Error()})
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(resErr)
		return
	}

	w.WriteHeader(statusCode)
	w.Write(result)
}
