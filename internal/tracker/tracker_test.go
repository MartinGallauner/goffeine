package tracker

import "testing"

type TestRepository struct {
	counter int
}

func (r *TestRepository) Fetch() int {
	return r.counter
}

func (r *TestRepository) Add(caffeineInMg int) {
	r.counter += caffeineInMg
}

func TestLevelIsZero(t *testing.T) {
	tracker := New(&TestRepository{})
	caffeineLevel := tracker.GetLevel()

	if caffeineLevel != 0 {
		t.Errorf("Expected zero but got '%v'", caffeineLevel)
	}
}

func TestAddCaffeine(t *testing.T) {
	tracker := New(&TestRepository{})
	tracker.Add(100)
	caffeineLevel := tracker.GetLevel()

	if caffeineLevel != 100 {
		t.Errorf("Expected 100 but got '%v'", caffeineLevel)
	}
}
