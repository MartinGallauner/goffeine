package tracker

import (
	"github.com/MartinGallauner/goffeine/internal/repository"
	"log"
	"time"
)

type Tracker struct {
	repository Repository
}

func New(repository Repository) *Tracker {
	return &Tracker{repository: repository}
}

type Repository interface {
	Fetch() ([]repository.Entry, error)
	Add(timestamp time.Time, caffeineInMg int) error
}

func (tracker *Tracker) GetLevel(now time.Time) (int, error) {
	//todo calculate level for right now
	//todo cleanup entries older than 24h

	entries, err := tracker.repository.Fetch()
	if err != nil {
		return 0, err
	}

	caffeineLevel := 0
	for _, entry := range entries {
		if now.Add(-1 * 24 * time.Hour).Before(entry.Timestamp) {
			caffeineLevel += entry.CaffeineInMg
		}
	}
	log.Printf("You have %vmg of caffeine in your system", caffeineLevel)
	return caffeineLevel, nil //todo calculate level
}

func (tracker *Tracker) Add(timestamp time.Time, caffeineInMg int) error {
	err := tracker.repository.Add(timestamp, caffeineInMg)
	if err != nil {
		log.Println("Error adding caffeine")
	}
	log.Printf("Added %vmg of caffeine", caffeineInMg)
	return nil
}
