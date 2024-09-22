// nolint:lll // ignore long lines for this file
package traveldomain

import (
	"time"

	citydomain "github.com/citywalker-app/go-api/pkg/city/domain"
	"github.com/citywalker-app/go-api/utils"
)

func CreateJourney(travel *Travel, path *[]citydomain.Place, matrixCost *MatrixCost) {
	// Initialize journeys
	travel.Itinerary = make([][]Itinerary, travel.Schedule.TotalDays+1)

	k := 0

	var i uint8
	for i = 0; i <= travel.Schedule.TotalDays; i++ {
		timeRemaining := travel.GetDayMinutes(i)
		travel.Itinerary[i] = make([]Itinerary, 0, 10)
		timeOfDay := getTimeOfDay(travel, i)
		first := true

		for timeRemaining > 0 && k < len(*path) {
			currentPlace := Itinerary{Place: (*path)[k]}
			if first {
				handleFirstPlace(travel, &currentPlace, &timeOfDay, &timeRemaining)
				travel.Itinerary[i] = append(travel.Itinerary[i], currentPlace)
				first = false
				k++
			} else {
				handleSubsequentPlaces(travel, &currentPlace, *path, matrixCost, &timeOfDay, &timeRemaining, &k)
				if timeRemaining <= 0 {
					break
				}
				travel.Itinerary[i] = append(travel.Itinerary[i], currentPlace)
			}
		}
	}
}

func getTimeOfDay(travel *Travel, day uint8) time.Time {
	if day == 0 {
		return travel.Schedule.StartDateTime
	}
	return travel.Schedule.StartTime.AddDate(0, 0, int(day))
}

func handleFirstPlace(travel *Travel, currentPlace *Itinerary, timeOfDay *time.Time, timeRemaining *int16) {
	currentPlace.Date = *timeOfDay
	if utils.Includes(travel.MustVisitPlaces, currentPlace.Place.Name) {
		*timeRemaining -= currentPlace.Place.Visit.All
		*timeOfDay = timeOfDay.Add(time.Duration(currentPlace.Place.Visit.All) * time.Minute)
	} else {
		*timeRemaining -= currentPlace.Place.Visit.Outside
		*timeOfDay = timeOfDay.Add(time.Duration(currentPlace.Place.Visit.Outside) * time.Minute)
	}
}

func handleSubsequentPlaces(travel *Travel, currentPlace *Itinerary, path []citydomain.Place, matrixCost *MatrixCost, timeOfDay *time.Time, timeRemaining *int16, k *int) {
	previousPlace := path[*k-1]
	currIndex := matrixCost.GetIndex(currentPlace.Place.Name)
	prevIndex := matrixCost.GetIndex(previousPlace.Name)

	isPlaceToEnter := utils.Includes(travel.MustVisitPlaces, currentPlace.Place.Name)
	visitTime := getVisitTime(currentPlace, isPlaceToEnter)
	travelTime := int16(matrixCost.Durations[currIndex][prevIndex] / 60)
	totalTime := visitTime + travelTime

	if *timeRemaining > totalTime {
		*timeRemaining -= totalTime
		updateTime(timeOfDay, travelTime)
		currentPlace.Date = *timeOfDay
		updateTime(timeOfDay, visitTime)
	} else {
		*timeRemaining = 0
	}

	*k++
}

func getVisitTime(place *Itinerary, isPlaceToEnter bool) int16 {
	if isPlaceToEnter {
		return place.Place.Visit.All
	}
	return place.Place.Visit.Outside
}

func updateTime(t *time.Time, duration int16) {
	t.Add(time.Duration(duration) * time.Minute)
}
