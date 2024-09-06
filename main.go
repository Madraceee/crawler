package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	args := os.Args[1:]

	if len(args) != 3 {
		fmt.Println("Wrong arguments provided")
		fmt.Println("Format is - 'crawler <website> <concurrencyCount> <maxPages>'")
		os.Exit(1)
	}

	maxConcurrency, err := strconv.Atoi(args[1])
	maxPages, err := strconv.Atoi(args[2])

	if err != nil {
		fmt.Println("Provide Concurrency and maxPages are integers")
		os.Exit(1)
	}

	log.Printf("starting crawl of: %s\n", args[0])

	crawlerCfg := configureCrawler(args[0], maxConcurrency, maxPages)

	// Starting timer
	start := time.Now()

	// Starting Crawler
	crawlerCfg.wg.Add(1)
	crawlerCfg.crawlPage(args[0])
	crawlerCfg.wg.Wait()

	printReport(crawlerCfg.pages, crawlerCfg.baseURL)

	log.Println("-------------")
	log.Println("TIME TAKEN - ", time.Since(start))
	log.Println("Exiting")
	os.Exit(0)
}
