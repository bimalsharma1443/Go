package main2

import "testing"

type test struct {
	data   []int
	answer int
}

func TestSum(t *testing.T) {
	tests := []test{
		test{[]int{1, 2, 3}, 42},
		test{[]int{10, 22, 34}, 12},
		test{[]int{1, 2, 3}, 6},
		test{[]int{0, -1, 0, 1}, 0},
	}

	for _, v := range tests {
		result, _ := sum(v.data...)
		if result != v.answer {
			t.Error("Expected", v.answer, "Got", result)
		}
	}
}
