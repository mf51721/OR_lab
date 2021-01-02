package fmodels

type WikiResponseWrapper struct {
	Status string `json:"status"`

	Message string `json:"message"`

	Response WikiResponse `json:"response"`
}
