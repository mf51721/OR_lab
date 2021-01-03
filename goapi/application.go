package goapi

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/mf51721/OR_lab/goapi/services"
)

type ApiServer struct {
	*gin.Engine
	db *gorm.DB
	ls services.LanguageService
	cs services.CreatorService
}

func NewServer(db *gorm.DB, router *gin.Engine) ApiServer {
	return ApiServer{
		Engine: router,
		db:     db,
	}

}

func (s *ApiServer) SetLanguageService(ls services.LanguageService) {
	s.ls = ls
}

func (s *ApiServer) SetCreatorService(cs services.CreatorService) {
	s.cs = cs
}
