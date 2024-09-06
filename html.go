package main

import (
	"errors"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	// Check for valid URL
	if len(rawURL) == 0 {
		return "", errors.New("Enter Valid URL")
	}

	resp, err := http.Get(rawURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return "", errors.New("Error while fetching data")
	}

	if !strings.Contains(resp.Header.Get("Content-Type"), "text/html") {
		return "", errors.New("received content is not html, it is " + resp.Header.Get("Content-Type"))
	}

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(respBytes[:]), nil
}
