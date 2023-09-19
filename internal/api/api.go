package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jexodusmercado/blog-system-go/internal/conf"
	"gorm.io/gorm"
)

type API struct {
	db *gorm.DB
	config conf.Config
	handler *gin.Engine
}

func NewAPI(database *gorm.DB, globalConfig conf.Config, handler *gin.Engine) *API {
	// initialize api
	api := &API{db: database, config: globalConfig, handler: handler}

	return api
}