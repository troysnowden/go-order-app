package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	// itemsOrdered param validation middleware
	getPackPaths := router.Group("/")
	getPackPaths.Use(validateItemsOrderedParam())

	// PUT JSON request body validation middleware
	putPackSizePaths := router.Group("/")
	putPackSizePaths.Use(validatePutJsonRequest())

	// frontend routes
	router.GET("/", renderIndexTemplate)

	getPackPaths.GET("pack-sizes", renderGetPacksResponse)

	// rest API routes
	getPackPaths.GET("api/packs", getPacks)

	putPackSizePaths.PUT("api/pack-sizes", putPackSizes)

	router.PUT("/api/reset-pack-sizes", resetPackSizesToDefault)

	router.Run(":" + port)
}
