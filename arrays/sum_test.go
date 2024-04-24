package main

import "testing"
import "reflect"

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{4,5,6,7,8}

		got := Sum(numbers)
		want := 30

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1,2}, []int{0,9})
	want := []int{3,9}

	if !reflect.DeepEqual(got,want) {
		t.Errorf("got %v want %v", got, want)
	}
}