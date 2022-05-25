package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

type SpySleeper struct {
	count int
}

func (ss *SpySleeper) Sleep() {
	ss.count++
}

func TestMockCountdown(t *testing.T) {
	ss := &SpySleeper{}
	buf := &bytes.Buffer{}
	Countdown(buf, ss)

	op := buf.String()
	assert.Equal(t, 4, ss.count, "sleep is called diff number of times than expected")
	assert.Equal(t, "3\n2\n1\nGo!", op, "got unexpected outout")
}
