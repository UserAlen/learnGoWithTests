package main

import "testing"

func TestReplace(t *testing.T) {
	t.Run("Contain with different cases", func(t *testing) {
		got := Replace("Hello, world!", "WORLD", "my friend")
		want := "Hello, my friend!"
		assertCorrectMessage(got, want, t)
	})
	t.Run("Contain with similar cases", func(t *testing) {
		got := Replace("Hello, world!", "world", "Boris")
		want := "Hello, Boris!"
		assertCorrectMessage(got, want, t)
	})
	t.Run("Not contain with different cases", func(t *testing) {
		got := Replace("Hello, world!", "APPLE", "pineapple")
		want := "Sentence \"Hello, world!\" doesn't have the word \"apple\""
		assertCorrectMessage(got, want, t)
	})
	t.Run("Not contain with similar cases")
}

func assertCorrectMessage(got, want string, t *testing.t) {
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
