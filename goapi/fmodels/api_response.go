package fmodels

type ApiResponse struct {
	Status string `json:"status"`

	Message string `json:"message"`

	Response *string `json:"response"`
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
