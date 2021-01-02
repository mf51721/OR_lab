package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mf51721/OR_lab/goapi/fmodels"
	"github.com/mf51721/OR_lab/goapi/services"
)

type Handler struct {
	ls services.LanguageService
}

// TODO: Handler should be an interface (service that is able to serve set methods, e.g. handleGetLanguages etc.).
func NewHandler(ls services.LanguageService) Handler {
	return Handler{ls: ls}
}

func (h *Handler) NoMethod(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, fmodels.ApiResponse{
		Status:  "Not implemented",
		Message: "Method not implemented for requested resource",
	})
}

func (h *Handler) NoRoute(c *gin.Context) {
	c.JSON(http.StatusNotFound, fmodels.ApiResponse{
		Status:   "Not Found",
		Message:  "Requested route does not exist",
		Response: nil,
	})
}
