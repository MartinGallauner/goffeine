package repl

import (
	"bufio"
	"fmt"
	"github.com/MartinGallauner/goffeine/internal/tracker"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*Config, string) error
}

func StartRepl(cfg *Config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Goffeine > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		var parameter string
		if len(words) == 2 {
			parameter = words[1]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, parameter)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"status": {
			name:        "status",
			description: "Shows your current caffeine level.",
			callback:    commandStatus,
		},
		"add": {
			name:        "add",
			description: "Adds the passed caffeine beverage.",
			callback:    commandAdd,
		},
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type Config struct {
	Tracker tracker.Tracker
}
