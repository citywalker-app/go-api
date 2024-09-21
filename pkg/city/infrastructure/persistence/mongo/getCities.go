package mongo

import (
	"context"

	citydomain "github.com/citywalker-app/go-api/pkg/city/domain"
)

func (mo *Repository) GetAll(lng string) (*[]citydomain.City, error) {
	filter := map[string]interface{}{"lng": lng}

	cur, err := mo.Collection.Find(context.Background(), filter)
	if err != nil {
		return nil, ErrCitiesNotFound
	}

	var cities []citydomain.City

	for cur.Next(context.Background()) {
		var city citydomain.City
		err := cur.Decode(&city)
		if err != nil {
			return nil, err
		}

		cities = append(cities, city)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return &cities, nil
}
