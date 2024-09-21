package mongo

import "errors"

var (
	ErrCitiesNotFound = errors.New("cities not found")
	ErrCityNotFound   = errors.New("city not found")
)
