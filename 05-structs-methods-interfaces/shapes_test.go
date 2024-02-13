package main

import "testing"

func TestPerimeter(t *testing.T) {

	rectangle := Rectangle{10.0, 10.0}
	got := rectangle.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	//table driven test

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		//for specific test "go test -run TestArea/Rectangle"
		{name: "Rectangle", shape: Rectangle{Width: 10.0, Height: 20.0}, hasArea: 200.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				// The %#v format string will print out our struct with the values in its field
				// as in Circle{Radius: 10}
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})
	}

	/* 	checkArea := func(t testing.TB, shape Shape, want float64) {
	   		t.Helper()
	   		got := shape.Area()
	   		if got != want {
	   			t.Errorf("got %g want %g", got, want)
	   		}
	   	}

	   	t.Run("rectangles", func(t *testing.T) {
	   		rectangle := Rectangle{Height: 10.0, Width: 20.0}
	   		want := 200.0

	   		checkArea(t, rectangle, want)
	   	})
	   	t.Run("circles", func(t *testing.T) {
	   		circle := Circle{10}
	   		want := 314.1592653589793

	   		checkArea(t, circle, want)
	   	}) */

}
