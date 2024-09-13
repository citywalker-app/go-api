package mongo

import "errors"

var (
	ErrUserNotInserted = errors.New("user can not be created")
	ErrUserNotUpdated  = errors.New("user can not be updated")
	ErrUserNotFound    = errors.New("user not found")
)
