package res

import (
	"encoding/json"
	"net/http"
)

// Response wrapper
type Response struct {
	w http.ResponseWriter
}

// ResponseData for json response data
type responseData struct {
	Data interface{} `json:"data"`
}

// ResponseError for json response success
type responseError struct {
	Message string `json:"message"`
}

// JSON response
func (res *Response) JSON(data interface{}, statusCode int) {
	result, err := json.Marshal(Data(data))
	if err != nil {
		res.Fail(err.Error(), http.StatusInternalServerError)
		return
	}

	res.w.Header().Set("Content-Type", "application/json")
	res.w.WriteHeader(statusCode)
	res.w.Write(result)
}

// Fail response with json format
func (res *Response) Fail(msg string, statusCode int) {
	result, _ := json.Marshal(Data(Error(msg)))

	res.w.Header().Set("Content-Type", "application/json")
	res.w.WriteHeader(statusCode)
	res.w.Write(result)
}

// Data wrapper
func Data(data interface{}) interface{} {
	return responseData{data}
}

// Error wrapper
func Error(data string) interface{} {
	return responseError{data}
}

// NewResponse instance
func NewResponse(w http.ResponseWriter) *Response {
	return &Response{w}
}
