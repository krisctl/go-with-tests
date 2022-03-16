package iteration

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepeat(t *testing.T) {
	repeatCount := 4
	got := Repeat("a", repeatCount)
	want := "aaaa"
	assert.Equal(t, want, got, fmt.Sprintf("got %q, want %q", got, want))
}
func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 1)
	}
}

func ExampleRepeat() {
	fmt.Println(Repeat("a", 2))
	//Output: aa
}
