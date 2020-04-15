package main

import "testing"

func TestSum(t *testing.T) {
	value, _ := sum(1, 2, 3)

	if value != 6 {
		t.Error("Expected", 6, "Got", value)
	}
}
