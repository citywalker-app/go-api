package traveldomain

import citydomain "github.com/citywalker-app/go-api/pkg/city/domain"

type Distance struct {
	Place    citydomain.Place
	Distance float32
}

type ByDistance []Distance

func (b ByDistance) Len() int {
	return len(b)
}

func (b ByDistance) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b ByDistance) Less(i, j int) bool {
	return b[i].Distance < b[j].Distance
}
