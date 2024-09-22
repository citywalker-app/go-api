package traveldomain

import "errors"

var (
	ErrPlacesNotFound  = errors.New("places not found")
	ErrCitiesNotFound  = errors.New("cities not found")
	ErrItinerary       = errors.New("itinerary not found")
	ErrNotEnoughPlaces = errors.New("not enough places to create an itinerary")
)
