package tracker

import "testing"

func TestLevelIsZero(t *testing.T) {
	caffeineLevel := GetLevel()

	if caffeineLevel != 0 {
		t.Errorf("Expected zero but got '%v'", caffeineLevel)
	}

}
