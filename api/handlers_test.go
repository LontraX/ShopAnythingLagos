package api

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetProducts(t *testing.T) {
	// Create a test server with the handler

	// Create a request with the correct method and headers
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080/products", nil)
	assert.NoError(t, err)
	req.Header.Set("Merchant-ID", "1234")

	// Perform the request
	resp, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)
	defer resp.Body.Close()

	// Check the response status code
	assert.Equal(t, http.StatusOK, resp.StatusCode)

}
