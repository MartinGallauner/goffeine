package main

import "testing"

func TestProcessArgs(t *testing.T) {

	command := processArgs([]string{"status"})
	if command != "status" {
		t.Errorf("Got command %q but expected 'status'", command)
	}

}
