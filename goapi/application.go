package goapi

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ApiServer struct {
	*gin.Engine
	db *gorm.DB
}

func NewServer(db *gorm.DB, router *gin.Engine) ApiServer {
	return ApiServer{
		Engine: router,
		db:     db,
	}

}
