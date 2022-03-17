package structs

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPerimeter(t *testing.T) {
	var want, got float64
	want = 44.0
	r := Rectangle{width: 12.0, height: 10.0}
	got = r.Perimeter()
	assert.Equal(t, want, got, fmt.Sprintf("Not equal, got %.2f, want %.2f", got, want))
}

func TestAreaWithInterface(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		assert.Equal(t, want, got, fmt.Sprintf("Not equal, got %g, want %g", got, want))
	}
	t.Run("Area of rectangle", func(t *testing.T) {
		shape := Rectangle{width: 2, height: 2}
		checkArea(t, shape, 4.0)
	})
	t.Run("Area of circle", func(t *testing.T) {
		shape := Circle{radius: 2}
		checkArea(t, shape, math.Pi*shape.radius*shape.radius)
	})
}

func TestAreaWithTables(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{"Area of Rectangle", Rectangle{width: 10, height: 12}, 120},
		{"Area of Circle", Circle{radius: 10}, 314.1592653589793},
		{"Area of Triangle", Triangle{base: 10, height: 12}, 60},
	}

	for _, test := range areaTests {
		t.Run(test.name, func(t *testing.T) {
			got := test.shape.Area()
			assert.Equal(t, test.want, got, fmt.Sprintf("Shape: %#v, Area is not equal, got %g, want %g", test.shape, got, test.want))
		})
	}
}
