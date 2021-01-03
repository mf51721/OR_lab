package fmodels

import (
	"strconv"

	"github.com/mf51721/OR_lab/goapi/models"
)

type WikiResponse struct {
	Handle string `json:"handle"`

	// list of links
	Links []Link `json:"links"`

	// Used only for link creation, extremely memory inefficient, maybe store only the id?
	language models.Language
}

func (w WikiResponse) FromModel(language models.Language) WikiResponse {
	w.Handle = language.Wikipedia
	w.language = language
	return w
}

func (w WikiResponse) SetLinks() WikiResponse {
	w.Links = []Link{
		{
			Href: "/api/languages/" + strconv.FormatInt(int64(w.language.ID), 10) + "/wikipedia",
			Rel:  "self",
			Type: "GET",
		},
		{
			Href: "/api/languages/" + strconv.FormatInt(int64(w.language.ID), 10),
			Rel:  "language",
			Type: "GET,PUT",
		},
		{
			Href: "/api/languages/" + strconv.FormatInt(int64(w.language.ID), 10) + "/creators",
			Rel:  "creators",
			Type: "GET,PUT",
		},
	}
	return w
}
