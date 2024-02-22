package models

import (
	"testing"
)

func TestAddProductForMerchant(t *testing.T) {
	// Initialize the MerchantProductStore
	MerchantProductStore = initDataBase()

	// Mock data
	merchantID := "merchant1"
	productDTO := ProductAddDTO{
		Name:        "Test Product",
		Description: "Test Description",
		Price:       9.99,
	}

	// Create a new Product instance
	product := Product{}

	// Call the method
	_, err := product.AddProductForMerchant(merchantID, productDTO)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Retrieve products for the merchant
	products, err := product.GetProductsByMerchant(merchantID)
	if err != nil {
		t.Errorf("Error retrieving products: %v", err)
	}

	// Check if the product was added
	if len(products) != 1 {
		t.Errorf("Expected 1 product, got %d", len(products))
	}

	// Check if product attributes match
	addedProduct := products[0]
	if addedProduct.Name != productDTO.Name ||
		addedProduct.Description != productDTO.Description ||
		addedProduct.Price != productDTO.Price {
		t.Errorf("Product attributes do not match")
	}
}

func TestEditProductForMerchant(t *testing.T) {
	// Initialize the MerchantProductStore
	MerchantProductStore = initDataBase()

	// Mock data
	merchantID := "merchant1"
	productDTO := ProductAddDTO{
		Name:        "Test Product",
		Description: "Test Description",
		Price:       9.99,
	}

	// Create a new Product instance
	product := Product{}

	// Add a product
	_, err := product.AddProductForMerchant(merchantID, productDTO)
	if err != nil {
		t.Errorf("Error adding product: %v", err)
	}

	// Retrieve products for the merchant
	products, err := product.GetProductsByMerchant(merchantID)
	if err != nil {
		t.Errorf("Error retrieving products: %v", err)
	}

	// Check if the product was added
	if len(products) != 1 {
		t.Errorf("Expected 1 product, got %d", len(products))
	}

	// Get SKU ID of the added product
	skuID := products[0].SKUID

	// Mock updated data
	updateDTO := ProductUpdateDTO{
		SKUID:       skuID,
		Name:        "Updated Product",
		Description: "Updated Description",
		Price:       19.99,
	}

	// Call the method to edit the product
	_, err = product.EditProductForMerchant(merchantID, updateDTO)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Retrieve products for the merchant after update
	updatedProducts, err := product.GetProductsByMerchant(merchantID)
	if err != nil {
		t.Errorf("Error retrieving products: %v", err)
	}

	// Check if product attributes are updated
	updatedProduct := updatedProducts[0]
	if updatedProduct.Name != updateDTO.Name ||
		updatedProduct.Description != updateDTO.Description ||
		updatedProduct.Price != updateDTO.Price {
		t.Errorf("Product attributes are not updated")
	}
}

func TestGetProductsByMerchant(t *testing.T) {
	// Initialize the MerchantProductStore
	MerchantProductStore = initDataBase()

	// Mock data
	merchantID := "merchant1"
	productDTO := ProductAddDTO{
		Name:        "Test Product",
		Description: "Test Description",
		Price:       9.99,
	}

	// Create a new Product instance
	product := Product{}

	// Add a product
	_, err := product.AddProductForMerchant(merchantID, productDTO)
	if err != nil {
		t.Errorf("Error adding product: %v", err)
	}

	// Call the method to retrieve products
	products, err := product.GetProductsByMerchant(merchantID)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Check if the product is retrieved
	if len(products) != 1 {
		t.Errorf("Expected 1 product, got %d", len(products))
	}
}

func TestDeleteProductForMerchant(t *testing.T) {
	// Initialize the MerchantProductStore
	MerchantProductStore = initDataBase()

	// Mock data
	merchantID := "merchant1"
	productDTO := ProductAddDTO{
		Name:        "Test Product",
		Description: "Test Description",
		Price:       9.99,
	}

	// Create a new Product instance
	product := Product{}

	// Add a product
	addedProduct, err := product.AddProductForMerchant(merchantID, productDTO)
	if err != nil {
		t.Errorf("Error adding product: %v", err)
	}

	// Call the method to delete the product
	err = product.DeleteProductForMerchant(merchantID, addedProduct.SKUID)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Retrieve products for the merchant
	products, err := product.GetProductsByMerchant(merchantID)
	if err != nil {
		t.Errorf("Error retrieving products: %v", err)
	}

	// Check if the product is deleted
	if len(products) != 0 {
		t.Errorf("Expected 0 products, got %d", len(products))
	}
}
