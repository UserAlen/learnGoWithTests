package main

import "testing"

func TestRepeat(t *testing.T) {
	received := RepeatContinue()
	expected := "1\n3\n5"

	if received != expected {
		t.Errorf("expected %q but got %q", expected, received)
	}
}
