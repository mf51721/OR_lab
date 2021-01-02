package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetCreators - Returns all programming language creators
func GetCreators(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// GetLangCreators -
func GetLangCreators(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// SetLangCreators - change list of language creators
func SetLangCreators(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
