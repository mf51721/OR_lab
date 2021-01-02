package fmodels

import "github.com/mf51721/OR_lab/goapi/models"

type Creator struct {
	Id int64 `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	// list of links
	Links *[]Link `json:"links,omitempty"`
}

func (c Creator) FromModel(creator models.Creator) Creator {
	c.Id = int64(creator.ID)
	c.Name = creator.Name
	return c
}
