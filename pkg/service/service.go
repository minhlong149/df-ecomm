package service

import (
	"gorm.io/gorm"

	"df-ecomm/pkg/model"
)

type Repository struct {
	Db *gorm.DB
}

type Product = model.Product

type CartItem = model.CartItem

type ProductService interface {
	GetAll() ([]Product, error)
	Create(Product) (Product, error)
	UpdateById(int, Product) (Product, error)
	DeleteById(int) error
}

type CartService interface {
	Add(uint, uint) (CartItem, error)
	Remove(uint) error
	Checkout() ([]CartItem, error)
}

type UserService interface {
	Login(model.Account, string) (string, error)
}
