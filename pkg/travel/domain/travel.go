package traveldomain

import (
	"time"

	citydomain "github.com/citywalker-app/go-api/pkg/city/domain"
)

type Travel struct {
	ID                 string        `json:"_id"`
	City               string        `json:"city" validate:"required"`
	Schedule           Schedule      `json:"schedule" validate:"required"`
	ExcludedCategories []string      `json:"excludedCategories" validate:"required"`
	MustVisitPlaces    []string      `json:"MustVisitPlaces" validate:"required"`
	Itinerary          [][]Itinerary `json:"itinerary,omitempty"`
	Geometry           []string      `json:"geometry"`
	Expenses           Expenses      `json:"expenses,omitempty"`
}

type Itinerary struct {
	Date  time.Time        `json:"date"`
	Place citydomain.Place `json:"place"`
}

type Schedule struct {
	StartDate     time.Time `json:"startDate"`
	EndDate       time.Time `json:"endDate"`
	StartDateTime time.Time `json:"startDateTime"`
	EndDateTime   time.Time `json:"endDateTime"`
	StartTime     time.Time `json:"startTime"`
	EndTime       time.Time `json:"endTime"`
	TotalDays     uint8     `json:"totalDays,omitempty"`
}

type Expenses struct {
	Items []Expense `json:"items"`
	Total float32   `json:"total"`
}

type Expense struct {
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Amount      float32 `json:"amount"`
}
