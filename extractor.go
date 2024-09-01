package main

import (
	"fmt"
	"log"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func getAbsoluteURLs(unNormalizedArr []string, rawBaseURL string) ([]string, error) {
	for i, link := range unNormalizedArr {
		urlLink, err := url.Parse(link)
		if err != nil {
			log.Printf("unable to parse url - %s\n", urlLink)
			continue
		}

		if !urlLink.IsAbs() {
			unNormalizedArr[i] = rawBaseURL + unNormalizedArr[i]
		}
	}
	return unNormalizedArr, nil
}

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	doc, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return []string{}, fmt.Errorf("error while parsing html - %s", err)
	}

	unNormalizedURLs := []string{}

	var parser func(*html.Node)
	parser = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					unNormalizedURLs = append(unNormalizedURLs, a.Val)
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			parser(c)
		}
	}

	parser(doc)

	unNormalizedURLs, err = getAbsoluteURLs(unNormalizedURLs, rawBaseURL)
	if err != nil {
		return []string{}, fmt.Errorf("error while getting absolute urls - %s", err)
	}

	return unNormalizedURLs, nil
}
