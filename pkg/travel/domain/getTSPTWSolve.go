package traveldomain

import (
	citydomain "github.com/citywalker-app/go-api/pkg/city/domain"
)

func (t *Travel) GetTSPTWSolve(places *[]citydomain.Place) error {
	matrixCost := GetMatrixCost(places)

	// getting clusters
	var cluster Cluster
	cluster.Create(places, matrixCost)

	// getting path
	path, err := CreatePath(&cluster, matrixCost)
	if err != nil {
		return err
	}

	CreateJourney(t, path, matrixCost)

	if err := t.GetGeometry(); err != nil {
		return err
	}

	return nil
}
