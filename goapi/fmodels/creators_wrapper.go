package fmodels

type CreatorsWrapper struct {
	Status string `json:"status"`

	Message string `json:"message"`

	// response with all of the languages in the system
	Response []Creator `json:"response"`
}
