package model

type ItemToAdd struct {
	ProductId int `json:"id" binding:"required"`
	Quantity  int `json:"quantity" binding:"required"`
}

type ItemToRemove struct {
	ProductId int `json:"id" binding:"required"`
}

type CartItem struct {
	Product  Product `json:"product"`
	Quantity int     `json:"quantity"`
}

type Receipt struct {
	Items []CartItem `json:"cart"`
	Total int         `json:"total"`
}

type CartService interface {
	AddItem(ItemToAdd) (CartItem, error)
	RemoveItem(ItemToRemove) error
	Checkout() Receipt
}
