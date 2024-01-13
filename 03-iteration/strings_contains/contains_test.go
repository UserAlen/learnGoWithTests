package main

import (
	"fmt"
	"testing"
)

func TestContains(t *testing.T) {
	t.Run("Contains the thing", func(t *testing.T) {
		got := Contains("Hello, world!", "world")
		want := true
		assertCorrectMessage(got, want, t)
	})
	t.Run("Doesn't contain the thing", func(t *testing.T) {
		got := Contains("Hello, world!", "Boris")
		want := false
		assertCorrectMessage(got, want, t)
	})
}

func assertCorrectMessage(got, want bool, t *testing.T) {
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func ExampleCompare() {
	result := Contains("pineapple", "apple")
	fmt.Println(result)
	// Output: true
}

func BenchmarkCompare(b *testing.b) {
	for i := 0; i <= b.N; i++ {
		Contains("my name is", "name")
	}
}
