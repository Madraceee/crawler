package main

import (
	"log"
	"os"
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
	os.Exit(0)
}
