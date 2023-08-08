package service

import (
	"df-ecomm/pkg/model"
	"df-ecomm/pkg/util"
)

type Account = model.Account

type UserService struct {
	SecretKey string
}

func (s *UserService) Login(account Account) (model.User, error) {
	var claims = map[string]interface{}{
		"username": account.Username,
	}

	token, err := util.CreateToken(claims, s.SecretKey)
	if err != nil {
		return model.User{}, err
	}

	return model.User{
		Username: account.Username,
		Token:    token,
	}, nil
}
