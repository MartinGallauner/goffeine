package tracker

import (
	"github.com/MartinGallauner/goffeine/internal/askopenai"
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

func (r *TestRepository) Add(timestamp time.Time, caffeineInMg int) error {
	entries := append(r.entries, repository.Entry{
		Timestamp:    timestamp,
		CaffeineInMg: caffeineInMg,
	})
	r.entries = entries
	return nil
}

type MockClient struct {
}

func (c *MockClient) AskLlm(input string) (askopenai.CaffeineIntake, error) {
	return askopenai.CaffeineIntake{
		Timestamp:    time.Time{},
		CaffeineInMg: 100,
	}, nil

}

func TestLevelIsZero(t *testing.T) {
	tracker := New(&TestRepository{}, &MockClient{})
	timestamp := time.Date(2024, time.August, 26, 11, 53, 25, 0, time.UTC)
	caffeineLevel, _ := tracker.GetLevel(timestamp)

	if caffeineLevel != 0 {
		t.Errorf("Expected zero but got '%v'", caffeineLevel)
	}
}

func TestAddCaffeineSameMoment(t *testing.T) {
	tracker := New(&TestRepository{})
	timestamp := time.Date(2024, time.August, 26, 11, 53, 25, 0, time.UTC)
	tracker.Add(timestamp, 100)
	caffeineLevel, _ := tracker.GetLevel(timestamp)

	if caffeineLevel != 100 {
		t.Errorf("Expected 100 but got '%v'", caffeineLevel)
	}
}

func TestAddCaffeineHalfLife(t *testing.T) {
	tracker := New(&TestRepository{})
	addTime := time.Date(2024, time.August, 26, 11, 53, 25, 0, time.UTC)
	tracker.Add(addTime, 100)

	checkTime := addTime.Add(time.Minute * 300)

	caffeineLevel, _ := tracker.GetLevel(checkTime)

	if caffeineLevel != 50 {
		t.Errorf("Expected half-life value of 50mg but got '%vmg'", caffeineLevel)
	}
}
