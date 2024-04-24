package main

import "testing"

func TestSum(t *testing.T) {
	numbers := [5]int{4,5,6,7,8}

	got := Sum(numbers)
	want := 30

	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}