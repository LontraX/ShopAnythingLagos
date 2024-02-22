# ShopAnythingLagos (SAL) Product Management

ShopAnythingLagos (SAL) is a Saas company that provides a platform for merchants to manage their product listings. This project includes a set of REST APIs for CRUD operations on merchant products.

## Table of Contents

- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Usage](#usage)
  - [API Endpoints](#api-endpoints)
- [Testing](#testing)

## Getting Started

### Prerequisites

- Go installed on your machine. You can download it from [here](https://golang.org/dl/).

### Installation

1. Clone the repository:

```bash
git clone https://github.com/LontraX/ShopAnythingLagos.git
cd your-repo
go run main.go
```
- The application should now be running on http://localhost:8080.


## Usage
### API Endpoints

- GET /products

  Retrieve all products listed by a merchant.
  
  Request Headers:
  Merchant-ID: Include the merchant ID in the request header.
  
  Example:
  ```
  curl -X GET -H "Merchant-ID: your-merchant-id" http://localhost:8080/products
  ```
   ``` json
  {
  "status": 200,
  "message": "Success",
  "data": [
    {
      "sku_id": "product-sku-1",
      "name": "Product 1",
      "description": "Description of Product 1",
      "price": 19.99,
      "date_added": "2024-02-18T12:34:56Z"
    },
    {
      "sku_id": "product-sku-2",
      "name": "Product 2",
      "description": "Description of Product 2",
      "price": 29.99,
      "date_added": "2024-02-18T12:45:00Z"
    }
  ]
  }
  ```

- POST /products/create

  Create a new product for a merchant.
  
  Request Headers:

  Merchant-ID: Include the merchant ID in the request header.
  Content-Type: Set to application/json in the request header.
  Example:
  ```
  curl -X POST -H "Merchant-ID: your-merchant-id" -H "Content-Type: application/json" -d '{"sku_id": "product-sku","name": "Product Name", "description": "Product Description", "price": 19.99}' http://localhost:8080/products/create
  ```

  ``` json
  {
  "status": 201,
  "message": "Product created successfully",
  "data": {
    "sku_id": "new-product-sku",
    "name": "New Product",
    "description": "Description of New Product",
    "price": 39.99,
    "date_added": "2024-02-18T13:00:00Z"
  }
  }

  ```

 


- PUT /products/edit

  Edit an existing product for a merchant.
  
  Request Headers:

  Merchant-ID: Include the merchant ID in the request header.
  Content-Type: Set to application/json in the request header.
  Example:
  ```
  curl -X PUT -H "Merchant-ID: your-merchant-id" -H "Content-Type: application/json" -d '{"sku_id": "product-sku", "name": "Updated Name", "description": "Updated Description", "price": 29.99}'         http://localhost:8080/products/edit
  ```

  ```json
  {
  "status": 200,
  "message": "Product edited successfully",
  "data": {
    "sku_id": "product-sku-1",
    "name": "Updated Product 1",
    "description": "Updated Description",
    "price": 49.99,
    "date_added": "2024-02-18T13:15:00Z"
  }
  }

  ```
- DELETE /products/delete
  Delete an existing product for a merchant.

  Request Headers:

  Merchant-ID: Include the merchant ID in the request header.
  URL Parameters:

  sku_id: Include SKU ID in the request URL parameters.
  Example:
  ```
  curl -X DELETE -H "Merchant-ID: your-merchant-id" http://localhost:8080/products/delete?sku_id=product-sku
  ```

  ```json
  {
  "status": 200,
  "message": "Product deleted successfully",
  "data": null
  }
  ```

## Testing
Run tests using the following command (API needs to be running):
```
go test ./...
```





