package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	"go-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()
	
	dbConnection, err := db.ConnectDB()
	if(err != nil) {
		panic(err)
	}
	
	//Camada de repository
	ProductRepository := repository.NewProductRepository(dbConnection)
	// Camada usecase
	ProductUsecase := usecase.NewProductUseCase(ProductRepository)
	//Camada de controllers	
	productController := controller.NewProductController(ProductUsecase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
	
	server.GET("/products", productController.GetProducts)

	server.Run(":8000")

}
