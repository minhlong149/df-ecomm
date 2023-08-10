# :shopping_cart: DF E-commerce

A simple e-commerce API built with [Gin](https://gin-gonic.com/) and [GORM](https://gorm.io/) to manage products and a shopping cart.

## Getting Started

After cloning [this repository](https://github.com/minhlong149/df-ecomm) and download all the required modules (using the command `go mod download`), you can run `make db` to start the PostgreSQL database in a [Docker](https://www.docker.com/) container.

```bash
docker-compose up -d db
```

> The database is initialized with some data. You can check out the [`.sql`](./assets/pg-init.sql) file to see the database schema and the initial data.

Then, you can run `make run` to start the [Go](https://golang.org/) server.

```bash
go run cmd/server/main.go
```

The server will be listening on port `8080` by default. You can change this by setting the `PORT` environment variable in the [`.env`](./.env) file. You can also change the secret key used to sign the JWT tokens with the `SECRET_KEY` environment variable.

## API Endpoints

All `/products` APIs requires a valid JWT token in the `Authorization` header, which can be obtained by calling the `/auth` API with a valid username and password. **Currently, only users with the `admin` role can manage the products.** Check out the [`.sql`](./assets/pg-init.sql) file to see the initial users and their roles.

You can check out the [`.http`](./assets/client.http) file to see some examples of requests to the API.

> If you are using [Visual Studio Code](https://code.visualstudio.com/), you can install the [REST Client](https://marketplace.visualstudio.com/items?itemName=humao.rest-client) extension to send requests directly from the editor.
> For [GoLand](https://www.jetbrains.com/go/) users, the [HTTP Client](https://plugins.jetbrains.com/plugin/13121-http-client) plugin also provides similar functionality. See their [documentation](https://www.jetbrains.com/help/idea/http-client-in-product-code-editor.html) for more details.

| Method | Path             | Description                     | Notes                                  |
| ------ | ---------------- | ------------------------------- | -------------------------------------- |
| POST   | `/products`      | Create a new product            | Receives product details               |
| PUT    | `/products/:id`  | Update a product's details      | Receives updated product details       |
| DELETE | `/products/:id`  | Delete a product by its ID      |                                        |
| GET    | `/products`      | Retrieve a list of all products |                                        |
| POST   | `/cart/add`      | Add items to the cart           | Receives product ID and quantity       |
| DELETE | `/cart/remove`   | Remove items from the cart      | Receives product ID                    |
| POST   | `/cart/checkout` | Checkout and clear the cart     | Returns a receipt with the total price |
