package model

import (
	"errors"
)

var ErrNotFound = errors.New("product not found")

var ErrAlreadyExists = errors.New("product already exists")

var ErrInvalidToken = errors.New("invalid token")
