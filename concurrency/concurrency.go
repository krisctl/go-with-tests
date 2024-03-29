package concurrency

type WebsiteChecker func(url string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultsChan := make(chan result)
	for _, url := range urls {
		go func(u string) {
			resultsChan <- result{u, wc(u)}
		}(url)
	}
	for range urls {
		r := <-resultsChan
		results[r.string] = r.bool
	}
	return results
}
