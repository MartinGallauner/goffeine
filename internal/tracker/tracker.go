package tracker

import "log"

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

func (tracker *Tracker) GetLevel() {
	//todo calculate level for right now
	//todo cleanup entries older than 24h
	data, _ := tracker.repository.Fetch()
	log.Println(data)

}

func (tracker *Tracker) Add(caffeineInMg int) {
	err := tracker.repository.Add(caffeineInMg)
	if err != nil {
		log.Println("Error adding caffeine")
	}
	log.Printf("Added %vmg of caffeine", caffeineInMg)
}
