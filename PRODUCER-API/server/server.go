package server

import (
	"message-queue/handlers"

	"github.com/gin-gonic/gin"
)

func InitRouter(productHandler handlers.ProductHandler) *gin.Engine {

	router := gin.New()

	// Set logger middleware
	router.Use(gin.Logger())

	// Init routes
	router = InitRoutes(router, productHandler)
	return router

}
