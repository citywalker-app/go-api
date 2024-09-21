package cityhandler

import citydomain "github.com/citywalker-app/go-api/pkg/city/domain"

type Response struct {
	Cities []citydomain.City `json:"cities"`
	City   citydomain.City   `json:"city"`
}
