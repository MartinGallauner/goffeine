package main

import (
	"github.com/MartinGallauner/goffeine/internal/llmclient"
	"github.com/MartinGallauner/goffeine/internal/repl"
	"github.com/MartinGallauner/goffeine/internal/repository"
	"github.com/MartinGallauner/goffeine/internal/tracker"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	log.Println("Started Goffeine")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed trying to load env variables.")
	}

	csvRepository := repository.New("data/data.csv")
	client := llmclient.New()

	t := tracker.New(csvRepository, *client)
	config := &repl.Config{Tracker: *t}
	repl.StartRepl(config)
}
