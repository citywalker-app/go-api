package traveldomain

import (
	"math"
	"math/rand"

	citydomain "github.com/citywalker-app/go-api/pkg/city/domain"
)

type Cluster struct {
	CentroID     [][]float64
	Clusters     [][]citydomain.Place
	ClusterAlloc [][]float64
	NClusters    uint8
}

func (c *Cluster) Create(placesToVisit *[]citydomain.Place, matrixCost *MatrixCost) {
	// Calculate num of clusters
	c.NClusters = uint8(math.Round(math.Sqrt(float64(len(*placesToVisit))))) + 1

	c.Initialize(len(*placesToVisit), matrixCost)

	isChange := false

	for !isChange {
		c.Fill(placesToVisit, &isChange)
		if !isChange {
			break
		}
		c.RedefineCentroID()
		isChange = false
	}

	filteredCluster := make([][]citydomain.Place, 0, len(c.Clusters))
	filteredCentroID := make([][]float64, 0, len(c.CentroID))

	for i := range c.Clusters {
		if len(c.Clusters[i]) > 0 {
			filteredCluster = append(filteredCluster, c.Clusters[i])
			filteredCentroID = append(filteredCentroID, c.CentroID[i])
		}
	}

	c.Clusters = filteredCluster
	c.CentroID = filteredCentroID
}

func (c *Cluster) Fill(visit *[]citydomain.Place, isChange *bool) {
	for i, place := range *visit {
		var cl uint8
		min := 100.0
		var j uint8
		for j = 0; j < c.NClusters; j++ {
			c.ClusterAlloc[i][j] = place.CalcDistance(c.CentroID[j][0], c.CentroID[j][1])
			if c.ClusterAlloc[i][j] < min {
				min = c.ClusterAlloc[i][j]
				cl = j
			}
		}
		if c.ClusterAlloc[i][c.NClusters] != float64(cl) {
			c.ClusterAlloc[i][c.NClusters] = float64(cl)
			*isChange = true
		}
		c.Clusters[cl] = append(c.Clusters[cl], place)
	}
}

func (c *Cluster) Initialize(lenVisit int, matrixCost *MatrixCost) {
	// center of clusters
	c.CentroID = make([][]float64, c.NClusters)

	// cluster to store places
	c.Clusters = make([][]citydomain.Place, c.NClusters)

	// cost of each cluster
	c.ClusterAlloc = make([][]float64, lenVisit)

	var i uint8
	for i = 0; i < c.NClusters; i++ {
		c.CentroID[i] = make([]float64, 2)
		c.Clusters[i] = make([]citydomain.Place, 0, lenVisit)
		c.CentroID[i][0], c.CentroID[i][1] = generateCentroID(matrixCost)
	}

	for i := 0; i < lenVisit; i++ {
		c.ClusterAlloc[i] = make([]float64, c.NClusters+1)
		c.ClusterAlloc[i][c.NClusters] = 100.0
	}
}

// nolint:gosec
func generateCentroID(m *MatrixCost) (float64, float64) {
	centroID0 := m.MinLatitude*1.001 + rand.Float64()*(m.MaxLatitude*0.999-m.MinLatitude*1.001)
	centroID1 := m.MinLongitude*1.001 + rand.Float64()*(m.MaxLongitude*0.999-m.MinLongitude*1.001)
	return centroID0, centroID1
}

func (c *Cluster) RedefineCentroID() {
	for i := range c.Clusters {
		if len(c.Clusters[i]) == 0 {
			continue
		}
		maxLong := -180.0
		maxLatitude := -90.0
		minLong := 180.00
		minLatitude := 90.0
		c.CentroID[i][0] = 0
		c.CentroID[i][1] = 0
		for _, place := range c.Clusters[i] {
			lat := place.Location.Coordinates[0]
			long := place.Location.Coordinates[1]
			if lat > maxLatitude {
				maxLatitude = lat
			}
			if lat < minLatitude {
				minLatitude = lat
			}
			if long > maxLong {
				maxLong = long
			}
			if long < minLong {
				minLong = long
			}
			c.CentroID[i][0] += lat
			c.CentroID[i][1] += long
		}
		c.CentroID[i][0] /= float64(len(c.Clusters[i]))
		c.CentroID[i][1] /= float64(len(c.Clusters[i]))
		c.Clusters[i] = c.Clusters[i][:0]
	}
}
