package fmodels

type ApiResponse struct {

	Status string `json:"status,omitempty"`

	Message string `json:"message,omitempty"`

	Response *string `json:"response,omitempty"`
}
