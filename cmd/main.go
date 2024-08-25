package main

import (
	"go-products-api/controller"
	"go-products-api/db"
	"go-products-api/repository"
	"go-products-api/usecase"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	// load .env file
	err := godotenv.Load()
	if err != nil {
		log.Error().Err(err).Msg("unable to load .env file")
	}
	environment := os.Getenv("ENVIRONMENT")
	version := os.Getenv("VERSION")
	log.Info().Str("enveiroment", environment).Str("version", version).Msg("Initializing the Products API Service")

	logger := zerolog.New(os.Stdout).Level(zerolog.DebugLevel).With().Timestamp().Caller().Logger()
	logger.Debug().Msg("Debug message")
	logger.Info().Msg("info message")
	logger.Debug().Str("username", "joshua").Send()

	server := gin.Default()

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	dbConnection, err := db.ConnectDB()
	if err != nil {
		logger.Error().Err(err)
	}
	//repository
	ProductRepository := repository.NewProductRepository(dbConnection)
	//usecase
	ProductUsecase := usecase.NewProductUsecase(ProductRepository)
	//controller
	ProductController := controller.NewProductController(ProductUsecase)
	server.GET("/products", ProductController.GetProducts)
	server.GET("/products/:id", ProductController.GetProductsById)
	server.POST("/products", ProductController.CreateProduct)
	server.DELETE("/products/:id", ProductController.DeleteProductById)

	server.Run(":8000")

}
