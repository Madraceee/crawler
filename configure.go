package main

import "sync"

type config struct {
	pages              map[string]int
	baseURL            string
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}

func (cfg *config) addPageVisit(url string) bool {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	_, present := cfg.pages[url]
	cfg.pages[url] += 1
	return present
}

func configureCrawler(baseURL string, maxConcurrency int) config {

	pages := make(map[string]int)
	return config{
		pages:              pages,
		baseURL:            baseURL,
		mu:                 &sync.Mutex{},
		wg:                 &sync.WaitGroup{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
	}
}
