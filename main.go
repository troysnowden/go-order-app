package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.Default()

	router.GET("/packs", getPacks)

	router.PUT("/change-pack-sizes", putPackSizes)

	router.PUT("/reset-pack-sizes", putPackSizes)

	router.Run(":" + port)
}
