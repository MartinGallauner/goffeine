package repl

import "fmt"

func commandHelp(cfg *Config, parameter string) error {
	fmt.Println("Welcome to the Goffeine!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s \n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil

}
