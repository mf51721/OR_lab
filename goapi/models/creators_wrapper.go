package models

type CreatorsWrapper struct {

	Status string `json:"status,omitempty"`

	Message string `json:"message,omitempty"`

	// response with all of the languages in the system
	Response []Creator `json:"response,omitempty"`
}
