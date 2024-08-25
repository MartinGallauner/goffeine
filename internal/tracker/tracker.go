package tracker

type Tracker struct {
	repository Repository
}

func New(repository Repository) *Tracker {
	return &Tracker{repository: repository}
}

type Repository interface {
	Fetch() int
	Add(caffeineInMg int)
}

func (tracker *Tracker) GetLevel() int {
	//todo calculate level for right now
	//todo cleanup entries older than 24h
	return tracker.repository.Fetch()
}

func (tracker *Tracker) Add(caffeineInMg int) {
	tracker.repository.Add(caffeineInMg)
}
