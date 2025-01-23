package main

import (
	"go-product-api/controller"
	"go-product-api/db"
	"go-product-api/repository"
	"go-product-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	// Camada Repository
	ProductRepository := repository.NewProductRepository(dbConnection)

	// Camada UseCase
	ProductUseCase := usecase.NewProductUseCase(ProductRepository)

	// Camada de Controllers
	ProductController := controller.NewProductController(ProductUseCase)

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", ProductController.GetProducts)
	server.POST("/product", ProductController.CreateProduct)

	server.Run(":8080")
}
