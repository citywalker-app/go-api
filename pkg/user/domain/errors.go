package userdomain

import "errors"

var (
	ErrUserNotFound     = errors.New("user not found")
	ErrUserExists       = errors.New("user already exists")
	ErrWrongCredentials = errors.New("wrong credentials")
	ErrBadRequest       = errors.New("user wrong structure")
)
