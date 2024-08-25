package main

import "testing"

func TestProcessArgsForStatus(t *testing.T) {

	command, _, _ := processArgs([]string{"status"})
	if command != "status" {
		t.Errorf("Got command %q but expected 'status'", command)
	}
}

func TestProcessArgsForAdd(t *testing.T) {
	command, num, _ := processArgs([]string{"add", "100"})
	if command != "add" {
		t.Errorf("Got command %q but expected 'add'", command)
	}
	if num != 100 {
		t.Errorf("Got value %v but expected '100'", num)
	}
}
