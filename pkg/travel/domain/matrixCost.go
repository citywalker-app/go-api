package traveldomain

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	citydomain "github.com/citywalker-app/go-api/pkg/city/domain"
	"github.com/citywalker-app/go-api/utils"
)

type MatrixCost struct {
	PlacesMapping map[string]int `json:"placesMapping"`
	Durations     [][]float32    `json:"durations"`
	MinLatitude   float64        `json:"minLatitude"`
	MinLongitude  float64        `json:"minLongitude"`
	MaxLatitude   float64        `json:"maxLatitude"`
	MaxLongitude  float64        `json:"maxLongitude"`
}

func NewMatrixCost() *MatrixCost {
	return &MatrixCost{
		Durations:     [][]float32{},
		PlacesMapping: map[string]int{},
		MinLatitude:   90,
		MinLongitude:  180,
		MaxLatitude:   -90,
		MaxLongitude:  -180,
	}
}

func GetMatrixCost(places *[]citydomain.Place) *MatrixCost {
	baseURL := "https://routing.openstreetmap.de/routed-foot/table/v1/walking/"
	response := NewMatrixCost()

	coords := make([]string, 0, len(*places))
	for i, place := range *places {
		response.PlacesMapping[place.Name] = i
		coords = append(coords, utils.Float64SliceToString(place.Location.Coordinates))
	}

	coordString := strings.Join(coords, ";")

	u, err := url.Parse(baseURL)
	if err != nil {
		return nil
	}

	u.Path = fmt.Sprintf("%s%s", u.Path, coordString)

	resp, err := http.Get(u.String())
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil
	}

	response.GetMinAndMax(places)

	return response
}

func (m *MatrixCost) GetMinAndMax(places *[]citydomain.Place) {
	for _, place := range *places {
		if place.Location.Coordinates[0] < m.MinLatitude {
			m.MinLatitude = place.Location.Coordinates[0]
		}
		if place.Location.Coordinates[0] > m.MaxLatitude {
			m.MaxLatitude = place.Location.Coordinates[0]
		}
		if place.Location.Coordinates[1] < m.MinLongitude {
			m.MinLongitude = place.Location.Coordinates[1]
		}
		if place.Location.Coordinates[1] > m.MaxLongitude {
			m.MaxLongitude = place.Location.Coordinates[1]
		}
	}
}

func (m *MatrixCost) GetIndex(name string) int {
	return m.PlacesMapping[name]
}
