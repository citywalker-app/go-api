package userhandler

import "errors"

var (
	ErrBadRequest    = errors.New("could not be parsed body to user")
	ErrJWTGeneration = errors.New("could not be generate jwt")
	ErrEmail         = errors.New("could not be send email")
	ErrUserNotFound  = errors.New("user not found")
)
