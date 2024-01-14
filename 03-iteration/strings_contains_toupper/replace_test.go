package main

import (
	"fmt"
	"testing"
)

func TestReplace(t *testing.T) {
	t.Run("Replace the initial word with different case with other word", func(t *testing.T) {
		got := Replace("Hello, world!", "WORLD", "my friend")
		want := "Hello, my friend!"
		assertCorrectMessage(got, want, t)
	})
	t.Run("Replace the initial word with similar case with other word", func(t *testing.T) {
		got := Replace("Hello, world!", "world", "Boris")
		want := "Hello, Boris!"
		assertCorrectMessage(got, want, t)
	})
	t.Run("Can't replace the initial word with different case with other word", func(t *testing.T) {
		got := Replace("Hello, world!", "APPLE", "pineapple")
		want := "Sentence \"Hello, world!\" doesn't have the word \"apple\". Can't replace with \"pineapple\""
		assertCorrectMessage(got, want, t)
	})
	t.Run("Can't replace the initial word with similar case with other word", func(t *testing.T) {
		got := Replace("Hello, world!", "sausage", "cheese")
		want := "Sentence \"Hello, world!\" doesn't have the word \"sausage\". Can't replace with \"cheese\""
		assertCorrectMessage(got, want, t)
	})
}

func assertCorrectMessage(got, want string, t *testing.T) {
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func ExampleReplace() {
	result := Replace("Hello, world", "HElLo", "Good afternoon")
	fmt.Println(result)
	// Output: Good afternoon, world
}

func BenchmarkReplace(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		Replace("Hello world", "hello", "Morning")
	}
}
