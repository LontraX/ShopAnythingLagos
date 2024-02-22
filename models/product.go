package models

import (
	"fmt"
	"time"
)

// MerchantProductStore holds the products for each merchant
var MerchantProductStore = initDataBase()

func initDataBase() *map[string][]Product {

	products := make(map[string][]Product)
	return &products

}

// Product struct represents the attributes of a product
type Product struct {
	SKUID       string    `json:"sku_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	DateAdded   time.Time `json:"date_added"`
	DateUpdated time.Time `json:"date_updated"`
}

// ProductAddDTO represents the data transfer object for adding a product
type ProductAddDTO struct {
	SKUID       string  `json:"sku_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

// ProductUpdateDTO represents the data transfer object for updating a product
type ProductUpdateDTO struct {
	SKUID       string  `json:"sku_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

// GetProductsByMerchant retrieves all products listed by a merchant
func (p *Product) GetProductsByMerchant(merchantID string) (products []Product, err error) {
	products, found := (*MerchantProductStore)[merchantID]
	if !found {
		err = fmt.Errorf("no merchant with id: %s found", merchantID)
		return
	}
	return
}

// AddProductForMerchant creates a new product for a merchant
func (p *Product) AddProductForMerchant(merchantID string, productDTO ProductAddDTO) (Product, error) {
	// Validate the input data in the DTO
	if productDTO.Name == "" || productDTO.Price <= 0 {
		return Product{}, fmt.Errorf("invalid product data")
	}

	newProduct := Product{
		SKUID:       productDTO.SKUID,
		Name:        productDTO.Name,
		Description: productDTO.Description,
		Price:       productDTO.Price,
		DateAdded:   time.Now(),
	}

	// Add the new product to the store
	(*MerchantProductStore)[merchantID] = append((*MerchantProductStore)[merchantID], newProduct)

	return newProduct, nil
}

// EditProductForMerchant edits an existing product for a merchant
func (p *Product) EditProductForMerchant(merchantID string, updateDTO ProductUpdateDTO) (*Product, error) {
	// Find the index of the product to be updated
	index := p.findProductIndex(merchantID, updateDTO.SKUID)
	if index == -1 {
		return nil, fmt.Errorf("Product not found for merchant %s with SKU ID %s", merchantID, updateDTO.SKUID)
	}

	// Create a copy of the existing product to preserve the original date
	existingProduct := (*MerchantProductStore)[merchantID][index]
	updatedProduct := Product{
		SKUID:       existingProduct.SKUID,
		Name:        updateDTO.Name,
		Description: updateDTO.Description,
		Price:       updateDTO.Price,
		DateAdded:   existingProduct.DateAdded, // Preserve the original date
	}

	// Update the product in the store
	(*MerchantProductStore)[merchantID][index] = updatedProduct

	return &updatedProduct, nil
}

// DeleteProductForMerchant deletes an existing product for a merchant
func (p *Product) DeleteProductForMerchant(merchantID string, skuID string) (err error) {
	// Find the index of the product to be deleted
	index := p.findProductIndex(merchantID, skuID)
	if index == -1 {
		err = fmt.Errorf("Product not found for merchant %s with SKU ID %s", merchantID, skuID)
		return
	}

	// Remove the product from the store
	(*MerchantProductStore)[merchantID] = append(
		(*MerchantProductStore)[merchantID][:index],
		(*MerchantProductStore)[merchantID][index+1:]...,
	)

	return
}

// findProductIndex finds the index of a product in the store based on SKU ID
func (p *Product) findProductIndex(merchantID, skuID string) int {
	for i, product := range (*MerchantProductStore)[merchantID] {
		if product.SKUID == skuID {
			return i
		}
	}
	return -1
}
