package handlers

import (
	"net/http"
	"strconv"

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
	c.JSON(http.StatusOK, gin.H{})
}
