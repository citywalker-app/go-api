package travelapplication

import (
	"context"

	cityapplication "github.com/citywalker-app/go-api/pkg/city/application"
	citydomain "github.com/citywalker-app/go-api/pkg/city/domain"
	traveldomain "github.com/citywalker-app/go-api/pkg/travel/domain"
	"github.com/citywalker-app/go-api/utils"
	"github.com/gofiber/fiber/v2/log"
)

func Create(travel *traveldomain.Travel) (*traveldomain.Travel, error) {
	var city *citydomain.City

	err := utils.GetCache(context.Background(), "/cities/"+travel.City, city)
	if err != nil {
		city, err = cityapplication.GetCity(&travel.City)
		if err != nil {
			return nil, err
		}

		if err := utils.SetCache(context.Background(), "/cities/"+travel.City, city); err != nil {
			log.Error("Error setting cache: %v", err)
		}
	}

	if err := travel.CreateItinerary(city); err != nil {
		return nil, err
	}

	if err := Repo.Create(travel); err != nil {
		return nil, err
	}

	return travel, nil
}
