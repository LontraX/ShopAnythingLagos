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
  Example:
  ```
  curl -X GET -H "Merchant-ID: your-merchant-id" http://localhost:8080/products
  ```

- POST /products/create

  Create a new product for a merchant.
  Example:
  ```
  curl -X POST -H "Merchant-ID: your-merchant-id" -H "Content-Type: application/json" -d '{"name": "Product Name", "description": "Product Description", "price": 19.99}' http://localhost:8080/products/create
  ```

- PUT /products/edit

  Edit an existing product for a merchant.
  Example:
  ```
  curl -X PUT -H "Merchant-ID: your-merchant-id" -H "Content-Type: application/json" -d '{"sku_id": "product-sku", "name": "Updated Name", "description": "Updated Description", "price": 29.99}'         http://localhost:8080/products/edit
  ```
- DELETE /products/delete
  Delete an existing product for a merchant.
  Example:
  ```
  curl -X DELETE -H "Merchant-ID: your-merchant-id" http://localhost:8080/products/delete?sku_id=product-sku
  ```

## Testing
Run tests using the following command (API needs to be running):
```
go test ./...
```





