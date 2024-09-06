package main

import (
	"fmt"
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

	isMaxPagesReached := cfg.checkMaxPagesReached()
	if isMaxPagesReached {
		return
	}

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
	log.Println("crawled website successfully - ", normalizedURL)
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

func printReport(pages map[string]int, baseURL string) {
	fmt.Println("=================")
	fmt.Println("  REPORT for", baseURL)
	fmt.Println("=================")

	// Sort and print
	for link, count := range pages {
		fmt.Printf("Found %d internal links to %s\n", count, link)
	}

}
