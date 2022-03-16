package integers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	sum := Add(2, 2)
	want := 4
	assert.Equal(t, want, sum, fmt.Sprintf("Not equal, got: %d, want: %d", sum, want))
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	//Output: 6

}
