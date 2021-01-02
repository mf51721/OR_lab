package fmodels

type LanguageWrapper struct {

	Status string `json:"status,omitempty"`

	Message string `json:"message,omitempty"`

	Response Language `json:"response,omitempty"`
}
