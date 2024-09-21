package mongo

import (
	"context"

	citydomain "github.com/citywalker-app/go-api/pkg/city/domain"
)

func (mo *Repository) GetCity(city string) (*citydomain.City, error) {
	filter := map[string]interface{}{"city": city}

	result := mo.Collection.FindOne(context.Background(), filter)
	if result.Err() != nil {
		return nil, ErrCityNotFound
	}

	var cityFounded citydomain.City

	if err := result.Decode(&cityFounded); err != nil {
		return nil, err
	}

	return &cityFounded, nil
}
