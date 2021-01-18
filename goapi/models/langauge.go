package models

import "gorm.io/gorm"

type Language struct {
	gorm.Model

	Name string `json:"name"`

	Year int32 `json:"year"`

	Wikipedia string `json:"wikipedia"`

	Imperative bool `json:"imperative,omitempty"`

	ObjectOriented bool `json:"objectOriented,omitempty"`

	Functional bool `json:"functional,omitempty"`

	Procedural bool `json:"procedural,omitempty"`

	Generic bool `json:"generic,omitempty"`

	Reflective bool `json:"reflective,omitempty"`

	Creators []Creator `gorm:"many2many:language_creators" json:"creators,omitempty"`

	Slika string `json:"slika"`
}
