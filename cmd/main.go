package main

import (
	"go-product-api/controller"
	"go-product-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	ProductUseCase := usecase.NewProductUseCase()
	// Camada de controllers
	ProductController := controller.NewProductController(ProductUseCase)

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", ProductController.GetProducts)

	server.Run(":8080")
}
