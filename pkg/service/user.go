package service

import (
	"df-ecomm/pkg/model"
	"df-ecomm/pkg/util"
)

func (r *Repository) Login(account model.Account, secretKey string) (model.User, error) {
	var claims = map[string]interface{}{
		"username": account.Username,
	}

	token, err := util.CreateToken(claims, secretKey)
	if err != nil {
		return model.User{}, err
	}

	return model.User{
		Username: account.Username,
		Token:    token,
	}, nil
}
