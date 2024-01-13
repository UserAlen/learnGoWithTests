package main

import (
	"fmt"
	"testing"
)

func TestCompare(t *testing.T) {
	t.Run("If it's equal", func(t *testing.T) {
		got := Compare("good", "good")
		want := "equal"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("If it's not equal", func(t *testing.T) {
		got := Compare("good", "bad")
		want := "not equal"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

func ExampleCompare() {
	result := Compare("good", "good")
	fmt.Println(result)
	// Output: equal
}

func BenchmarkCompare(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		Compare("good", "good")
	}
}
