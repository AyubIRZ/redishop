package controller

type controller struct{}

// ApiResponse
type ApiResponse struct {
	Status  bool          `json:"status"`
	Message string        `json:"message"`
}

// ErrorResponse wraps the failed api response in a proper format
func ErrorResponse(message string) ApiResponse {
	resp := ApiResponse{
		Status:  false,
		Message: message,
	}

	return resp
}
