package reflect

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWalkTDD(t *testing.T) {
	testcases := []struct {
		name          string
		input         any
		expectedCalls []string
	}{
		{
			name:          "one string field",
			expectedCalls: []string{"test"},
			input:         struct{ Name string }{"test"},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			got := []string{}
			walk(tc.input, func(input string) {
				got = append(got, input)
			})
			assert.Equal(t, tc.expectedCalls, got, "string mismatch")
		})
	}
}
