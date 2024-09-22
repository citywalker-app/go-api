package traveldomain

import (
	citydomain "github.com/citywalker-app/go-api/pkg/city/domain"
	"github.com/citywalker-app/go-api/utils"
)

func (t *Travel) CreateItinerary(city *citydomain.City) error {
	totalMinutes := t.GetTotalMinutes()

	placesToVisit := []citydomain.Place{}

	for _, place := range city.Places {
		if totalMinutes <= 0 {
			break
		}

		if shouldSkipPlace(t, place, totalMinutes) {
			continue
		}

		visitDuration := getVisitDuration(t, place)
		placesToVisit = append(placesToVisit, place)
		totalMinutes -= visitDuration
	}

	err := t.GetTSPTWSolve(&placesToVisit)

	return err
}

func shouldSkipPlace(t *Travel, place citydomain.Place, totalMinutes int16) bool {
	if utils.Includes(t.ExcludedCategories, place.Category) {
		return true
	}

	if utils.Includes(t.MustVisitPlaces, place.Name) {
		return place.Visit.All+10 > totalMinutes
	}

	if place.Visit.Outside == 0 {
		return true
	}

	return place.Visit.Outside+10 > totalMinutes
}

func getVisitDuration(t *Travel, place citydomain.Place) int16 {
	if utils.Includes(t.MustVisitPlaces, place.Name) {
		return place.Visit.All + 10
	}
	return place.Visit.Outside + 10
}
