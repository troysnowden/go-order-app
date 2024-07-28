package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/packs", getPacks)

	router.PUT("/change-pack-sizes", putPackSizes)

	router.PUT("/reset-pack-sizes", putPackSizes)

	router.Run("localhost:8080")
}
