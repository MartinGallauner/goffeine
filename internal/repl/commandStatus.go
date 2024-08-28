package repl

import (
	"time"
)

func commandStatus(cfg *Config, parameter string) error {
	cfg.Tracker.GetLevel(time.Now())
	return nil

}
