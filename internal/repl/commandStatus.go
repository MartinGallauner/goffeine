package repl

import (
	"time"
)

func commandStatus(cfg *Config, parameter string) error {
	_, err := cfg.Tracker.GetLevel(time.Now().UTC())
	if err != nil {
		return err
	}
	return nil

}
