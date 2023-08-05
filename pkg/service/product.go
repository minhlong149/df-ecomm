package service

import "df-ecomm/pkg/model"

type Product = model.Product

type ProductService struct{}

var products []Product = []Product{
	{Id: 1, Name: "Product 1", Price: 100},
	{Id: 2, Name: "Product 2", Price: 200},
	{Id: 3, Name: "Product 3", Price: 300},
}

func (s *ProductService) GetAll() []Product {
	return products
}
