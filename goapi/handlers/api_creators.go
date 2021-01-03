package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

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

// GetLangCreators - Returns creators of a desired programming language
func (h *Handler) GetLangCreators(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// SetLangCreators - change list of language creators
func (h *Handler) SetLangCreators(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
