package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	finalWord      = "Go!"
	startingNumber = 3
)

type Sleeper interface {
	Sleep()
}

type RealSleeper struct {
}

func (rs *RealSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func main() {
	Countdown(os.Stdout, &RealSleeper{})
}

func Countdown(w io.Writer, s Sleeper) {
	for i := startingNumber; i > 0; i-- {
		s.Sleep()
		fmt.Fprintln(w, i)
	}
	s.Sleep()
	fmt.Fprint(w, finalWord)
}
