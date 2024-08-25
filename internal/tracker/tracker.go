package tracker

type tracker struct {
	repository Repository
}

func New(repository Repository) *tracker {
	return &tracker{repository: repository}
}

type Repository interface {
	fetch() int
	add(caffeineInMg int)
}

func (tracker *tracker) GetLevel() int {
	//todo calculate level for right now
	//todo cleanup entries older than 24h
	return tracker.repository.fetch()
}

func (tracker *tracker) Add(caffeineInMg int) {
	tracker.repository.add(caffeineInMg)
}
