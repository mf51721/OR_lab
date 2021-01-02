package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetCreators - Returns all programming language creators
func (h *Handler) GetCreators(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// GetLangCreators - Returns creators of a desired programming language
func (h *Handler) GetLangCreators(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// SetLangCreators - change list of language creators
func (h *Handler) SetLangCreators(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
