package tracker

import (
	"testing"
	"time"
)

func TestAdd(t *testing.T) {
	repository := NewMemoryRepository()

	repository.Add(time.Now(), 100)
	if len(repository.Entries[1]) != 1  {
		t.Fatal("Adding a new entry failed")
	}

	repository.Add(time.Now(), 50)
	if len(repository.Entries[1]) != 2  {
		t.Fatal("Adding a second entry failed")
	}

	entries, _ := repository.Fetch()

	entryOne := entries[0]
	if entryOne.CaffeineInMg != 100 {
		t.Fatal("intake was persisted incorrectly")
	} 

	entryTwo := entries[1]
	if entryTwo.CaffeineInMg != 50 {
		t.Fatal("intake was persisted incorrectly")
	}
}
