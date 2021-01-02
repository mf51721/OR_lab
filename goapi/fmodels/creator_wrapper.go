package fmodels

type CreatorWrapper struct {

	Status string `json:"status,omitempty"`

	Message string `json:"message,omitempty"`

	Response Creator `json:"response,omitempty"`
}
