package cityhandler

type TestCase struct {
	Name       string
	URL        string
	StatusCode int
}

var GetCitiesTestCases = []TestCase{
	{
		Name:       "Get all cities in english(success)",
		URL:        "/cities/all/en",
		StatusCode: 200,
	},
}

var GetCityTestCases = []TestCase{
	{
		Name:       "Get city by name(success)",
		URL:        "/cities/london",
		StatusCode: 200,
	},
	{
		Name:       "Get city by name(fail)",
		URL:        "/cities/berlin",
		StatusCode: 404,
	},
}
