package goapi

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/mf51721/OR_lab/goapi/handlers"
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
	return router
}

// Index is the index handler.
func Index(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}

func (s *ApiServer) ConfigureDefaults() {
	handler := handlers.NewHandler(s.ls, s.cs)
	s.NoMethod(handler.NoMethod)
	s.NoRoute(handler.NoRoute)
}

func (s *ApiServer) GetRoutes() Routes {
	handler := handlers.NewHandler(s.ls, s.cs)

	return Routes{
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
			handler.GetCreators,
		},
		{
			"GetCreator",
			http.MethodGet,
			"/creators/:creatorId",
			handler.GetCreatorById,
		},
		{
			"GetLangCreators",
			http.MethodGet,
			"/languages/:languageId/creators",
			handler.GetLangCreators,
		},

		{
			"SetLangCreators",
			http.MethodPut,
			"/languages/:languageId/creators",
			handler.SetLangCreators,
		},

		{
			"AddLanguage",
			http.MethodPost,
			"/languages",
			handler.AddLanguage,
		},

		{
			"DeleteLang",
			http.MethodDelete,
			"/languages/:languageId",
			handler.DeleteLang,
		},

		{
			"GetLangById",
			http.MethodGet,
			"/languages/:languageId",
			handler.GetLangById,
		},

		{
			"GetLangWiki",
			http.MethodGet,
			"/languages/:languageId/wikipedia",
			handler.GetLangWiki,
		},

		{
			"LanguagesGet",
			http.MethodGet,
			"/languages",
			handler.GetLanguages,
		},

		{
			"UpdateLang",
			http.MethodPut,
			"/languages/:languageId",
			handler.UpdateLangWithForm,
		},
	}

}

func (s *ApiServer) SetRoutes(r Routes) {
	rg := s.Group("/api")

	for _, route := range r {
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
}
