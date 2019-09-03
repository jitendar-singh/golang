package main

import "testing"

func TestMySum(t *testing.T) {
	type test struct {
		data   []int
		answer int
	}
	tests := []test{
		test{[]int{21, 21}, 42},
		test{[]int{3, 4, 5}, 12},
		test{[]int{1, 1}, 2},
		test{[]int{-1, 0, 1}, 0},
	}

	for _, y := range tests {
		x := mySum(y.data...)
		if x != y.answer {
			t.Error("Expected", y.answer, "got", x)
		}
	}
}
func TestMySub(t *testing.T) {
	type testSub struct {
		data   []int
		answer int
	}
	testsub := []testSub{
		testSub{[]int{21, 21}, 0},
		testSub{[]int{3, -5}, -8},
		testSub{[]int{2, 21}, 19},
		testSub{[]int{21, 121}, 100},
	}

	for _, y := range testsub {
		x := mySub(y.data...)
		if x != y.answer {
			t.Error("Expected", y.answer, "got", x)
		}
	}
}
