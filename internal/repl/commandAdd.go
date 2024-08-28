package repl

func commandAdd(cfg *Config, userInput string) error {
	err := cfg.Tracker.Add(userInput)
	if err != nil {
		return err
	}

	return nil

}
