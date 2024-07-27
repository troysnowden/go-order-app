package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/packs", getPacks)

	// router.PUT("/change-pack-size", changePackSize)

	// router.PUT("/add-pack-size", addPackSize)

	// router.PUT("/remove-pack-size", removePackSize)

	router.Run("localhost:8080")
}

func getPacks(c *gin.Context) {
	var itemsOrderedParam = c.Query("itemsOrdered")
	var itemsOrderedInt, err = strconv.Atoi(itemsOrderedParam)
	if err != nil {
		// change
		c.IndentedJSON(http.StatusOK, jsonErrorResppnse{
			ErrorMessage: "Invalid itemsOrdered value",
			Response:     "error",
		})
	}

	c.IndentedJSON(http.StatusOK, jsonSuccessResponse{RequiredPacks: getPackResponse(itemsOrderedInt), Response: "success"})
}
