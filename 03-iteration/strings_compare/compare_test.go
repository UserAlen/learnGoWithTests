package main

import (
	"fmt"
	"testing"
)

func TestCompare(t *testing.T) {
	t.Run("If it's equal", func(t *testing.T) {
		got := Compare("good", "good")
		want := "equal"
		assertCorrectMessage(got, want, t)
	})
	t.Run("If it's not equal", func(t *testing.T) {
		got := Compare("good", "bad")
		want := "not equal"
		assertCorrectMessage(got, want, t)
	})
}

func assertCorrectMessage(got string, want string, t *testing.T) {
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
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
