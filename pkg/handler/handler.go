package handler

import (
	"log"

	"df-ecomm/pkg/service"
	"df-ecomm/pkg/util"
)

type Handler struct {
	Config  util.Config
	Logger  *log.Logger
	Product service.ProductService
	Cart    service.CartService
	User    service.UserService
}

func Setup(logger *log.Logger, config util.Config) *Handler {
	repository := &service.Repository{}
	return &Handler{
		Logger:  logger,
		Config:  config,
		Product: repository,
		Cart:    repository,
		User:    repository,
	}
}
