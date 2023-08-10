package model

type Account struct {
	Id       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role"`
}
