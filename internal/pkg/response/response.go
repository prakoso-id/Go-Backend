package response

// Response is the standard API response structure
type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// SuccessResponse creates a success response with data
func SuccessResponse(message string, data interface{}) Response {
	return Response{
		Status:  "success",
		Message: message,
		Data:    data,
	}
}

// ErrorResponse creates an error response
func ErrorResponse(message string) Response {
	return Response{
		Status:  "error",
		Message: message,
		Data:    nil,
	}
}

// EmptySuccessResponse creates a success response without data
func EmptySuccessResponse(message string) Response {
	return Response{
		Status:  "success",
		Message: message,
		Data:    nil,
	}
}
