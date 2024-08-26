package tracker

import (
	"github.com/MartinGallauner/goffeine/internal/repository"
	"testing"
	"time"
)

type TestRepository struct {
	entries []repository.Entry
}

func (r *TestRepository) Fetch() ([]repository.Entry, error) {
	return r.entries, nil
}

func (r *TestRepository) Add(timestamp string, caffeineInMg int) error {
	time, err := time.Parse("2006-01-02T15:04:05", timestamp)
	if err != nil {
		return err
	}
	entries := append(r.entries, repository.Entry{
		Timestamp:    time,
		CaffeineInMg: caffeineInMg,
	})
	r.entries = entries
	return nil
}

func TestLevelIsZero(t *testing.T) {
	tracker := New(&TestRepository{})
	timestamp := time.Date(2024, time.August, 26, 11, 53, 25, 0, time.UTC)
	caffeineLevel, _ := tracker.GetLevel(timestamp)

	if caffeineLevel != 0 {
		t.Errorf("Expected zero but got '%v'", caffeineLevel)
	}
}

func TestAddCaffeine(t *testing.T) {
	tracker := New(&TestRepository{})
	tracker.Add(100)
	timestamp := time.Date(2024, time.August, 26, 11, 53, 25, 0, time.UTC)
	caffeineLevel, _ := tracker.GetLevel(timestamp)

	if caffeineLevel != 100 {
		t.Errorf("Expected 100 but got '%v'", caffeineLevel)
	}
}
