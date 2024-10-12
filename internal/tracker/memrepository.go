package tracker

import (
	"time"
)


type MemoryRepository struct {
	Entries map[int][]Entry //key == userID; not used yet.
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{Entries: make(map[int][]Entry)}
}


func (r *MemoryRepository) Fetch() ([]Entry, error) {
	return r.Entries[1], nil
}

func (r *MemoryRepository) Add(timestamp time.Time, caffeineInMg int) error {
	e := Entry{
		Timestamp: timestamp,
		CaffeineInMg: caffeineInMg,
	}


	if r.Entries[1] == nil {
		r.Entries[1] = []Entry{e}
	} else {
		r.Entries[1] = append(r.Entries[1], e)
	}

	return nil
}