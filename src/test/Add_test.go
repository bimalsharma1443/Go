package main

import "testing"

func TestAdd(t *testing.T) {
	x := add(3, 2)
	if x != 5 {
		t.Error("expected value is 5 got : ", x)
	}
}
