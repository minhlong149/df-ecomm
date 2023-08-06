package model

type Product struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type NewProduct struct {
	Name  string `json:"name" binding:"required"`
	Price int    `json:"price" binding:"required"`
}

type ProductService interface {
	GetAll() []Product
	Create(NewProduct) Product
	UpdateById(int, Product) (Product, error)
	DeleteById(int) error
}
