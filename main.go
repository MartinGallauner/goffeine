package main

import (
	"log"
	"os"
)

func main() {
	log.Println("Started Goffeine")

	processArgs(os.Args[1:])

	log.Println("Stopping Goffeine")
}

func processArgs(args []string) string {
	if len(os.Args) < 1 {
		log.Println("Please provide an argument")
		return ""
	}

	command := args[0]
	return command

}
