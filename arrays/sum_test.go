package arrays

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumArraySlice(t *testing.T) {
	t.Run("collection of 5 numbers in arrays", func(t *testing.T) {
		data := [5]int{1, 2, 3, 4, 5}
		got := SumArray(data)
		want := 15
		assert.Equal(t, want, got, fmt.Sprintf("Not equal, got %d, want %d", got, want))
	})
	t.Run("collection of numbers in slice", func(t *testing.T) {
		data := []int{1, 2, 3, 4}
		got := SumSlice(data)
		want := 10
		assert.Equal(t, want, got, fmt.Sprintf("Not equal, got %d, want %d", got, want))
	})
	t.Run("tesing SumAll functionality", func(t *testing.T) {
		// SumAll([]int{1,2}, []int{0,9}) would return []int{3, 9}
		want := []int{3, 9}
		// fmt.Println("Just for removing cache")
		got := SumAll([]int{1, 2}, []int{0, 9})
		assert.Equal(t, want, got, fmt.Sprintf("Not equal, got %v, want %v", got, want))
	})
	t.Run("testing sum of all tails", func(t *testing.T) {
		// SumAllTails([]int{1,2,3,4}, []int{1,9,2}, []int{1}) would return []int{9, 11, 0}
		want := []int{9, 11, 0, 0}
		got := SumAllTails([]int{1, 2, 3, 4}, []int{1, 9, 2}, []int{1}, []int{})
		assert.Equal(t, want, got, fmt.Sprintf("Not equal, got %v, want %v", got, want))
	})
}
