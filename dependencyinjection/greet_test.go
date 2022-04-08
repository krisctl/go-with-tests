package dependencyinjection

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGreet(t *testing.T) {
	buf := bytes.Buffer{}
	Greet(&buf, "krish")
	want := "hello, krish!"
	got := buf.String()
	assert.Equal(t, want, got, "greeting not equal, want: %s, got: %s", got, want)
}
