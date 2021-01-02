package fmodels

import (
	"strconv"

	"github.com/mf51721/OR_lab/goapi/models"
)

type Creator struct {
	Id int64 `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	// list of links
	Links []Link `json:"links,omitempty"`

	// Used to determine links
	languages []models.Language
}

func (c Creator) FromModel(creator models.Creator) Creator {
	c.Id = int64(creator.ID)
	c.Name = creator.Name
	c.languages = creator.Languages
	return c
}

func (c Creator) SetLinks() Creator {
	c.Links = []Link{
		{
			Href: "/api/creators/" + strconv.FormatInt(c.Id, 10),
			Rel:  "self",
			Type: "GET,PUT",
		},
	}
	for _, lang := range c.languages {
		c.Links = append(c.Links, Link{
			Href: "/api/languages/" + strconv.FormatUint(uint64(lang.ID), 10),
			Rel:  "CreatedLanguage",
			Type: "GET,PUT,DELETE",
		})
	}
	return c
}
