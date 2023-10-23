package handlers

type ApiResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewApiError(message string) *ApiResponse {
	return &ApiResponse{
		Status:  "error",
		Message: message,
		Data:    nil,
	}
}

func NewApiResponse(data interface{}, message string) *ApiResponse {
	return &ApiResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	}
}
