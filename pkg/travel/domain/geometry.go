package traveldomain

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/citywalker-app/go-api/utils"
)

type Response struct {
	Routes []Routes `json:"routes"`
}

type Routes struct {
	Geometry string `json:"geometry"`
}

func (t *Travel) GetGeometry() error {
	baseURL := "https://routing.openstreetmap.de/routed-foot/route/v1/walking/"
	t.Geometry = make([]string, t.Schedule.TotalDays+1)

	for i, day := range t.Itinerary {
		var coords []string
		for _, it := range day {
			coords = append(coords, utils.Float64SliceToString(it.Place.Location.Coordinates))
		}

		coordString := strings.Join(coords, ";")

		u, err := url.Parse(baseURL)
		if err != nil {
			return err
		}

		u.Path = fmt.Sprintf("%s%s", u.Path, coordString)

		resp, err := http.Get(u.String())
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		var response Response

		err = json.Unmarshal(body, &response)
		if err != nil {
			return err
		}

		t.Geometry[i] = response.Routes[0].Geometry
	}

	return nil
}
