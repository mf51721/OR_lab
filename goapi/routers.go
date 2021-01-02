package goapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Route is the information for every URI.
type Route struct {
	// Name is the name of this Route.
	Name string
	// Method is the string for the HTTP method. ex) GET, POST etc..
	Method string
	// Pattern is the pattern of the URI.
	Pattern string
	// HandlerFunc is the handler function of this route.
	HandlerFunc gin.HandlerFunc
}

// Routes is the list of the generated Route.
type Routes []Route

// NewRouter returns a new router.
func NewRouter() *gin.Engine {
	router := gin.Default()
	rg := router.Group("/api")
	for _, route := range routes {
		switch route.Method {
		case http.MethodGet:
			rg.GET(route.Pattern, route.HandlerFunc)
		case http.MethodPost:
			rg.POST(route.Pattern, route.HandlerFunc)
		case http.MethodPut:
			rg.PUT(route.Pattern, route.HandlerFunc)
		case http.MethodDelete:
			rg.DELETE(route.Pattern, route.HandlerFunc)
		}
	}

	return router
}

// Index is the index handler.
func Index(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}

var routes = Routes{
	{
		"Index",
		http.MethodGet,
		"/",
		Index,
	},

	{
		"GetCreators",
		http.MethodGet,
		"/creators",
		GetCreators,
	},

	{
		"GetLangCreators",
		http.MethodGet,
		"/languages/:languageId/creators",
		GetLangCreators,
	},

	{
		"SetLangCreators",
		http.MethodPut,
		"/languages/:languageId/creators",
		SetLangCreators,
	},

	{
		"AddLanguage",
		http.MethodPost,
		"/languages",
		AddLanguage,
	},

	{
		"DeleteLang",
		http.MethodDelete,
		"/languages/:languageId",
		DeleteLang,
	},

	{
		"GetLangById",
		http.MethodGet,
		"/languages/:languageId",
		GetLangById,
	},

	{
		"GetLangWiki",
		http.MethodGet,
		"/languages/:languageId/wikipedia",
		GetLangWiki,
	},

	{
		"LanguagesGet",
		http.MethodGet,
		"/languages",
		LanguagesGet,
	},

	{
		"UpdateLang",
		http.MethodPut,
		"/languages/:languageId",
		UpdateLangWithForm,
	},
}
