package main

import (
	"log"
	"strings"
)

func (cfg *config) crawlPage(rawCurrentUrl string) {
	// We wish to stay in the same domain for now
	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()

	if !strings.Contains(rawCurrentUrl, cfg.baseURL) {
		log.Println("Going outside the required base URL")
		return
	}

	normalizedURL, err := normalize(rawCurrentUrl)
	if err != nil {
		log.Println(err)
		return
	}

	present := cfg.addPageVisit(normalizedURL)
	if present {
		return
	}

	log.Println("Crawling website - ", normalizedURL)
	data, err := getHTML(rawCurrentUrl)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Crawled website successfully - ", normalizedURL)
	newLinks, err := getURLsFromHTML(data, cfg.baseURL)
	if err != nil {
		log.Println(err)
		return
	}

	for _, link := range newLinks {
		normalizedLink, _ := normalize(link)
		if normalizedURL == normalizedLink {
			continue
		}
		cfg.wg.Add(1)
		go cfg.crawlPage(link)
	}
}
