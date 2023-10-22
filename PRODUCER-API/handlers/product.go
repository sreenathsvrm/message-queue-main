package handlers

import (
	"fmt"
	"message-queue/internals/helpers"
	"message-queue/internals/models"
	"message-queue/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productService service.ProductService
}

func NewProductService(productService service.ProductService) ProductHandler {
	return ProductHandler{
		productService: productService,
	}
}

func (p ProductHandler) AddProduct(c *gin.Context) {
	var product models.Products

	fmt.Println("product details from front end ", product)

	if err := c.ShouldBindJSON(&product); err != nil {
		errRes := helpers.ClientResponse(http.StatusBadRequest, "payload error", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	err := p.productService.AddProduct(product)
	if err != nil {
		errRes := helpers.ClientResponse(http.StatusBadRequest, "payload error", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := helpers.ClientResponse(http.StatusCreated, "Product Added Successfully", nil, nil)
	c.JSON(http.StatusOK, successRes)

	// pass it to service for db storage and passing to message queue

	// success or error handling
}
