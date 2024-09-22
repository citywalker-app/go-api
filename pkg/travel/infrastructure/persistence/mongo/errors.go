package mongo

import "errors"

var (
	ErrTravelNotCreated = errors.New("travel not created")
	ErrConvertID        = errors.New("error converting ID")
)
