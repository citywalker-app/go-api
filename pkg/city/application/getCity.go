package cityapplication

import citydomain "github.com/citywalker-app/go-api/pkg/city/domain"

func GetCity(city *string) (*citydomain.City, error) {
	cityFounded, err := repo.GetCity(*city)
	if err != nil {
		return nil, err
	}

	return cityFounded, nil
}
