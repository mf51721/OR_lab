package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/mf51721/OR_lab/goapi/fmodels"
	"github.com/mf51721/OR_lab/goapi/models"
)

// AddLanguage - Add a new programming language to the collection
func (h *Handler) AddLanguage(c *gin.Context) {
	req := models.Language{}
	err := c.Bind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmodels.RespBadRequest)
	}
	err = h.ls.Add(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmodels.BuildErrorResponse("unsuccessful POST", err))
	}
	c.Header("Location", "/api/languages/"+strconv.FormatUint(uint64(req.ID), 10))
	c.JSON(http.StatusCreated, fmodels.LanguagesWrapper{
		Status:   "Created",
		Message:  "Successfully added new language entry",
		Response: nil,
	})
}

// DeleteLang - Deletes a language entry
func (h *Handler) DeleteLang(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// GetLangById - Find language by Id
func (h *Handler) GetLangById(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// GetLangWiki - Find wikipedia handle string
func (h *Handler) GetLangWiki(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// GetLanguages -
func (h *Handler) GetLanguages(c *gin.Context) {
	query, err := h.ls.GetAll("")
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmodels.BuildErrorResponse("Error occurred", err))
		return
	}

	var fLangs []fmodels.Language
	if query != nil {
		for _, lang := range *query {
			fLangs = append(fLangs, fmodels.Language{}.FromModel(lang).SetLinks())
		}
	}

	c.JSON(http.StatusOK, fmodels.LanguagesWrapper{
		Status:   "OK",
		Message:  "Successfully queried all desired language entries",
		Response: fLangs,
	})
}

// UpdateLangWithForm - Updates a programming language entry using form data
func (h *Handler) UpdateLangWithForm(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
