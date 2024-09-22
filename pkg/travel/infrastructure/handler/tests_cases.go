package travelhandler

import (
	"time"

	traveldomain "github.com/citywalker-app/go-api/pkg/travel/domain"
)

type TestCase struct {
	Name       string
	Travel     traveldomain.Travel
	StatusCode int
}

var location, _ = time.LoadLocation("UTC")

var travel = traveldomain.Travel{
	City: "London",
	Schedule: traveldomain.Schedule{
		StartDate:     time.Date(2024, 2, 10, 10, 0, 0, 0, location),
		EndDate:       time.Date(2024, 2, 13, 10, 0, 0, 0, location),
		StartDateTime: time.Date(2024, 2, 10, 16, 0, 0, 0, location),
		EndDateTime:   time.Date(2024, 2, 10, 15, 0, 0, 0, location),
		StartTime:     time.Date(2024, 2, 10, 9, 0, 0, 0, location),
		EndTime:       time.Date(2024, 2, 10, 20, 0, 0, 0, location),
	},
	ExcludedCategories: []string{"Museum"},
	MustVisitPlaces:    []string{"Imperial War Museum"},
}

var CreateTestCases = []TestCase{
	{
		Name: "Travel with no city(fail)",
		Travel: traveldomain.Travel{
			Schedule:           travel.Schedule,
			ExcludedCategories: travel.ExcludedCategories,
			MustVisitPlaces:    travel.MustVisitPlaces,
		},
		StatusCode: 400,
	},
	{
		Name: "Travel with no schedule(fail)",
		Travel: traveldomain.Travel{
			City:               travel.City,
			ExcludedCategories: travel.ExcludedCategories,
			MustVisitPlaces:    travel.MustVisitPlaces,
		},
		StatusCode: 400,
	},
	{
		Name: "Travel with no excluded categories(fail)",
		Travel: traveldomain.Travel{
			City:            travel.City,
			Schedule:        travel.Schedule,
			MustVisitPlaces: travel.MustVisitPlaces,
		},
		StatusCode: 400,
	},
	{
		Name: "Travel with no must visit places(fail)",
		Travel: traveldomain.Travel{
			City:               travel.City,
			Schedule:           travel.Schedule,
			ExcludedCategories: travel.ExcludedCategories,
		},
		StatusCode: 400,
	},
	{
		Name: "Travel with no existing city(fail)",
		Travel: traveldomain.Travel{
			City:               "City",
			Schedule:           travel.Schedule,
			ExcludedCategories: travel.ExcludedCategories,
			MustVisitPlaces:    travel.MustVisitPlaces,
		},
		StatusCode: 404,
	},
	{
		Name:       "Travel valid(success)",
		Travel:     travel,
		StatusCode: 200,
	},
}
