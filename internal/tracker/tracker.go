package tracker

import (
	"github.com/MartinGallauner/goffeine/internal/ask"
	"log"
	"math"
	"time"
)

const halfLife = 5 * time.Hour //half life of caffeine TODO move to config

type LlmClient interface {
	Ask(input string) (ask.CaffeineIntake, error)
}

type Tracker struct {
	repository Repository
	client     LlmClient
}

func New(repository Repository, client LlmClient) *Tracker {
	return &Tracker{repository: repository, client: client}
}

type Repository interface {
	Fetch() ([]Entry, error)
	Add(timestamp time.Time, caffeineInMg int) error
}

type Entry struct {
	Timestamp    time.Time
	CaffeineInMg int
}

func (tracker *Tracker) GetLevel(now time.Time) (int, error) {
	// TODO cleanup entries older than 24h?

	entries, err := tracker.repository.Fetch()
	if err != nil {
		return 0, err
	}

	caffeineLevel := 0
	for _, entry := range entries {
		if now.Add(-1 * 24 * time.Hour).Before(entry.Timestamp) { //filter out entries older than 24h

			remainingCaffeine := calculateRemainingCaffeine(entry.CaffeineInMg, now.Sub(entry.Timestamp), halfLife)
			caffeineLevel += remainingCaffeine
		}
	}
	log.Printf("You have %vmg of caffeine in your system", caffeineLevel)
	return caffeineLevel, nil
}

// CalculateRemainingCaffeine calculates the remaining quantity of a substance
// after a given time period based on its half-life.
func calculateRemainingCaffeine(initialAmount int, elapsed time.Duration, halfLife time.Duration) int {
	// Convert elapsed time and half-life to hours
	elapsedMinutes := elapsed.Minutes()
	halfLifeMinutes := halfLife.Minutes()

	// Apply the exponential decay formula
	remainingAmount := float64(initialAmount) * math.Pow(0.5, elapsedMinutes/halfLifeMinutes)
	return int(remainingAmount) // TODO conversion is a bit risky
}

func (tracker *Tracker) Add(userInput string) error {
	caffeineIntake, err := tracker.client.Ask(userInput)
	if err != nil {
		return err
	}

	err = tracker.repository.Add(caffeineIntake.Timestamp, caffeineIntake.CaffeineInMg)
	if err != nil {
		log.Println("Error adding caffeine")
	}
	log.Printf("Added %vmg of caffeine", caffeineIntake.CaffeineInMg)
	return nil
}
