package main

import (
	"fmt"
	"github.com/MartinGallauner/goffeine/internal/ask"
	"github.com/MartinGallauner/goffeine/internal/server"
	"github.com/MartinGallauner/goffeine/internal/tracker"
	"github.com/joho/godotenv"
	"log"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil { //TODO I need to figure out how to handle .env files in deployment
		log.Print("Failed trying to load the .env file.")
	}

	repository := tracker.NewMemoryRepository()
	client := ask.New()
	t := tracker.New(repository, client)
	goffeineServer := server.NewGoffeineServer(t)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	addr := fmt.Sprintf(":%s", port)
	slog.Info("Started Goffeine on:", "port", port)
	log.Fatal(http.ListenAndServe(addr, goffeineServer)) /* #nosec */
}
