package goapi

import (
	"net/http"

	"github.com/mf51721/OR_lab/goapi/handlers"

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
		handlers.GetCreators,
	},

	{
		"GetLangCreators",
		http.MethodGet,
		"/languages/:languageId/creators",
		handlers.GetLangCreators,
	},

	{
		"SetLangCreators",
		http.MethodPut,
		"/languages/:languageId/creators",
		handlers.SetLangCreators,
	},

	{
		"AddLanguage",
		http.MethodPost,
		"/languages",
		handlers.AddLanguage,
	},

	{
		"DeleteLang",
		http.MethodDelete,
		"/languages/:languageId",
		handlers.DeleteLang,
	},

	{
		"GetLangById",
		http.MethodGet,
		"/languages/:languageId",
		handlers.GetLangById,
	},

	{
		"GetLangWiki",
		http.MethodGet,
		"/languages/:languageId/wikipedia",
		handlers.GetLangWiki,
	},

	{
		"LanguagesGet",
		http.MethodGet,
		"/languages",
		handlers.LanguagesGet,
	},

	{
		"UpdateLang",
		http.MethodPut,
		"/languages/:languageId",
		handlers.UpdateLangWithForm,
	},
}
