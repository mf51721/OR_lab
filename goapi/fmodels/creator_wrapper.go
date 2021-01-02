package fmodels

type CreatorWrapper struct {
	Status string `json:"status"`

	Message string `json:"message"`

	Response Creator `json:"response"`
}
