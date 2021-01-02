package handlers

import (
	"github.com/mf51721/OR_lab/goapi/services"
)

type Handler struct {
	ls services.LanguageService
}

// TODO: Handler should be an interface (service that is able to serve set methods, e.g. handleGetLanguages etc.).
func NewHandler(ls services.LanguageService) Handler {
	return Handler{ls: ls}
}
