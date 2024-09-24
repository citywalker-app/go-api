package travelhandler_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/citywalker-app/go-api/database"
	cityapplication "github.com/citywalker-app/go-api/pkg/city/application"
	citymongo "github.com/citywalker-app/go-api/pkg/city/infrastructure/persistence/mongo"
	travelapplication "github.com/citywalker-app/go-api/pkg/travel/application"
	traveldomain "github.com/citywalker-app/go-api/pkg/travel/domain"
	travelhandler "github.com/citywalker-app/go-api/pkg/travel/infrastructure/handler"
	travelmongo "github.com/citywalker-app/go-api/pkg/travel/infrastructure/persistence/mongo"
	userapplication "github.com/citywalker-app/go-api/pkg/user/application"
	userdomain "github.com/citywalker-app/go-api/pkg/user/domain"
	userhandler "github.com/citywalker-app/go-api/pkg/user/infrastructure/handler"
	usermongo "github.com/citywalker-app/go-api/pkg/user/infrastructure/persistence/mongo"
	"github.com/citywalker-app/go-api/server"
	"github.com/citywalker-app/go-api/utils"
	"github.com/gofiber/fiber/v2"
	"gopkg.in/go-playground/assert.v1"
)

var (
	app *fiber.App
	jwt string
)

var user = userdomain.User{
	Email:    "testtravel@gmail.com",
	FullName: "Test",
	Password: "12345689",
}

func init() {
	os.Setenv("MDB_COLLECTION_TRAVELS", "travels")
	os.Setenv("MDB_COLLECTION_USERS", "users")
	os.Setenv("MDB_COLLECTION_CITIES", "cities")
	os.Setenv("REDIS_ADDR", "localhost:6379")
	os.Setenv("REDIS_DB", "0")
	os.Setenv("REDIS_PASSWORD", "")

	app = server.Setup()

	var db database.MongoDB
	db.ConnectTest()
	database.DB = &db

	cityapplication.Repo = citymongo.NewMongoRepository()
	userapplication.Repo = usermongo.NewMongoRepository()
	travelapplication.Repo = travelmongo.NewMongoRepository()

	jwt = register()
}

func register() string {
	userJSON, _ := json.Marshal(user)
	req, err := http.NewRequest(http.MethodPost, "/user/register", bytes.NewBuffer(userJSON))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var response userhandler.Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		panic(err)
	}
	return response.JWT
}

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func GenericCheck(t *testing.T, tc *travelhandler.TestCase) *http.Response {
	travelJSON, _ := json.Marshal(tc.Travel)
	req, err := http.NewRequest(http.MethodPost, "/travels/create", bytes.NewBuffer(travelJSON))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+jwt)

	if err != nil {
		t.Fatal(err)
	}

	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	assert.Equal(t, resp.StatusCode, tc.StatusCode)

	if resp.StatusCode != http.StatusOK {
		return nil
	}

	return resp
}

func IsIncluidingCategory(travel *traveldomain.Travel) bool {
	for _, category := range travel.ExcludedCategories {
		for _, day := range travel.Itinerary {
			for _, place := range day {
				if utils.Includes(travel.MustVisitPlaces, place.Place.Name) {
					continue
				}
				if place.Place.Category == category {
					return true
				}
			}
		}
	}
	return false
}

func IsIncluidingPlace(travel *traveldomain.Travel) bool {
	for _, placeMustVisit := range travel.MustVisitPlaces {
		for _, day := range travel.Itinerary {
			for _, place := range day {
				fmt.Println(place.Place.Name, "-", placeMustVisit)
				if place.Place.Name == placeMustVisit {
					return true
				}
			}
		}
	}
	return false
}

func TestCreate(t *testing.T) {
	for _, tc := range travelhandler.CreateTestCases {
		t.Run(tc.Name, func(t *testing.T) {
			resp := GenericCheck(t, &tc)
			if resp == nil {
				return
			}
			defer resp.Body.Close()

			var response travelhandler.Response
			if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
				t.Fatalf("Error decoding response body: %v", err)
			}

			emptyItinerary := make([][]traveldomain.Itinerary, 0)
			assert.NotEqual(t, response.Travel.Itinerary, emptyItinerary)
			assert.Equal(t, response.Travel.Schedule.TotalDays, uint8(3))
			assert.Equal(t, false, IsIncluidingCategory(&response.Travel))
			assert.Equal(t, true, IsIncluidingPlace(&response.Travel))
		})
	}
}

func TestTravelAddedtoUser(t *testing.T) {
	userJSON, _ := json.Marshal(user)
	req, err := http.NewRequest(http.MethodPost, "/user/login", bytes.NewBuffer(userJSON))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var response userhandler.Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		panic(err)
	}
	emptyArray := make([]string, 0)
	assert.NotEqual(t, len(emptyArray), len(response.User.Travels))
}
