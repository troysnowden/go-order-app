package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const errorResponse = "error"
const successResponse = "success"

func getPacks(c *gin.Context) {
	const itemsOrderedKey = "itemsOrdered"
	if itemsOrderedInt, err := strconv.Atoi(c.Query(itemsOrderedKey)); err != nil {
		c.IndentedJSON(http.StatusBadRequest,
			jsonErrorResponse{
				ErrorMessage: errorMessage(itemsOrderedKey),
				Response:     errorResponse,
			})
	} else {
		c.IndentedJSON(http.StatusOK,
			jsonSuccessGetResponse{RequiredPacks: getPackResponse(itemsOrderedInt), Response: successResponse})
	}
}

func putPackSizes(c *gin.Context) {
	const newPackSizesKey = "NewPackSizes"
	var requestBody jsonChangePackSizesPutRequest

	if err := c.BindJSON(&requestBody); err != nil {
		c.IndentedJSON(http.StatusBadRequest,
			jsonErrorResponse{
				ErrorMessage: errorMessage(newPackSizesKey),
				Response:     errorResponse,
			})
	} else {
		if changePackSizes(requestBody.NewPackSizes) {
			c.IndentedJSON(http.StatusOK, jsonSuccessPutResponse{Response: successResponse})
		} else {
			c.IndentedJSON(http.StatusBadRequest,
				jsonErrorResponse{
					ErrorMessage: errorMessage(newPackSizesKey),
					Response:     errorResponse,
				})
		}
	}
}

func resetPackSizesToDefault(c *gin.Context) {
	changePackSizesToDefault()
	c.IndentedJSON(http.StatusOK, jsonSuccessPutResponse{Response: successResponse})
}

func errorMessage(value string) string {
	return fmt.Sprintf("Invalid %s value", value)
}
