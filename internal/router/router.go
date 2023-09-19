package router

import "github.com/gin-gonic/gin"

type GRouter struct {}

func NewRouter() *GRouter {
	return &GRouter{}
}

func (r *GRouter) InitRouter() *gin.Engine {
	router := gin.Default()

	// set additional configuration for gin here
	// auth and cors should be here
	// r.Use(middleware)

	return router
}