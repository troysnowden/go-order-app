package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetPacksWithNoItemsOrderedParam(t *testing.T) {
	r := gin.Default()

	r.GET("/packs", getPacks)

	// Create a request to send to the above route
	req, _ := http.NewRequest("GET", "/packs", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var response jsonErrorResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Invalid JSON Response")
	}

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "error", response.Response)
	assert.Equal(t, "Invalid itemsOrdered value", response.ErrorMessage)
}

func TestGetPacksWithInvalidItemsOrderedParam(t *testing.T) {
	r := gin.Default()

	r.GET("/packs", getPacks)

	// Create a request to send to the above route
	req, _ := http.NewRequest("GET", "/packs?itemsOrdered=", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var response jsonErrorResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Invalid JSON Response")
	}

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "error", response.Response)
	assert.Equal(t, "Invalid itemsOrdered value", response.ErrorMessage)

	//non-int input
	req, _ = http.NewRequest("GET", "/packs?itemsOrdered=3.14", nil)

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Invalid JSON Response")
	}

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "error", response.Response)
	assert.Equal(t, "Invalid itemsOrdered value", response.ErrorMessage)
}

func TestGetPacksWithValidItemsOrdered(t *testing.T) {
	r := gin.Default()

	r.GET("/packs", getPacks)

	// Create a request to send to the above route
	req, _ := http.NewRequest("GET", "/packs?itemsOrdered=1", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var response jsonSuccessGetResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Invalid JSON Response")
	}

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "success", response.Response)

	assert.Equal(t, 250, response.RequiredPacks[0].PackSize)
	assert.Equal(t, 1, response.RequiredPacks[0].Amount)
}

func TestPutPackSizesWithInvalidBody(t *testing.T) {
	r := gin.Default()

	r.PUT("/change-pack-sizes", putPackSizes)

	// Create a request to send to the above route
	req, _ := http.NewRequest("PUT", "/change-pack-sizes", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var response jsonErrorResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Invalid JSON Response")
	}

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "error", response.Response)
	assert.Equal(t, "Invalid NewPackSizes value", response.ErrorMessage)

	// Create a request to send to the above route
	var requestBody, e = json.Marshal(jsonChangePackSizesPutRequest{NewPackSizes: []int{}})
	if e != nil {
		t.Errorf("Invalid JSON Request")
	}
	req, _ = http.NewRequest("PUT", "/change-pack-sizes", bytes.NewBuffer(requestBody))

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Invalid JSON Response")
	}

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "error", response.Response)
	assert.Equal(t, "Invalid NewPackSizes value", response.ErrorMessage)
}

func TestPutPackSizesWithValidBody(t *testing.T) {
	r := gin.Default()

	r.PUT("/change-pack-sizes", putPackSizes)

	var response jsonSuccessPutResponse

	// Create a request to send to the above route
	var requestBody, err = json.Marshal(jsonChangePackSizesPutRequest{NewPackSizes: []int{250, 500}})
	if err != nil {
		t.Errorf("Invalid JSON Request")
	}
	req, _ := http.NewRequest("PUT", "/change-pack-sizes", bytes.NewBuffer(requestBody))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Invalid JSON Response")
	}

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "success", response.Response)
	// check the reverse order is done correctly
	assert.Equal(t, []int{500, 250}, packSizes())
	// reset to default
	setPackSizesToDefault()
}

func TestPutPackSizesBackToDefault(t *testing.T) {
	r := gin.Default()

	r.PUT("/reset-pack-sizes", resetPackSizesToDefault)

	var response jsonSuccessPutResponse

	req, _ := http.NewRequest("PUT", "/reset-pack-sizes", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Invalid JSON Response")
	}

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "success", response.Response)
	// check that the packSizes are set back to default
	assert.Equal(t, []int{5000, 2000, 1000, 500, 250}, packSizes())
}
