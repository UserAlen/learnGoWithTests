package main

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d givern, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	/* Starting from Go 1.21 it now has slices package with equal.
	But it doesn't work with 2D slices if it's 1D slice

	reflect.DeepEqual(thing1, thing2) doesn't have the "type safe" thing.
	It can compare a slice with a string..... which makes no sense
	*/
	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t *testing.T, got, want []int) {
		t.Helper()

		// choosing !slices.Equal or !reflect.DeepEqual doesn't matter - the func covers type-safety
		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	//checking the checksums type-safety
	/* 	t.Run("checking the checkSums type-safety", func(t *testing.T) {
		got := 1234
		want := "hello"
		checkSums(t, got, want)
	}) */

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{3, 5, 6})
		want := []int{2, 5 + 6}

		checkSums(t, got, want)
	})
	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{4, 9})
		want := []int{0, 9}

		checkSums(t, got, want)
	})

}
