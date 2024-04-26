package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("Test Perimeter", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := Perimeter(rectangle)
		want := 40.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
}

func TestArea(t *testing.T) {
	
	areaTests := []struct {
		name string
		shape Shape
		want float64
	}{
		{name: "Rectangle 1", shape: Rectangle{12, 6}, want: 72.00},
		{name: "Circle 1", shape: Circle{10}, want: 314.1592653589793},
		{name: "Rectangle 2", shape: Rectangle{10, 10}, want: 100.0},
		{name: "Triangle 1", shape: Triangle{12, 6}, want: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.want)
			}
		})
	}
}