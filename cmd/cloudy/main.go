package main

import (
	"os"

	"github.com/blue-monads/potatoverse/backend/distro/cloudy"
)

func main() {

	masterSecret := os.Getenv("CLOUDY_MASTER_SECRET")
	if masterSecret == "" {
		panic("CLOUDY_MASTER_SECRET env variable is required")
	}

	tursoAuthToken := os.Getenv("TURSO_AUTH_TOKEN")
	if tursoAuthToken == "" {
		panic("TURSO_AUTH_TOKEN env variable is required")
	}

	tursoRemoteURL := os.Getenv("TURSO_REMOTE_URL")
	if tursoRemoteURL == "" {
		panic("TURSO_REMOTE_URL env variable is required")
	}

	capp, err := cloudy.New(&cloudy.Config{
		Port:           8080,
		WorkingDir:     "./cloudy_data",
		MasterSecret:   masterSecret,
		TursoAuthToken: tursoAuthToken,
		TursoRemoteURL: tursoRemoteURL,
	})
	if err != nil {
		panic(err)
	}

	if err := capp.Build(); err != nil {
		panic(err)
	}

	if err := capp.Run(); err != nil {
		panic(err)
	}

}
