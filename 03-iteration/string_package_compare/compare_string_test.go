package iteration

import (
	"fmt"
	"testing"
)

// Comparing symbols
func TestCompare(t *testing.T) {
	got := Compare("good", "good")
	want := "Equal"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	got := Compare("food", "foos")
	want := "Not Equal"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func ExampleCompare() {
	result := Compare("good", "bad")
	fmt.Println(result)
	// Output: "Not Equal"
}

func BenchmarkCompare(b *testing.B) {

}
