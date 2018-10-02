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
type ResponseData struct {
	Data interface{} `json:"data"`
}

// ResponseMsg for json response success
type ResponseMsg struct {
	Message string `json:"message"`
}

// ResponseValidationError for json response success
type ResponseValidationError struct {
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
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
	result, _ := json.Marshal(Data(Message(msg)))

	res.w.Header().Set("Content-Type", "application/json")
	res.w.WriteHeader(statusCode)
	res.w.Write(result)
}

// ValidationError response with json format
func (res *Response) ValidationError(errors interface{}, statusCode int) {
	result, _ := json.Marshal(Data(ValidationError(errors)))

	res.w.Header().Set("Content-Type", "application/json")
	res.w.WriteHeader(statusCode)
	res.w.Write(result)
}

// Data wrapper
func Data(data interface{}) ResponseData {
	return ResponseData{data}
}

// Message wrapper
func Message(data string) ResponseMsg {
	return ResponseMsg{data}
}

// ValidationError wrapper
func ValidationError(errors interface{}) ResponseValidationError {
	return ResponseValidationError{
		Message: "validation_errors",
		Errors:  errors,
	}
}

// NewResponse instance
func NewResponse(w http.ResponseWriter) *Response {
	return &Response{w}
}
