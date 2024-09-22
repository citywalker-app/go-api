package traveldomain

import (
	"math"
	"sort"

	citydomain "github.com/citywalker-app/go-api/pkg/city/domain"
)

func CreatePath(cluster *Cluster, matrixCost *MatrixCost) (*[]citydomain.Place, error) {
	initialPath := InitialPath(&cluster.Clusters, matrixCost)

	var finalPath []citydomain.Place
	if len(*initialPath) == 0 {
		return nil, ErrNotEnoughPlaces
	}

	finalPath = append(finalPath, (*initialPath)[0]...)
	(*initialPath) = (*initialPath)[1:]

	for range *initialPath {
		cl, posAdd, posCloser := findClosestCluster(initialPath, finalPath, matrixCost)

		shouldReverseInitialSegment := (posAdd == 0 && posCloser == 0) || (posAdd != 0 && posCloser != 0)
		pathSegment := (*initialPath)[cl]

		if shouldReverseInitialSegment {
			pathSegment = reverse(pathSegment)
		}

		if posAdd == 0 {
			finalPath = append(pathSegment, finalPath...)
		} else {
			finalPath = append(finalPath, pathSegment...)
		}

		finalPath = OptimizePath(finalPath, matrixCost)

		(*initialPath) = append((*initialPath)[:cl], (*initialPath)[cl+1:]...)
	}

	return &finalPath, nil
}

// nolint:lll
func findClosestCluster(initialPath *[][]citydomain.Place, finalPath []citydomain.Place, matrixCost *MatrixCost) (int, int, int) {
	var min float32 = 1000.0
	cl := 0
	posAdd := 0
	posCloser := 0

	for i := 0; i < len(*initialPath); i++ {
		finalIndex := matrixCost.GetIndex(finalPath[0].Name)
		initialIndex := matrixCost.GetIndex((*initialPath)[i][0].Name)
		lastInitialIndex := matrixCost.GetIndex((*initialPath)[i][len((*initialPath)[i])-1].Name)
		lastFinalIndex := matrixCost.GetIndex(finalPath[len(finalPath)-1].Name)

		scenarios := []struct {
			cost         float32
			newPosAdd    int
			newPosCloser int
		}{
			{matrixCost.Durations[finalIndex][initialIndex], 0, 0},
			{matrixCost.Durations[finalIndex][lastInitialIndex], 0, len((*initialPath)[i]) - 1},
			{matrixCost.Durations[lastFinalIndex][initialIndex], len(finalPath) - 1, 0},
			{matrixCost.Durations[lastFinalIndex][lastInitialIndex], len(finalPath) - 1, len((*initialPath)[i]) - 1},
		}

		for _, scenario := range scenarios {
			if scenario.cost < min {
				min = scenario.cost
				cl = i
				posAdd = scenario.newPosAdd
				posCloser = scenario.newPosCloser
			}
		}
	}

	return cl, posAdd, posCloser
}

func OptimizePath(path []citydomain.Place, matrixCost *MatrixCost) []citydomain.Place {
	improvement := true
	for improvement {
		improvement = false
		for i := 0; i < len(path)-1; i++ {
			for k := i + 1; k < len(path); k++ {
				newPath := twoOptSwap(path, i, k)
				if calculateTotalDistance(newPath, matrixCost) < calculateTotalDistance(path, matrixCost) {
					path = newPath
					improvement = true
				}
			}
		}
	}
	return path
}

func calculateTotalDistance(path []citydomain.Place, matrixCost *MatrixCost) float32 {
	var totalDistance float32
	for i := 0; i < len(path)-1; i++ {
		totalDistance += matrixCost.Durations[matrixCost.GetIndex(path[i].Name)][matrixCost.GetIndex(path[i+1].Name)]
	}
	return totalDistance
}

