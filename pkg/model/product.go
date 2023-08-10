package model

type Product struct {
	Id    uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Price uint   `json:"price"`
}
