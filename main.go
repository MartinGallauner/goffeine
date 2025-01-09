package main

import (
	"fmt"
	"github.com/MartinGallauner/goffeine/internal/askopenai"
	"github.com/MartinGallauner/goffeine/internal/server"
	"github.com/MartinGallauner/goffeine/internal/tracker"
	"github.com/alexedwards/scs/v2"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"
)

const DEFAULT_PORT = "8080"

func main() {
	err := godotenv.Load()
	if err != nil { //todo I need to figure out how to handle that
		//log.Fatal("Failed trying to load env variables.")
		log.Print("Failed trying to load the .env file.")
	}

	repository := tracker.NewMemoryRepository()
	client := askopenai.New()

	t := tracker.New(repository, client)

	sessionManager := scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	goffeineServer := server.NewGoffeineServer(t, sessionManager)

	port := os.Getenv("PORT")
	if port == "" {
		log.Printf("No port set. Using %s as default.", DEFAULT_PORT)
		port = DEFAULT_PORT
	}

	log.Printf("Starting Goffeine on port %s", port)
	addr := fmt.Sprintf(":%s", port)
	log.Fatal(http.ListenAndServe(addr, goffeineServer)) /* #nosec */
}
