package repl

import (
	"strconv"
	"time"
)

func commandAdd(cfg *Config, parameter string) error {
	num, err := strconv.Atoi(parameter)
	if err != nil {
		return err
	}

	err = cfg.Tracker.Add(time.Now(), num)
	if err != nil {
		return err
	}

	return nil

}
