package model

import (
	"errors"
)

var (
	ErrProductNotFound  = errors.New("the product you're looking for has been abducted by aliens")
	ErrInvalidProductId = errors.New("your product ID is so invalid, it's not even a real number")
	ErrInvalidProduct   = errors.New("i'm sorry, but your product is not a valid member of the Product Club")
)

var (
	ErrInvalidCart = errors.New("your cart is invalid, maybe you should try adding some real items instead of just dreams")
)

var (
	ErrInvalidCredentials = errors.New("i can't let you in without knowing who you are, please enter your username and password")
	ErrWrongCredentials = errors.New("that's not the password I was expecting, are you sure you're not trying to log in to someone else's account?")
)

var (
	ErrInvalidAuthHeader = errors.New("your authorization header is so invalid, it makes me question your entire life")
	ErrInvalidToken      = errors.New("your token is so invalid, it got rejected from Hogwarts")
	ErrUnauthorized      = errors.New("looks like you're not authorized to access this site, maybe you should try logging in as someone else?")
)
