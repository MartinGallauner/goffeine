package tracker

import "testing"

type TestRepository struct {
}

func (r *TestRepository) fetch() int {
	return 0
}

func TestLevelIsZero(t *testing.T) {
	tracker := New(&TestRepository{})
	caffeineLevel := tracker.GetLevel()

	if caffeineLevel != 0 {
		t.Errorf("Expected zero but got '%v'", caffeineLevel)
	}

}
