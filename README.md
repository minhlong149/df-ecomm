# Go Dwarves E-commerce

A simple E-commerce server with product and shopping cart management APIs, built with Go Gin framework. *This is a solution to the web development assignment for the [Golang course at Dwarves Foundation](https://github.com/minhlong149/go-dwarves-assignment).*

## Getting Started

After cloning the repository and instal all the packages, you can run the server with the following command:

```bash
go run cmd/server/main.go
```

## Endpoints

> Note: All `/products` APIs requires a valid JWT token in the `Authorization` header.
> A token can be obtained by calling the `/auth` API with *any* credentials.

| Method | Path            | Description                     | Notes                                  |
| ------ | --------------- | ------------------------------- | -------------------------------------- |
| POST   | /products       | Create a new product            | Receives product details               |
| PUT    | /products/`:id` | Update a product's details      | Receives updated product details       |
| DELETE | /products/`:id` | Delete a product by its ID      |                                        |
| GET    | /products       | Retrieve a list of all products |                                        |
| POST   | /cart/add       | Add items to the cart           | Receives product ID and quantity       |
| DELETE | /cart/remove    | Remove items from the cart      | Receives product ID                    |
| POST   | /cart/checkout  | Checkout and clear the cart     | Returns a receipt with the total price |
