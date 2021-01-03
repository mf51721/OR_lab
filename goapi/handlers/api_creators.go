package handlers

import (
	"net/http"
	"strconv"

	"github.com/mf51721/OR_lab/goapi/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/mf51721/OR_lab/goapi/fmodels"
)

// GetCreators - Returns all programming language creators
func (h *Handler) GetCreators(c *gin.Context) {
	query, err := h.cs.GetAll("")
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmodels.BuildErrorResponse("Error occurred", err))
		return
	}

	var fCreators []fmodels.Creator
	if query != nil {
		for _, creator := range *query {
			fCreators = append(fCreators, fmodels.Creator{}.FromModel(creator).SetLinks())
		}
	}

	c.JSON(http.StatusOK, fmodels.CreatorsWrapper{
		Status:   "OK",
		Message:  "Successfully queried all desired language creators",
		Response: fCreators,
	})
}

// GetCreatorById - Find language by Id
func (h *Handler) GetCreatorById(c *gin.Context) {
	rawCreatorId, ok := c.Params.Get("creatorId")
	if !ok {
		c.JSON(http.StatusBadRequest, fmodels.RespBadRequest)
		return
	}
	creatorId, err := strconv.Atoi(rawCreatorId)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmodels.RespBadRequest)
		return
	}
	modCreator, err := h.cs.Get(uint(creatorId))
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, fmodels.ApiResponse{
			Status:   "Not Found",
			Message:  "Requested language creator could not be found",
			Response: nil,
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmodels.BuildErrorResponse("Error returning creator", err))
		return
	}

	c.JSON(http.StatusOK, fmodels.CreatorWrapper{
		Status:   "OK",
		Message:  "Found desired language creator (author)",
		Response: fmodels.Creator{}.FromModel(*modCreator).SetLinks(),
	})
}

// GetLangCreators - Returns creators of a desired programming language
func (h *Handler) GetLangCreators(c *gin.Context) {
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
	modCreators, err := h.cs.GetByLanguage(uint(langId))
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, fmodels.ApiResponse{
			Status:   "Not Found",
			Message:  "Requested language creators could not be found",
			Response: nil,
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmodels.BuildErrorResponse("Error occurred", err))
		return
	}

	var fCreators []fmodels.Creator
	if modCreators != nil {
		for _, creator := range *modCreators {
			fCreators = append(fCreators, fmodels.Creator{}.FromModel(creator).SetLinks())
		}
	}

	c.JSON(http.StatusOK, fmodels.CreatorsWrapper{
		Status:   "OK",
		Message:  "Successfully queried all desired language creators",
		Response: fCreators,
	})
}

// SetLangCreators - change list of language creators
func (h *Handler) SetLangCreators(c *gin.Context) {
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
	var modCreators []models.Creator
	err = c.Bind(&modCreators)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmodels.RespBadRequest)
	}
	err = h.ls.SetCreators(uint(langId), modCreators)
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmodels.BuildErrorResponse("unsuccessful POST", err))
		return
	}
	c.Header("Location", "/api/languages/"+strconv.Itoa(langId)+"/creators")
	c.JSON(http.StatusCreated, fmodels.LanguagesWrapper{
		Status:   "Created",
		Message:  "Successfully set creators of desired language",
		Response: nil,
	})
}

// AddLangCreators - append to list of language creators
// This is currently not used, maybe it would found its use-case on POST /api/languages/:id/creators
func (h *Handler) AddLangCreators(c *gin.Context) {
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
			Message:  "language with id " + strconv.Itoa(langId) + " could not be found",
			Response: nil,
		})
		return
	}
	var modCreators []models.Creator
	err = c.Bind(&modCreators)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmodels.RespBadRequest)
	}
	for _, creator := range modCreators {
		creator.Languages = append(creator.Languages, *modLang)
		err = h.cs.Add(&creator)
		if err != nil {
			c.JSON(http.StatusInternalServerError, fmodels.BuildErrorResponse("unsuccessful POST", err))
			return
		}
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmodels.BuildErrorResponse("unsuccessful POST", err))
	}
	c.Header("Location", "/api/languages/"+strconv.Itoa(langId)+"/creators")
	c.JSON(http.StatusCreated, fmodels.LanguagesWrapper{
		Status:   "Created",
		Message:  "Successfully set creators of desired language",
		Response: nil,
	})
}
