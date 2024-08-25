package tracker

type tracker struct {
	repository Repository
}

func New(repository Repository) *tracker {
	return &tracker{repository: repository}
}

type Repository interface {
	fetch() int
}

func (tracker *tracker) GetLevel() int {
	return tracker.repository.fetch()
}
