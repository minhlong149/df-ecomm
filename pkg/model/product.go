package model

type Product struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type ProductService interface {
	GetAll() []Product
}
