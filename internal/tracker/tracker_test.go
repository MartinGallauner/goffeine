package tracker

import (
	"github.com/MartinGallauner/goffeine/internal/askopenai"
	"testing"
	"time"
)

type TestRepository struct {
	entries []Entry
}

func (r *TestRepository) Fetch() ([]Entry, error) {
	return r.entries, nil
}

func (r *TestRepository) Add(timestamp time.Time, caffeineInMg int) error {
	entries := append(r.entries, Entry{
		Timestamp:    timestamp,
		CaffeineInMg: caffeineInMg,
	})
	r.entries = entries
	return nil
}

type MockClient struct {
}

func (c *MockClient) Ask(input string) (askopenai.CaffeineIntake, error) {
	timestamp := time.Date(2024, time.August, 26, 11, 53, 25, 0, time.UTC)
	return askopenai.CaffeineIntake{
		Timestamp:    timestamp,
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
	tracker := New(&TestRepository{}, &MockClient{})
	timestamp := time.Date(2024, time.August, 26, 11, 53, 25, 0, time.UTC)
	tracker.Add("I'm drinking one big mug of coffee right now.")
	caffeineLevel, _ := tracker.GetLevel(timestamp)

	if caffeineLevel != 100 {
		t.Errorf("Expected 100 but got '%v'", caffeineLevel)
	}
}

func TestAddCaffeineHalfLife(t *testing.T) {
	tracker := New(&TestRepository{}, &MockClient{})
	addTime := time.Date(2024, time.August, 26, 11, 53, 25, 0, time.UTC)
	tracker.Add("I'm drinking one big mug of coffee right now.")

	checkTime := addTime.Add(time.Minute * 300)

	caffeineLevel, _ := tracker.GetLevel(checkTime)

	if caffeineLevel != 50 {
		t.Errorf("Expected half-life value of 50mg but got '%vmg'", caffeineLevel)
	}
}
