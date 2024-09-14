package userhandler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"testing"

	"github.com/citywalker-app/go-api/database"
	_ "github.com/citywalker-app/go-api/envLoader"
	userapplication "github.com/citywalker-app/go-api/pkg/user/application"
	userhandler "github.com/citywalker-app/go-api/pkg/user/infrastructure/handler"
	"github.com/citywalker-app/go-api/pkg/user/infrastructure/persistence/mongo"
	"github.com/citywalker-app/go-api/server"
	"github.com/gofiber/fiber/v2"
	"gopkg.in/go-playground/assert.v1"
)

var app *fiber.App

func init() {
	os.Setenv("MDB_COLLECTION_USERS", "users")

	app = server.Setup()

	var db database.MongoDB
	db.ConnectTest()
	database.DB = &db

	userapplication.Repo = mongo.NewMongoRepository()
}

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func GenericCheck(t *testing.T, tc *userhandler.TestCase, route string) *http.Response {
	userJSON, _ := json.Marshal(tc.User)
	req, err := http.NewRequest(http.MethodPost, route, bytes.NewBuffer(userJSON))
	req.Header.Set("Content-Type", "application/json")

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

func TestRegister(t *testing.T) {
	for _, tc := range userhandler.RegisterTestCases {
		t.Run(tc.Name, func(t *testing.T) {
			resp := GenericCheck(t, &tc, "/user/register")
			if resp == nil {
				return
			}
			defer resp.Body.Close()

			var response userhandler.Response
			if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
				t.Fatalf("Error decoding response body: %v", err)
			}

			assert.Equal(t, response.User.Email, tc.User.Email)
			assert.NotEqual(t, response.User.Password, tc.User.Password)
			assert.NotEqual(t, response.JWT, "")
		})
	}
}

func TestLogin(t *testing.T) {
	for _, tc := range userhandler.LoginTestCases {
		t.Run(tc.Name, func(t *testing.T) {
			resp := GenericCheck(t, &tc, "/user/login")
			if resp == nil {
				return
			}
			defer resp.Body.Close()

			var response userhandler.Response
			if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
				t.Fatalf("Error decoding response body: %v", err)
			}

			assert.NotEqual(t, response.JWT, "")
		})
	}
}

func TestResetPassword(t *testing.T) {
	for _, tc := range userhandler.ResetPasswordTestCases {
		t.Run(tc.Name, func(t *testing.T) {
			resp := GenericCheck(t, &tc, "/user/resetPassword")
			if resp == nil {
				return
			}
			defer resp.Body.Close()
		})
	}
}

func TestCanLoginWithNewPassword(t *testing.T) {
	lenTestCases := len(userhandler.ResetPasswordTestCases)
	user := userhandler.ResetPasswordTestCases[lenTestCases-1].User

	userJSON, _ := json.Marshal(user)
	req, err := http.NewRequest(http.MethodPost, "/user/login", bytes.NewBuffer(userJSON))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		t.Fatal(err)
	}

	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
	defer resp.Body.Close()

	assert.Equal(t, resp.StatusCode, 200)

	var response userhandler.Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Fatalf("Error decoding response body: %v", err)
	}

	assert.NotEqual(t, response.JWT, "")
}

// nolint:dupl
func TestConfirmCode(t *testing.T) {
	for _, tc := range userhandler.ConfirmCodeTestCases {
		t.Run(tc.Name, func(t *testing.T) {
			resp := GenericCheck(t, &tc, "/user/confirmCode")
			if resp == nil {
				return
			}
			defer resp.Body.Close()

			var response userhandler.Response
			if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
				t.Fatalf("Error decoding response body: %v", err)
			}

			assert.NotEqual(t, response.ConfirmCode, "")
		})
	}
}

func TestContinueWithGoogle(t *testing.T) {
	for _, tc := range userhandler.ContinueWithGoogleTestCases {
		t.Run(tc.Name, func(t *testing.T) {
			resp := GenericCheck(t, &tc, "/user/continueWithGoogle")
			if resp == nil {
				return
			}
			defer resp.Body.Close()

			var response userhandler.Response
			if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
				t.Fatalf("Error decoding response body: %v", err)
			}

			assert.NotEqual(t, response.JWT, "")
		})
	}
}
