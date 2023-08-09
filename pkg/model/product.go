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


