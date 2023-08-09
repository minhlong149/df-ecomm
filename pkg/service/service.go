package service

import (
	"df-ecomm/pkg/model"
)

type Repository struct{}

type CartService interface {
	Add(model.ItemToAdd) (model.CartItem, error)
	Remove(model.ItemToRemove) error
	Checkout() model.Receipt
}

type ProductService interface {
	GetAll() []model.Product
	Create(model.NewProduct) model.Product
	UpdateById(int, model.Product) (model.Product, error)
	DeleteById(int) error
}

type UserService interface {
	Login(model.Account, string) (model.User, error)
}
