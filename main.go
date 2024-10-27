package main

import (
	"log"
	"net/http"
	"time"
	"github.com/MartinGallauner/goffeine/internal/askopenai"
	"github.com/MartinGallauner/goffeine/internal/server"
	"github.com/MartinGallauner/goffeine/internal/tracker"
	"github.com/alexedwards/scs/v2"
	"github.com/joho/godotenv"
)


func main() {
	log.Println("Started Goffeine")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed trying to load env variables.")
	}

	repository := tracker.NewMemoryRepository()
	client := askopenai.New()
	
	t := tracker.New(repository, client)
	
	
	sessionManager := scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	server := server.NewGoffeineServer(t, sessionManager)
	
	log.Fatal(http.ListenAndServe(":5001", server))
}