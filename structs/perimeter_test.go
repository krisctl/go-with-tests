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

func TestArea(t *testing.T) {
	t.Run("Area of rectangle", func(t *testing.T) {
		var want, got float64
		want = 120.0
		r := Rectangle{width: 12.0, height: 10.0}
		got = r.Area()
		assert.Equal(t, want, got, fmt.Sprintf("Not equal, got %g, want %g", got, want))
	})
	t.Run("Area of circle", func(t *testing.T) {
		var want, got float64
		want = math.Pi * 100
		c := Circle{radius: 10.0}
		got = c.Area()
		assert.Equal(t, want, got, fmt.Sprintf("Not equal, got %g, want %g", got, want))
	})
}
