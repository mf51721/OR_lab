package models

type WikiResponseWrapper struct {

	Status string `json:"status,omitempty"`

	Message string `json:"message,omitempty"`

	Response WikiResponse `json:"response,omitempty"`
}
