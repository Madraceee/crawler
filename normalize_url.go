package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalize(inputURL string) (string, error) {
	parsedUrl, err := url.Parse(inputURL)
	if err != nil {
		return "", nil
	}

	outputURL := fmt.Sprintf("%s%s", parsedUrl.Hostname(), parsedUrl.Path)
	if len(parsedUrl.Path) == 0 {
		outputURL += "/"
	}

	if len(parsedUrl.Path) > 0 && parsedUrl.Path[len(parsedUrl.Path)-1] == '/' {
		outputURL = outputURL[:len(outputURL)-1]
	}

	outputURL = strings.ToLower(outputURL)

	if len(parsedUrl.RawQuery) > 0 {
		outputURL += "?" + parsedUrl.RawQuery
	}

	return outputURL, nil
}
