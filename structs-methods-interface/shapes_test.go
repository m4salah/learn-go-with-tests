package structsmethodsinterface

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := rectangle.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f and we want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Circle", shape: Circle{10.0}, want: 314.1592653589793},
		{name: "Rectangle", shape: Rectangle{12, 6}, want: 72.0},
		{name: "Triangle", shape: Triangle{12, 6}, want: 36.0},
	}
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %.2f and we want %.2f", got, want)
		}

	}
	t.Run("Table driven tests", func(t *testing.T) {
		for _, tt := range areaTests {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("%#v got %g want %g", tt.name, got, tt.want)
			}
		}
	})
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		want := 100.0

		checkArea(t, rectangle, want)

	})

	t.Run("Circle", func(t *testing.T) {
		circle := Circle{10.0}
		want := 314.1592653589793

		checkArea(t, circle, want)
	})
}
