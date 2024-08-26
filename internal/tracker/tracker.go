package tracker

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

type Tracker struct {
	repository Repository
}

func New(repository Repository) *Tracker {
	return &Tracker{repository: repository}
}

type Repository interface {
	Fetch() ([][]string, error)
	Add(caffeineInMg int) error
}

type Entry struct {
	timestamp    time.Time
	caffeineInMg int
}

func (tracker *Tracker) GetLevel() (int, error) {
	//todo calculate level for right now
	//todo cleanup entries older than 24h
	data, _ := tracker.repository.Fetch()

	//parse data
	data = data[1:] //ignore the header row
	layout := "2006-01-02T15:04:05Z07:00"
	caffeineLevel := 0
	for _, row := range data {
		timestamp, err := time.Parse(layout, row[0])
		if err != nil {
			return 0, err
		}

		value := parseInt(row[1])

		if time.Now().Add(-1 * 24 * time.Hour).Before(timestamp) {
			caffeineLevel += value
		}
	}
	log.Printf("You have %q of caffeine in your system")
	return caffeineLevel, nil //todo calculate level
}

func parseInt(s string) int {
	// Implement error handling and conversion logic as needed
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Error converting to int:", err)
		return 0 // Or handle the error differently
	}
	return i
}

func (tracker *Tracker) Add(caffeineInMg int) {
	err := tracker.repository.Add(caffeineInMg)
	if err != nil {
		log.Println("Error adding caffeine")
	}
	log.Printf("Added %vmg of caffeine", caffeineInMg)
}
