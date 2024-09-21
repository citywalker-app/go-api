package cityapplication

import citydomain "github.com/citywalker-app/go-api/pkg/city/domain"

func GetCities(lng *string) (*[]citydomain.City, error) {
	cities, err := repo.GetAll(*lng)
	if err != nil {
		return nil, err
	}

	return cities, nil
}
