package racer

import (
	"fmt"
	"net/http"
	"time"
)

const TIMEOUT time.Duration = 10 * time.Second

func Racer(u1, u2 string) (string, error) {
	select {
	case <-ping(u1):
		return u1, nil
	case <-ping(u2):
		return u2, nil
	case <-time.After(10 * time.Second):
		return "", fmt.Errorf("timed out waiting for: %v, %v", u1, u2)
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
