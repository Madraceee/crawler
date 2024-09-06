package main

import "sync"

type config struct {
	pages              map[string]int
	baseURL            string
	maxPages           int
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

func (cfg *config) checkMaxPagesReached() bool {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	return cfg.maxPages == len(cfg.pages)
}

func configureCrawler(baseURL string, maxConcurrency, maxPages int) config {

	pages := make(map[string]int)
	return config{
		pages:              pages,
		baseURL:            baseURL,
		maxPages:           maxPages,
		mu:                 &sync.Mutex{},
		wg:                 &sync.WaitGroup{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
	}
}
