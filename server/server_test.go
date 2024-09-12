// nolint: gci
package server_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"syscall"
	"testing"
	"time"

	_ "github.com/citywalker-app/go-api/envLoader"
	"github.com/citywalker-app/go-api/server"
	_ "github.com/joho/godotenv/autoload"

	"github.com/stretchr/testify/assert"
)

func TestGracefulShutdown(t *testing.T) {
	app := server.Setup()

	go func() {
		server.StartServerWithGracefulShutdown(app)
	}()

	// Give the server a moment to start
	time.Sleep(1 * time.Second)

	// Send a request to the server
	req := httptest.NewRequest(http.MethodGet, "/livez", nil)
	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatalf("Error on sending request: %v", err)
	}
	defer resp.Body.Close()

	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Send an interrupt signal to the server
	p, err := os.FindProcess(os.Getpid())
	assert.NoError(t, err)
	err = p.Signal(syscall.SIGINT)
	assert.NoError(t, err)

	// Give the server a moment to shut down
	time.Sleep(1 * time.Second)

	// Check if the server has shut down
	err = app.Shutdown()
	if err != nil {
		t.Error("Server did not shut down gracefully:", err)
	}
}

func TestStartServer(t *testing.T) {
	// Initialize the Fiber app.
	app := server.Setup()

	t.Run("serverFail", func(t *testing.T) {
		os.Setenv("SERVER_PORT", "abcd")
		// Use a goroutine to start the server since it will block.
		var err error
		go func() {
			err = server.StartServer(app)
		}()

		time.Sleep(1 * time.Second)

		assert.NotNil(t, err)
	})

	t.Run("serverSuccess", func(t *testing.T) {
		// Use a goroutine to start the server since it will block.
		var err error
		go func() {
			err = server.StartServer(app)
		}()

		time.Sleep(1 * time.Second)

		// Check if the server is running by sending a test request.
		req := httptest.NewRequest(http.MethodGet, "/livez", nil)
		resp, err := app.Test(req, -1)
		if err != nil {
			t.Fatalf("Error on sending request: %v", err)
		}
		defer resp.Body.Close()

		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
	})
}
