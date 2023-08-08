# DF E-commerce

A simple E-commerce server with product and shopping cart management APIs. Built with Go Gin framework.

| Method | Path            | Description                     | Notes                                  |
| ------ | --------------- | ------------------------------- | -------------------------------------- |
| POST   | /products       | Create a new product            | Receives product details               |
| PUT    | /products/`:id` | Update a product's details      | Receives updated product details       |
| DELETE | /products/`:id` | Delete a product by its ID      |                                        |
| GET    | /products       | Retrieve a list of all products |                                        |
| POST   | /cart/add       | Add items to the cart           | Receives product ID and quantity       |
| DELETE | /cart/remove    | Remove items from the cart      | Receives product ID                    |
| POST   | /cart/checkout  | Checkout and clear the cart     | Returns a receipt with the total price |

> Note: The `/products` API requires a valid JWT token in the `Authorization` header.
> The token can be obtained by calling the `/auth` API with any credentials.
