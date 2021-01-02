package openapi

type WikiResponse struct {

	Handle string `json:"handle,omitempty"`

	// list of links
	Links []Link `json:"links,omitempty"`
}
