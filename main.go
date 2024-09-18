package main

import (
	"github.com/MartinGallauner/goffeine/internal/askopenai"
	"github.com/MartinGallauner/goffeine/internal/repl"
	"github.com/MartinGallauner/goffeine/internal/repository"
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

	csvRepository := repository.New("data/data.csv")
	client := askopenai.New()

	t := tracker.New(csvRepository, client)
	config := &repl.Config{Tracker: *t}

	handler := http.HandlerFunc(server.GoffeineServer)
	log.Fatal(http.ListenAndServe(":5000", handler))

	repl.StartRepl(config)
}