func twoOptSwap(path []citydomain.Place, i, k int) []citydomain.Place {
	var newPath []citydomain.Place
	// 1. Take route[0] to route[i-1] and add them in order to newPath
	for j := 0; j <= i-1; j++ {
		newPath = append(newPath, path[j])
	}
	// 2. Take route[i] to route[k] and add them in reverse order to newPath
	for j := k; j >= i; j-- {
		newPath = append(newPath, path[j])
	}
	// 3. Take route[k+1] to end and add them in order to newPath
	for j := k + 1; j < len(path); j++ {
		newPath = append(newPath, path[j])
	}
	return newPath
}

func reverse(s []citydomain.Place) []citydomain.Place {
	var newS []citydomain.Place
	for i := len(s) - 1; i >= 0; i-- {
		newS = append(newS, s[i])
	}
	return newS
}

// nolint:lll
func InitialPath(clusters *[][]citydomain.Place, matrixCost *MatrixCost) *[][]citydomain.Place {
	distances := CalculateClusterDistances(clusters, matrixCost)

	initialPath := make([][]citydomain.Place, len(*distances))

	for i := range *distances {
		initialPath[i] = make([]citydomain.Place, 0, 10)

		switch len((*distances)[i]) {
		case 1:
			initialPath[i] = append(initialPath[i], (*distances)[i][0].Place)
			continue
		case 2:
			initialPath[i] = append(initialPath[i], (*distances)[i][0].Place, (*distances)[i][1].Place)
			continue
		default:
			initialPath[i] = append(initialPath[i], (*distances)[i][0].Place, (*distances)[i][1].Place)
		}

		for j := 2; j < len((*distances)[i]); j++ {
			var min float32 = 100000.0
			pos := -1
			p := (*distances)[i][j].Place

			for k := 0; k < len(initialPath[i])+1; k++ {
				var cost float32
				for d := 0; d < len(initialPath[i]); d++ {
					iPlace := matrixCost.GetIndex(p.Name)
					iPlaceID := matrixCost.GetIndex(initialPath[i][d].Name)

					switch {
					case k == d:
						cost += matrixCost.Durations[iPlace][iPlaceID]
					case k-d == 1:
						cost += matrixCost.Durations[iPlaceID][iPlace]
					case k-d < 0:
						iPlaceIDBefore := matrixCost.GetIndex(initialPath[i][d-1].Name)
						cost += matrixCost.Durations[iPlaceID][iPlaceIDBefore]
					default:
						iPlaceIDAfter := matrixCost.GetIndex(initialPath[i][d+1].Name)
						cost += matrixCost.Durations[iPlaceID][iPlaceIDAfter]
					}
				}
				if cost < min {
					min = cost
					pos = k
				}
			}

			if pos >= len(initialPath[i]) {
				initialPath[i] = append(initialPath[i], p)
			} else {
				initialPath[i] = append(initialPath[i][:pos+1], initialPath[i][pos:]...)
				initialPath[i][pos] = p
			}
		}
	}

	return &initialPath
}

func CalculateClusterDistances(clusters *[][]citydomain.Place, matrixCost *MatrixCost) *[][]Distance {
	distance := make([][]Distance, len(*clusters))

	for i := range *clusters {
		for j, place := range (*clusters)[i] {
			d := Distance{
				Place:    place,
				Distance: 0.0,
			}
			for k, secondPlace := range (*clusters)[i] {
				if j != k {
					d.Distance += matrixCost.Durations[matrixCost.GetIndex(place.Name)][matrixCost.GetIndex(secondPlace.Name)]
				}
			}
			distance[i] = append(distance[i], d)
		}
		sort.Sort(ByDistance(distance[i]))
	}

	return &distance
}

func CalcDistance(lat1, lon1, lat2, lon2 float64) float64 {
	PI := 3.1415926535897932384
	lat1 *= (PI / 180)
	lat2 *= (PI / 180)
	lon1 *= (PI / 180)
	lon2 *= (PI / 180)
	return 6371.0 * math.Acos(math.Cos(lat1)*math.Cos(lat2)*math.Cos(lon2-lon1)+math.Sin(lat1)*math.Sin(lat2))
}
