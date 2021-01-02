package openapi

type LanguagesWrapper struct {

	Status string `json:"status,omitempty"`

	Message string `json:"message,omitempty"`

	// response with all of the languages in the system
	Response []Language `json:"response,omitempty"`
}
