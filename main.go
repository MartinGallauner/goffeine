package main

import (
	"github.com/MartinGallauner/goffeine/internal/repl"
	"github.com/MartinGallauner/goffeine/internal/repository"
	"github.com/MartinGallauner/goffeine/internal/tracker"
	"log"
)

func main() {
	log.Println("Started Goffeine")

	csvRepository := repository.New("data/data.csv")
	t := tracker.New(csvRepository)
	config := &repl.Config{Tracker: *t}
	repl.StartRepl(config)
}
