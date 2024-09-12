package middleware

import "errors"

var (
	ErrInvalidTokenFormat = errors.New("invalid token format, must be Bearer {token}")
	ErrInvalidToken       = errors.New("invalid token")
	ErrExpiredToken       = errors.New("expired token")
)
