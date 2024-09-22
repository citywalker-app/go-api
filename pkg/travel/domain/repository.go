package traveldomain

type Repository interface {
	Create(travel *Travel) error
}
