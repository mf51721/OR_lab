package fmodels

type LanguageWrapper struct {
	Status string `json:"status"`

	Message string `json:"message"`

	Response Language `json:"response"`
}
