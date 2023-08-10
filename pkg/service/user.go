package service

import (
	"gorm.io/gorm"

	"df-ecomm/pkg/model"
	"df-ecomm/pkg/util"
)

func (r *Repository) Login(account model.Account, secretKey string) (token string, err error) {
	if err = r.Db.Where("username = ? AND password = ?", account.Username, account.Password).First(&account).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", model.ErrWrongCredentials
		}
		return "", err
	}

	var claims = map[string]interface{}{
		"username": account.Username,
		"role":     account.Role,
	}

	token, err = util.CreateToken(claims, secretKey)
	if err != nil {
		return "", err
	}

	return token, nil
}
