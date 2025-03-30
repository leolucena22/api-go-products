package main

import (
	usecase "api/Usecase"
	"api/controller"
	"api/db"
	"api/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	//Camada de Repository
	ProductRepository := repository.NewProductRepository(dbConnection)
	//Camada usecase
	ProductsUsecase := usecase.NewProductUsecase(ProductRepository)
	//Camada de Controllers
	ProductController := controller.NewProductController(ProductsUsecase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", ProductController.GetProducts)

	server.POST("product", ProductController.CreateProduct)

	server.GET("/product/:productId", ProductController.GetProductById)

	server.Run(":8000")
}
