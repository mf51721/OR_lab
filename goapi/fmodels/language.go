package fmodels

import (
	"strconv"

	"github.com/mf51721/OR_lab/goapi/models"
)

type Language struct {
	Id int64 `json:"id,omitempty"`

	Name string `json:"name"`

	Year int32 `json:"year"`

	Wikipedia string `json:"wikipedia"`

	Imperative bool `json:"imperative"`

	ObjectOriented bool `json:"objectOriented"`

	Functional bool `json:"functional"`

	Procedural bool `json:"procedural"`

	Generic bool `json:"generic"`

	Reflective bool `json:"reflective"`

	Creators []Creator `json:"creators"`

	// list of links
	Links []Link `json:"links"`
}

func (l Language) FromModel(language models.Language) Language {
	l.Id = int64(language.ID)
	l.Name = language.Name
	l.Year = language.Year
	l.Wikipedia = language.Wikipedia
	l.Imperative = language.Imperative
	l.ObjectOriented = language.ObjectOriented
	l.Functional = language.Functional
	l.Procedural = language.Procedural
	l.Generic = language.Generic
	l.Reflective = language.Reflective
	var c []Creator
	for _, creator := range language.Creators {
		c = append(c, Creator{}.FromModel(creator))
	}
	l.Creators = c

	return l
}

func (l Language) SetLinks() Language {
	l.Links = []Link{
		{
			Href: "/api/languages/" + strconv.FormatInt(l.Id, 10),
			Rel:  "self",
			Type: "GET,PUT",
		},
		{
			Href: "/api/languages/" + strconv.FormatInt(l.Id, 10) + "/wikipedia",
			Rel:  "wikiPageHandle",
			Type: "GET",
		},
		{
			Href: "/api/languages/" + strconv.FormatInt(l.Id, 10) + "/creators",
			Rel:  "creators",
			Type: "GET,PUT",
		},
	}
	return l
}
