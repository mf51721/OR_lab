package models

type Creator struct {

	Id int64 `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	// list of links
	Links []Link `json:"links,omitempty"`
}
