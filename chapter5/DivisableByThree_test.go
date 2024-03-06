package main

import "testing"

func TestName(t *testing.T) {
	var answer = divisable()
	var result = 165
	if answer != result {
		t.Errorf("output %q not equal to % q ", answer, result)
	}

}
