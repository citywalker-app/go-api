package citydomain

type Repository interface {
	GetAll(lng string) (*[]City, error)
	GetCity(city string) (*City, error)
}
