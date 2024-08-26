package main

import (
	"fmt"
	"github.com/MartinGallauner/goffeine/internal/repository"
	"github.com/MartinGallauner/goffeine/internal/tracker"
	"log"
	"os"
	"strconv"
)

func main() {
	log.Println("Started Goffeine")

	command, num, err := processArgs(os.Args[1:])

	if err != nil {
		log.Fatalf("Unable to handle passed arguments. Got %q", os.Args)
	}

	csvRepository := repository.New("data/data.csv")
	t := tracker.New(csvRepository)

	switch command {
	case "add":
		t.Add(num)
	case "status":
		t.GetLevel()
	}
	log.Println("Finished Goffeine")
}

func processArgs(args []string) (string, int, error) {
	if len(args) < 1 {
		log.Fatal("Please provide an argument")
	}

	command := args[0]

	if command == "add" {

		if len(args) < 2 {
			log.Fatal("Please provide a number as 2nd argument")
		}

		num, err := strconv.Atoi(args[1])
		if err != nil {
			// handle error
			fmt.Println("Error:", err)
			return "", 0, nil
		}
		return command, num, nil
	}
	return command, 0, nil

}
