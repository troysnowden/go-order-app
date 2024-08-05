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

// this doesn't work due to being unable to unmarshal json request more than once.
// func validatePutJsonRequest() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var requestBody jsonChangePackSizesPutRequest
// 		if err := c.ShouldBindJSON(&requestBody); err != nil {
// 			c.AbortWithStatusJSON(http.StatusBadRequest,
// 				jsonErrorResponse{
// 					ErrorMessage: errorMessage("NewPackSizes"),
// 					Response:     errorResponse,
// 				})
// 			c.Abort()
// 			return
// 		}
// 		c.Next()
// 	}
// }
