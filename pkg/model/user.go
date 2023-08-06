package model

import "errors"

type Account struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type User struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

type UserService interface {
	ExistsByUsername(string) bool
	Login(Account) (User, error)
}

var ErrUnauthorized = errors.New("unauthorized")
