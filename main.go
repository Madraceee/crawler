package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		log.Fatal("no website provided")
		os.Exit(1)
	}

	if len(args) > 1 {
		log.Fatal("too many arguments provided")
		os.Exit(1)
	}

	log.Printf("starting crawl of: %s\n", args[0])

	maxConcurrency := 4
	crawlerCfg := configureCrawler(args[0], maxConcurrency)

	// Starting timer
	start := time.Now()

	// Starting Crawler
	crawlerCfg.wg.Add(1)
	crawlerCfg.crawlPage(args[0])
	crawlerCfg.wg.Wait()

	log.Println("\n\n----REPORT----")
	for url, count := range crawlerCfg.pages {
		fmt.Println(url, count)
	}
	log.Println("-------------")
	log.Println("TIME TAKEN - ", time.Since(start))
	log.Println("Exiting")
	os.Exit(0)
}
