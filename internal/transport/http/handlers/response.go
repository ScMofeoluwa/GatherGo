package handlers

type ApiResponse struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Status  string      `json:"status"`
}

func NewApiError(message string) *ApiResponse {
	return &ApiResponse{
		Data:    nil,
		Message: message,
		Status:  "error",
	}
}

func NewApiResponse(data interface{}, message string) *ApiResponse {
	return &ApiResponse{
		Data:    data,
		Message: message,
		Status:  "success",
	}
}
