package cityhandler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"os"
	"testing"

	"github.com/citywalker-app/go-api/database"
	cityapplication "github.com/citywalker-app/go-api/pkg/city/application"
	citydomain "github.com/citywalker-app/go-api/pkg/city/domain"
	cityhandler "github.com/citywalker-app/go-api/pkg/city/infrastructure/handler"
	citymongo "github.com/citywalker-app/go-api/pkg/city/infrastructure/persistence/mongo"
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

func init() {
	os.Setenv("MDB_COLLECTION_CITIES", "cities")
	os.Setenv("MDB_COLLECTION_USERS", "users")
	os.Setenv("REDIS_ADDR", "localhost:6379")
	os.Setenv("REDIS_DB", "0")
	os.Setenv("REDIS_PASSWORD", "")

	app = server.Setup()

	var db database.MongoDB
	db.ConnectTest()
	database.DB = &db

	cityapplication.Repo = citymongo.NewMongoRepository()
	userapplication.Repo = usermongo.NewMongoRepository()

	jwt = login()
}

func login() string {
	user := userdomain.User{
		Email:    "testcity@gmail.com",
		FullName: "Test",
		Password: "12345689",
	}

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

func TestGetCities(t *testing.T) {
	for _, tc := range cityhandler.GetCitiesTestCases {
		t.Run(tc.Name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, tc.URL, nil)
			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}
			req.Header.Set("Authorization", "Bearer "+jwt)

			resp, err := app.Test(req, -1)
			if err != nil {
				t.Fatalf("could not send request: %v", err)
			}
			defer resp.Body.Close()

			assert.Equal(t, tc.StatusCode, resp.StatusCode)

			var response cityhandler.Response
			if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
				t.Fatalf("could not decode response: %v", err)
			}

			emptyArray := make([]citydomain.City, 0)
			assert.NotEqual(t, len(response.Cities), len(emptyArray))
		})
	}
}

func TestCacheGetCities(t *testing.T) {
	for _, tc := range cityhandler.GetCitiesTestCases {
		t.Run(tc.Name, func(t *testing.T) {
			if tc.StatusCode == http.StatusOK {
				var value []citydomain.City
				err := utils.GetCache(context.Background(), tc.URL, &value)
				if err != nil {
					t.Fatalf("could not get cache: %v", err)
				}

				emptyArray := make([]citydomain.City, 0)
				assert.NotEqual(t, len(value), len(emptyArray))
			}
		})
	}
}

func TestGetCity(t *testing.T) {
	for _, tc := range cityhandler.GetCityTestCases {
		t.Run(tc.Name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, tc.URL, nil)
			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}
			req.Header.Set("Authorization", "Bearer "+jwt)

			resp, err := app.Test(req, -1)
			if err != nil {
				t.Fatalf("could not send request: %v", err)
			}
			defer resp.Body.Close()

			assert.Equal(t, tc.StatusCode, resp.StatusCode)

			var response cityhandler.Response
			if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
				t.Fatalf("could not decode response: %v", err)
			}

			assert.NotEqual(t, response.City.ID, 0)
		})
	}
}

func TestCacheGetCity(t *testing.T) {
	for _, tc := range cityhandler.GetCityTestCases {
		t.Run(tc.Name, func(t *testing.T) {
			if tc.StatusCode == http.StatusOK {
				var value citydomain.City
				err := utils.GetCache(context.Background(), tc.URL, &value)
				if err != nil {
					t.Fatalf("could not get cache: %v", err)
				}

				var emptyCity citydomain.City
				assert.NotEqual(t, emptyCity, value)
			}
		})
	}
}
