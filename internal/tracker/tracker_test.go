package tracker

import "testing"

func TestLevelIsZero(t *testing.T) {
	caffeineLevel := 0

	if caffeineLevel != 0 {
		t.Errorf("Expected zero but got ")
	}

}
