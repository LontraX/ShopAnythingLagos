package main

import (
	"fmt"
	"net/http"
	"sal/api"
)

func main() {

	// Define API endpoints
	http.HandleFunc("/products", api.GetProducts)
	http.HandleFunc("/products/create", api.CreateProduct)
	http.HandleFunc("/products/edit", api.EditProduct)
	http.HandleFunc("/products/delete", api.DeleteProduct)

	// Start the server
	port := 8080
	fmt.Printf("Server is running on port %d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}

}
