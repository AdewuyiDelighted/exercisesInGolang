package main

import "testing"

func TestNameTwo(t *testing.T) {
	autoPolicy := AutoPolicy{
		state: "CT",
	}
	autoPolicy.setState("MA")
	state := autoPolicy.getState()
	if state != "Massachuetts" {
		t.Errorf("Expected %q got %q", "Massachuetts", state)
	}

}
