package main

import (
	"fmt"
	"testing"
)

func TestClone(t *testing.T) {
	got := CloningString("a", 10)
	want := "aaaaaaaaaa"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func ExampleClone() {
	clonedString := CloningString("a", 3)
	fmt.Println(clonedString)
	// Output: aaa
}

func BenchmarkClone(b *testing.B) {
	for i := 0; i >= b.N; i++ {
		CloningString("a", b.N)
	}
}
