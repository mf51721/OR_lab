package handlers

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

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
	rawLangId, ok := c.Params.Get("languageId")
	if !ok {
		c.JSON(http.StatusBadRequest, fmodels.RespBadRequest)
		return
	}
	langId, err := strconv.Atoi(rawLangId)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmodels.RespBadRequest)
		return
	}
	err = h.ls.Delete(uint(langId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmodels.BuildErrorResponse("Unsuccessful deletion", err))
		return
	}
	c.JSON(http.StatusOK, fmodels.ApiResponse{
		Status:  "Entry deleted",
		Message: fmt.Sprintf("Languge with ID %v was successfully deleted", langId),
	})
}

// GetLangById - Find language by Id
func (h *Handler) GetLangById(c *gin.Context) {
	rawLangId, ok := c.Params.Get("languageId")
	if !ok {
		c.JSON(http.StatusBadRequest, fmodels.RespBadRequest)
		return
	}
	langId, err := strconv.Atoi(rawLangId)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmodels.RespBadRequest)
		return
	}
	modLang, err := h.ls.Get(uint(langId))
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, fmodels.ApiResponse{
			Status:   "Not Found",
			Message:  "Requested language could not be found",
			Response: nil,
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmodels.BuildErrorResponse("Error returning language", err))
		return
	}

	c.JSON(http.StatusOK, fmodels.LanguageWrapper{
		Status:   "OK",
		Message:  "Found desired language",
		Response: fmodels.Language{}.FromModel(*modLang).SetLinks(),
	})
}

// GetLangWiki - Find wikipedia handle string
func (h *Handler) GetLangWiki(c *gin.Context) {
	rawLangId, ok := c.Params.Get("languageId")
	if !ok {
		c.JSON(http.StatusBadRequest, fmodels.RespBadRequest)
		return
	}
	langId, err := strconv.Atoi(rawLangId)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmodels.RespBadRequest)
		return
	}
	modLang, err := h.ls.Get(uint(langId))
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, fmodels.ApiResponse{
			Status:   "Not Found",
			Message:  "Requested language wiki handle could not be found",
			Response: nil,
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmodels.BuildErrorResponse("Error returning language wiki handle", err))
		return
	}

	c.JSON(http.StatusOK, fmodels.WikiResponseWrapper{
		Status:   "OK",
		Message:  "Found desired language wikipedia handle",
		Response: fmodels.WikiResponse{}.FromModel(*modLang).SetLinks(),
	})
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
	// Get request body
	var req gin.H
	err := c.Bind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmodels.RespBadRequest)
		return
	}

	// Get language ID from URL param
	rawLangId, ok := c.Params.Get("languageId")
	if !ok {
		c.JSON(http.StatusBadRequest, fmodels.RespBadRequest)
		return
	}
	langId, err := strconv.Atoi(rawLangId)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmodels.RespBadRequest)
		return
	}
	// TODO: find more elegant solution to disable changing language id
	delete(req, "id")
	delete(req, "ID")

	err = h.ls.Update(uint(langId), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmodels.BuildErrorResponse("Unsuccessful update", err))
		return
	}
	c.JSON(http.StatusOK, fmodels.ApiResponse{
		Status:  "Entry updated",
		Message: fmt.Sprintf("Languge with ID %v was successfully updated", langId),
	})
}

//GetLangPic -
func (h *Handler) GetLangPic(c *gin.Context) {
	rawLangId, ok := c.Params.Get("languageId")
	if !ok {
		c.JSON(http.StatusBadRequest, fmodels.RespBadRequest)
		return
	}
	langId, err := strconv.Atoi(rawLangId)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmodels.RespBadRequest)
		return
	}

	h.ls.GetPic(uint(langId))

	var fs http.FileSystem = http.Dir("")
	filename := "resources/" + fmt.Sprint(langId) + ".png"
	_, err = os.Stat(filename)

	if os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, fmodels.ApiResponse{
			Status:  "Image not found",
			Message: "The remote source does not have the picture specified",
		})
	} else {
		c.FileFromFS(filename, fs)
	}

}
