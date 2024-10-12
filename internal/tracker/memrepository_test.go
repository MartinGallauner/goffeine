package tracker

import (
	"testing"
	"time"
)

func TestAdd(t *testing.T) {

	repository := NewMemoryRepository()

	time := time.Now()
	repository.Add(time, 100)

	if len(repository.Entries[1]) != 1  {
		t.Fatal("Adding a new entry failed")
	}

}
