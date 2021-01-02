package fmodels

type WikiResponse struct {
	Handle string `json:"handle"`

	// list of links
	Links []Link `json:"links"`
}
