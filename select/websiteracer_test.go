package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestWebsiteRacer(t *testing.T) {
	slowServer := initiateDelayedServer(13 * time.Second)
	fastServer := initiateDelayedServer(11 * time.Second)
	t.Cleanup(func() {
		fastServer.Close()
		slowServer.Close()
	})

	fastUrl := fastServer.URL
	got, err := Racer(slowServer.URL, fastServer.URL)
	assert.Error(t, err, "Did not get an error here, was expecting a timeout error")
	assert.Equal(t, fastUrl, got, "got: %s, wanted: %s", got, fastUrl)
}

func initiateDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))

}
