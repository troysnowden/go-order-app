package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const errorResponse = "error"
const successResponse = "success"

// these are the endpoints used by the frontend
func renderIndexTemplate(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func renderGetPacksResponse(c *gin.Context) {
	itemsOrderedInt, _ := strconv.Atoi(c.Query("itemsOrdered"))

	var packsRequired, minimumItemsToSend = getPackResponse(itemsOrderedInt)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":        "Home Page",
		"packs":        packsRequired,
		"itemsOrdered": itemsOrderedInt,
		"itemsToSend":  minimumItemsToSend})
}

// these are the rest API endpoints
func getPacks(c *gin.Context) {
	itemsOrderedInt, _ := strconv.Atoi(c.Query("itemsOrdered"))

	var packsRequired, _ = getPackResponse(itemsOrderedInt)

	c.IndentedJSON(http.StatusOK,
		jsonSuccessGetResponse{RequiredPacks: packsRequired, Response: successResponse})
}

func putPackSizes(c *gin.Context) {
	var requestBody jsonChangePackSizesPutRequest
	c.BindJSON(&requestBody)

	if changePackSizes(requestBody.NewPackSizes) {
		c.IndentedJSON(http.StatusOK, jsonSuccessPutResponse{Response: successResponse})
	} else {
		c.IndentedJSON(http.StatusBadRequest,
			jsonErrorResponse{
				ErrorMessage: errorMessage("NewPackSizes"),
				Response:     errorResponse,
			})
	}
}

func resetPackSizesToDefault(c *gin.Context) {
	changePackSizesToDefault()
	c.IndentedJSON(http.StatusOK, jsonSuccessPutResponse{Response: successResponse})
}

func errorMessage(value string) string {
	return fmt.Sprintf("Invalid %s value", value)
}
