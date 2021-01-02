package fmodels

type LanguagesWrapper struct {
	Status string `json:"status"`

	Message string `json:"message"`

	// response with all of the languages in the system
	Response []Language `json:"response"`
}
