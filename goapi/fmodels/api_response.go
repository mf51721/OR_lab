package fmodels

type ApiResponse struct {
	Status string `json:"status,omitempty"`

	Message string `json:"message,omitempty"`

	Response *string `json:"response,omitempty"`
}

var (
	RespBadRequest = ApiResponse{
		Status:  "Failed",
		Message: "There was an error processing the request, please check request format",
	}
)

func BuildErrorResponse(status string, err error) ApiResponse {
	return ApiResponse{
		Status:  status,
		Message: err.Error(),
	}
}
