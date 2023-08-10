package handler

import (
	"gorm.io/gorm"

	"df-ecomm/pkg/service"
	"df-ecomm/pkg/util"
)

type Handler struct {
	Config  util.Config
	Logger  util.Logger
	Product service.ProductService
	Cart    service.CartService
	User    service.UserService
}

func Setup(logger util.Logger, config util.Config, dataSource *gorm.DB) *Handler {
	repository := &service.Repository{Db: dataSource}
	return &Handler{
		Logger:  logger,
		Config:  config,
		Product: repository,
		Cart:    repository,
		User:    repository,
	}
}
