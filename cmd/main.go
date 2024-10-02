package main

import (
	"github.com/gin-gonic/gin"
	"go-api/controller"
	"go-api/usecase"
	"go-api/repository"
	"go-api/db"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}
	
	ProductRepository := repository.NewProductRepository(dbConnection)

	ProductUsecase := usecase.NewProductUseCase(ProductRepository)
	productController := controller.NewProductController(ProductUsecase)

	server.GET("/ping", func(ctx *gin.Context){
		ctx.JSON(200, gin.H {
			"message": "pong",		
		})
	})

	server.GET("/products", productController.GetProducts)
	server.GET("/product/:productId", productController.GetProductById)
	server.POST("/product", productController.CreateProduct)

	server.Run(":8001")
}