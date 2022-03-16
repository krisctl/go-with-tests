package structs

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPerimeter(t *testing.T) {
	var want, got float64
	want = 44.0
	r := Rectangle{width: 12.0, height: 10.0}
	got = r.Perimeter(12.0, 10.0)
	assert.Equal(t, want, got, fmt.Sprintf("Not equal, got %.2f, want %.2f", got, want))
}

func TestArea(t *testing.T) {
	var want, got float64
	want = 120.0
	r := Rectangle{width: 12.0, height: 10.0}
	got = r.Area(12.0, 10.0)
	assert.Equal(t, want, got, fmt.Sprintf("Not equal, got %.2f, want %.2f", got, want))
}
