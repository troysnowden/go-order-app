package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test the function that returns a pack response based on number of items ordered
func TestGetPackResponse(t *testing.T) {
	var packResponse = getPackResponse(1)
	assertPackInfo(t, 250, packResponse[0].PackSize, 1, packResponse[0].Amount)

	packResponse = getPackResponse(250)
	assertPackInfo(t, 250, packResponse[0].PackSize, 1, packResponse[0].Amount)

	packResponse = getPackResponse(251)
	assertPackInfo(t, 500, packResponse[0].PackSize, 1, packResponse[0].Amount)

	packResponse = getPackResponse(501)
	assertPackInfo(t, 500, packResponse[0].PackSize, 1, packResponse[0].Amount)
	assertPackInfo(t, 250, packResponse[1].PackSize, 1, packResponse[1].Amount)

	packResponse = getPackResponse(12001)
	assertPackInfo(t, 5000, packResponse[0].PackSize, 2, packResponse[0].Amount)
	assertPackInfo(t, 2000, packResponse[1].PackSize, 1, packResponse[1].Amount)
	assertPackInfo(t, 250, packResponse[2].PackSize, 1, packResponse[2].Amount)

}

func assertPackInfo(t *testing.T, expectedPackSize int, actualPackSize int, expectedAmount int, actualAmount int) {
	assert.Equal(t, expectedPackSize, actualPackSize)
	assert.Equal(t, expectedAmount, actualAmount)
}
