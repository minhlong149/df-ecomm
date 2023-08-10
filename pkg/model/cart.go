package model

type CartItem struct {
	ProductId uint    `json:"-" gorm:"primaryKey"`
	Product   Product `json:"product" gorm:"constraint:OnDelete:CASCADE;"`
	Quantity  uint    `json:"quantity"`
}
