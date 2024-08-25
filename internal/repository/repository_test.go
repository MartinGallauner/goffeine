package repository

import "testing"

func TestReadCSV(t *testing.T) {
	repository := New("testdata/data.csv")
	value := repository.fetch()

	if value != 0 {
		t.Errorf("Got %q but expected '0'", value)
	}

}
