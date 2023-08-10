package service

import (
	"df-ecomm/pkg/model"
)

func (r *Repository) Add(productId, quantity uint) (cartItem CartItem, err error) {
	var product Product
	if err = r.Db.First(&product, productId).Error; err != nil {
		return CartItem{}, model.ErrProductNotFound
	}

	cartItem = CartItem{
		Product:   product,
		ProductId: productId,
		Quantity:  quantity,
	}

	err = r.Db.Save(&cartItem).Error
	if err != nil {
		return CartItem{}, err
	}

	return cartItem, nil
}

func (r *Repository) Remove(productId uint) error {
	return r.Db.Where("product_id = ?", productId).Delete(&CartItem{}).Error
}

func (r *Repository) Checkout() (cartItems []CartItem, err error) {
	if err = r.Db.Joins("Product").Find(&cartItems).Error; err != nil {
		return []CartItem{}, err
	}

	if err = r.Db.Exec("DELETE FROM cart_items").Error; err != nil {
		return []CartItem{}, err
	}

	return cartItems, nil
}
