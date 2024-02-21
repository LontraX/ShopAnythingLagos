package api

// api/handlers.go
import (
	"encoding/json"
	"net/http"
	"sal/models"
)

var productInstance models.Product

// GetProducts retrieves all products listed by a merchant
func GetProducts(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	// Extract merchant ID from the request header
	merchantID := r.Header.Get("Merchant-ID")
	if merchantID == "" {
		http.Error(w, "Include a Merchant-ID in your request", http.StatusUnauthorized)
		return
	}

	// Retrieve products for the given merchant
	products, err := productInstance.GetProductsByMerchant(merchantID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	respondWithJSON(w, http.StatusOK, "Success", products)
}

// CreateProduct creates a new product for a merchant
func CreateProduct(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	// Extract merchant ID from the request header
	merchantID := r.Header.Get("Merchant-ID")
	if merchantID == "" {
		http.Error(w, "Include a Merchant-ID in your request", http.StatusUnauthorized)
		return
	}

	// Parse the request body to get the new product details
	var newProductToAdd models.ProductAddDTO
	err := json.NewDecoder(r.Body).Decode(&newProductToAdd)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	addedProduct, err := productInstance.AddProductForMerchant(merchantID, newProductToAdd)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	respondWithJSON(w, http.StatusCreated, "Product created successfully", addedProduct)
}

// EditProduct edits an existing product for a merchant
func EditProduct(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPut {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	// Extract merchant ID from the request header
	merchantID := r.Header.Get("Merchant-ID")

	if merchantID == "" {
		http.Error(w, "Include a Merchant-ID in your request", http.StatusUnauthorized)
		return
	}

	// Parse the request body to get the updated product details
	var updatedProduct models.ProductUpdateDTO
	err := json.NewDecoder(r.Body).Decode(&updatedProduct)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Call the method to edit the product for the merchant
	editedProduct, err := productInstance.EditProductForMerchant(merchantID, updatedProduct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, "Product edited successfully", editedProduct)
}

// DeleteProduct deletes an existing product for a merchant
func DeleteProduct(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodDelete {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	// Extract merchant ID from the request header
	merchantID := r.Header.Get("Merchant-ID")
	if merchantID == "" {
		http.Error(w, "Include a Merchant-ID in your request", http.StatusUnauthorized)
		return
	}

	// Extract SKU ID from the URL parameters
	skuID := r.URL.Query().Get("sku_id")
	if skuID == "" {
		http.Error(w, "Include SKU ID in your request URL parameters", http.StatusBadRequest)
		return
	}

	// Call the method to delete the product for the merchant
	err := productInstance.DeleteProductForMerchant(merchantID, skuID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, "Product deleted successfully", nil)
}

// Common function for responding with JSON
func respondWithJSON(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	response := map[string]interface{}{
		"status":  statusCode,
		"message": message,
		"data":    data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}
