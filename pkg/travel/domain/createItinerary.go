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

		if shouldSkipPlace(t, place, totalMinutes, *city) {
			continue
		}

		visitDuration := getVisitDuration(t, place, *city)
		placesToVisit = append(placesToVisit, place)
		totalMinutes -= visitDuration
	}

	err := t.GetTSPTWSolve(&placesToVisit)

	return err
}

func shouldSkipPlace(t *Travel, place citydomain.Place, totalMinutes uint16, city citydomain.City) bool {
	if utils.Includes(t.ExcludedCategories, place.Category) {
		return true
	}

	if utils.Includes(t.MustVisitPlaces, place.Name) {
		return place.Visit.All+city.AverageCost > totalMinutes
	}

	if place.Visit.Outside == 0 {
		return true
	}

	return place.Visit.Outside+city.AverageCost > totalMinutes
}

func getVisitDuration(t *Travel, place citydomain.Place, city citydomain.City) uint16 {
	if utils.Includes(t.MustVisitPlaces, place.Name) {
		return place.Visit.All + city.AverageCost
	}
	return place.Visit.Outside + city.AverageCost
}

// func (t *Travel) CreateItinerary(city *citydomain.City) error {
// 	totalMinutes := t.GetTotalMinutes()

// 	placesToVisit := []citydomain.Place{}

// 	for _, place := range city.Places {
// 		if totalMinutes > 0 {
// 			if utils.Includes(t.DiscardedCategories, place.Category) {
// 				continue
// 			}
// 			if utils.Includes(t.PlacesToEnter, place.Name) {
// 				if place.Visit.All+city.AverageCost > totalMinutes {
// 					continue
// 				}
// 				placesToVisit = append(placesToVisit, place)
// 				totalMinutes -= place.Visit.All + city.AverageCost
// 				continue
// 			}
// 			if place.Visit.Outside == 0 {
// 				continue
// 			}
// 			if place.Visit.Outside+city.AverageCost > totalMinutes {
// 				continue
// 			}
// 			placesToVisit = append(placesToVisit, place)
// 			totalMinutes -= place.Visit.Outside + city.AverageCost
// 		} else {
// 			break
// 		}
// 	}

// 	err := t.GetTSPTWSolve(&placesToVisit)

// 	return err
// }
