package service

import (
	"df-ecomm/pkg/model"
	"df-ecomm/pkg/util"
)

type Account = model.Account

type UserService struct{
	SecretKey string
}

var accounts []Account = []Account{
	{Username: "root", Password: "password"},
}

func (s *UserService) ExistsByUsername(username string) bool {
	for i := range accounts {
		if accounts[i].Username == username {
			return true
		}
	}

	return false
}

func (s *UserService) Login(account Account) (model.User, error) {
	for i := range accounts {
		if accounts[i].Username == account.Username && accounts[i].Password == account.Password {
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
	}

	return model.User{}, model.ErrUnauthorized
}
