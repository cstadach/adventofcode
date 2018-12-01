package main

import "testing"

func Test1(t *testing.T) {
	deltas := []int{+1, -1}
	result := loopDeltas(deltas)
	if result != 0 {
		t.Error(
			"Expected", 0,
			"But received",  result,
		)
	}
}

func Test2(t *testing.T) {
	deltas := []int{+3, +3, +4, -2, -4}
	result := loopDeltas(deltas)
	if result != 10 {
		t.Error(
			"Expected", 10,
			"But received",  result,
		)
	}
}

func Test3(t *testing.T) {
	deltas := []int{-6, +3, +8, +5, -6}
	result := loopDeltas(deltas)
	if result != 5 {
		t.Error(
			"Expected", 5,
			"But received",  result,
			)
	}
}

func Test4(t *testing.T) {
	deltas := []int{+7, +7, -2, -7, -4}
	result := loopDeltas(deltas)
	if result != 14 {
		t.Error(
			"Expected", 14,
			"But received",  result,
		)
	}
}
