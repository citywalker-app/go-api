package travelhandler

import traveldomain "github.com/citywalker-app/go-api/pkg/travel/domain"

type Response struct {
	Travel traveldomain.Travel `json:"travel"`
}
