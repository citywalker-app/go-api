package citydomain

type City struct {
	ID          string  `bson:"_id"`
	UUID        string  `json:"uuid"`
	City        string  `json:"city"`
	Continent   string  `json:"continent"`
	Language    string  `json:"lng"`
	LngSpoken   string  `json:"lngSpoken"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Currency    string  `json:"currency"`
	Visa        string  `json:"visa"`
	AdapterPlug string  `json:"adapterPlug"`
	Places      []Place `json:"places"`
	AverageCost uint16  `json:"averageCost"`
}

type Place struct {
	Priority    uint8  `json:"priority"`
	Price       uint8  `json:"price"`
	City        string `json:"city"`
	Lng         string `json:"lng"`
	UUID        string `json:"uuid"`
	Name        string `json:"name"`
	Category    string `json:"category"`
	Information string `json:"information"`
	Description string `json:"description"`
	Website     string `json:"website"`
	Image       struct {
		Credits string `json:"credits"`
		License string `json:"license"`
	} `json:"image"`
	Location struct {
		Type        string    `json:"type"`
		Coordinates []float64 `json:"coordinates"`
	} `json:"location"`
	Visit struct {
		Schedule [][]string `json:"schedule"`
		All      int16      `json:"all"`
		Outside  int16      `json:"outside"`
	} `json:"visit"`
}
