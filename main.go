package main

import (
	"fmt"
	"github.com/MartinGallauner/goffeine/internal/ask"
	"github.com/MartinGallauner/goffeine/internal/server"
	"github.com/MartinGallauner/goffeine/internal/tracker"
	"github.com/alexedwards/scs/v2"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	err := godotenv.Load()
	if err != nil { //todo I need to figure out how to handle that
		log.Print("Failed trying to load the .env file.")
	}

	repository := tracker.NewMemoryRepository()
	client := ask.New()

	t := tracker.New(repository, client)

	sessionManager := scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	goffeineServer := server.NewGoffeineServer(t, sessionManager)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	addr := fmt.Sprintf(":%s", port)
	log.Printf("Started Goffeine on port: %s", port)
	log.Fatal(http.ListenAndServe(addr, goffeineServer)) /* #nosec */
}
