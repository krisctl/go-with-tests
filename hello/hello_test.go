package hello

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T) {
	got := hello()
	want := "hello, world!"
	assert.Equal(t, want, got, fmt.Sprintf("got %q, want %q", got, want))
}

func TestGreeting(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		language string
		expected string
	}{
		{
			name:     "user provided",
			input:    "user",
			language: "english",
			expected: "hello, user!",
		},
		{
			name:     "empty user",
			input:    "",
			expected: "hello, world!",
		},
		{
			name:     "spanish language",
			input:    "user",
			language: "Spanish",
			expected: "hola, user!",
		},
		{
			name:     "french language",
			input:    "user",
			language: "French",
			expected: "bonjour, user!",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			got := helloGreeting(tC.input, tC.language)
			want := tC.expected
			assert.Equal(t, want, got, fmt.Sprintf("got %q, want %q", got, want))
		})
	}
}
