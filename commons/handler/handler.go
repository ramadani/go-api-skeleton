package handler

// Handler helpers
type Handler struct {
	Response Response
}

// ResponseData for json response data
type ResponseData struct {
	Data interface{} `json:"data"`
}

// ResponseError for json response success
type ResponseError struct {
	Message string `json:"message"`
}
