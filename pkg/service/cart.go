package service

import "df-ecomm/pkg/model"

type ItemToAdd = model.ItemToAdd

type CartService struct{}

var cart = []ItemToAdd{
	{ProductId: 1, Quantity: 1},
	{ProductId: 2, Quantity: 2},
}

func (s *CartService) AddItem(newItem ItemToAdd) (model.CartItem, error) {
	foundProduct := model.Product{}
	for _, product := range products {
		if product.Id == newItem.ProductId {
			foundProduct = product
			break
		}
	}

	if foundProduct == (model.Product{}) {
		return model.CartItem{}, model.ErrNotFound
	}

	for i := range cart {
		if cart[i].ProductId == newItem.ProductId {
			cart[i].Quantity += newItem.Quantity
			return model.CartItem{
				Product:  foundProduct,
				Quantity: cart[i].Quantity,
			}, nil
		}
	}

	cart = append(cart, newItem)
	return model.CartItem{
		Product:  foundProduct,
		Quantity: newItem.Quantity,
	}, nil
}

func (s *CartService) RemoveItem(item model.ItemToRemove) error {
	for i := range cart {
		if cart[i].ProductId == item.ProductId {
			cart = append(cart[:i], cart[i+1:]...)
			return nil
		}
	}
	return model.ErrNotFound
}

func (s *CartService) Checkout() model.Receipt {
	receipt := model.Receipt{
		Items: []model.CartItem{},
		Total: 0,
	}

	for _, product := range products {
		for _, item := range cart {
			if product.Id == item.ProductId {
				receipt.Items = append(receipt.Items, model.CartItem{
					Product:  product,
					Quantity: item.Quantity,
				})
				receipt.Total += product.Price * item.Quantity
			}
		}
	}

	cart = []ItemToAdd{}
	return receipt
}