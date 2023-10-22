package server

import (
	"message-queue/handlers"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine, productHandler handlers.ProductHandler) *gin.Engine {
	router.POST("/add-product", productHandler.AddProduct)
	return router
}
