package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func validateItemsOrderedParam() gin.HandlerFunc {
	return func(c *gin.Context) {
		const itemsOrderedKey = "itemsOrdered"
		if _, err := strconv.Atoi(c.Query(itemsOrderedKey)); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, jsonErrorResponse{
				ErrorMessage: errorMessage(itemsOrderedKey),
				Response:     errorResponse,
			})
			return
		}
	}
}

func validatePutJsonRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestBody jsonChangePackSizesPutRequest
		if err := c.BindJSON(&requestBody); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				jsonErrorResponse{
					ErrorMessage: errorMessage("NewPackSizes"),
					Response:     errorResponse,
				})
			return
		}
	}
}
