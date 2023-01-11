package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestWebsiteRacer(t *testing.T) {
	slowServer := initiateDelayedServer(1 * time.Millisecond)
	fastServer := initiateDelayedServer(0 * time.Millisecond)
	t.Cleanup(func() {
		fastServer.Close()
		slowServer.Close()
	})

	fastUrl := fastServer.URL
	got, err := Racer(slowServer.URL, fastUrl)
	assert.NoError(t, err, "unexpected error")
	assert.Equal(t, fastUrl, got, "got: %s, wanted: %s", got, fastUrl)
}

func TestWebsiteRacerTimeout(t *testing.T) {
	server := initiateDelayedServer(2 * time.Millisecond)
	t.Cleanup(func() {
		server.Close()
	})

	_, err := ConfigurableRacer(server.URL, server.URL, 1*time.Millisecond)
	assert.Error(t, err, "Did not get an error here, was expecting a timeout error")
}

func initiateDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))

}
