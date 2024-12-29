package main

import (
	"github.com/MartinGallauner/goffeine/internal/askopenai"
	"github.com/MartinGallauner/goffeine/internal/server"
	"github.com/MartinGallauner/goffeine/internal/tracker"
	"github.com/alexedwards/scs/v2"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"time"
)

func main() {
	log.Println("Started Goffeine")

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
	server := server.NewGoffeineServer(t, sessionManager)

	log.Fatal(http.ListenAndServe(":8080", server))
}
