package main

import (
	"fmt"
	"testing"
)

// Comparing symbols
func TestCompare(t *testing.T) {
	t.Run("Equal test", func(t *testing.T) {
		got := Compare("good", "good")
		want := "Equal"
		assertCorrectMessage(got, want, t)
	})
	t.Run("Not Equal test", func(t *testing.T) {
		got := Compare("food", "foos")
		want := "Not Equal"
		assertCorrectMessage(got, want, t)
	})
}
func assertCorrectMessage(got, want string, t *testing.T) {
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func ExampleCompare() {
	result := Compare("good", "bad")
	fmt.Println(result)
	// Output: Not Equal
}

func BenchmarkCompare(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		Compare("food", "food")
	}
}
