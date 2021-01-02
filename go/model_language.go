package openapi

type Language struct {

	Id int64 `json:"id,omitempty"`

	Name string `json:"name"`

	Year int32 `json:"year"`

	Wikipedia string `json:"wikipedia"`

	Imperative bool `json:"imperative,omitempty"`

	ObjectOriented bool `json:"objectOriented,omitempty"`

	Functional bool `json:"functional,omitempty"`

	Procedural bool `json:"procedural,omitempty"`

	Generic bool `json:"generic,omitempty"`

	Reflective bool `json:"reflective,omitempty"`

	Creators []Creator `json:"creators,omitempty"`

	// list of links
	Links []Link `json:"links,omitempty"`
}
