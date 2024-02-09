package main

import "testing"

type Rectangle struct {
	Height float64
	Width  float64
}

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{Height: 10.0, Width: 20.0}
		got := Area(rectangle)
		want := 200.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		got := Area(circle)
		want := 314.1592653589793

	})

}
