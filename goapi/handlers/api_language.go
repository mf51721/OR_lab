package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AddLanguage - Add a new programming language to the collection
func AddLanguage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// DeleteLang - Deletes a language entry
func DeleteLang(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// GetLangById - Find language by Id
func GetLangById(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// GetLangWiki - Find wikipedia handle string
func GetLangWiki(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// LanguagesGet -
func LanguagesGet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// UpdateLangWithForm - Updates a programming language entry using form data
func UpdateLangWithForm(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
