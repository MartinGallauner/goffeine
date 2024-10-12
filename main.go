package main

import (
	"github.com/MartinGallauner/goffeine/internal/askopenai"
	"github.com/MartinGallauner/goffeine/internal/server"
	"github.com/MartinGallauner/goffeine/internal/tracker"
	"github.com/joho/godotenv"
	"log"
	"net/http"
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
	server := server.NewGoffeineServer(t)
	log.Fatal(http.ListenAndServe(":5001", server))
}