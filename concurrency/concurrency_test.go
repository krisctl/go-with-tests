package concurrency

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func mockWebsiteChecker(url string) bool {
	if url == "http://test2" {
		return false
	}
	return true
}

func slowMockWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func TestWebsiteChecker(t *testing.T) {
	testCases := []struct {
		mock WebsiteChecker
		urls []string
		want map[string]bool
	}{
		{
			mock: mockWebsiteChecker,
			urls: []string{"http://test1", "http://test2", "http://test3"},
			want: map[string]bool{"http://test1": true,
				"http://test2": false,
				"http://test3": true},
		},
	}
	for _, tC := range testCases {
		t.Run("test run", func(t *testing.T) {
			got := CheckWebsites(tC.mock, tC.urls)
			assert.Equal(t, tC.want, got, "results don't match")
		})
	}
}

func BenchmarkCheckWebsite(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "some url"
	}

	for i := 0; i < b.N; i++ {
		CheckWebsites(slowMockWebsiteChecker, urls)
	}
}
