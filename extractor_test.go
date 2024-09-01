package main

import (
	"reflect"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
	}{
		{
			name:     "Absolute and relative URLs",
			inputURL: "https://www.test.com",
			inputBody: `
			<html>
				<body>
					<a href="/path">
						<span>Boot.dev</span>
					</a>
					<a href="https://www.google.com">
						<span>Boot.dev</span>
					</a>
				</body>
			</html>
		`,
			expected: []string{"https://www.test.com/path", "https://www.google.com"},
		},
		{
			name:     "Absolute and relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
			<html>
				<body>
					<a href="/path/one">
						<span>Boot.dev</span>
					</a>
					<a href="https://other.com/path/one">
						<span>Boot.dev</span>
					</a>
				</body>
			</html>
			`,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getURLsFromHTML(tc.inputBody, tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected errorL %v", i, tc.name, err)
			}

			if reflect.DeepEqual(actual, tc.expected) == false {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}

func TestGetAbsoluteURLs(t *testing.T) {
	tests := []struct {
		name  string
		input struct {
			inputURLs  []string
			rawBaseURL string
		}
		outputURLs []string
	}{
		{
			name: "Get Absolute URLs",
			input: struct {
				inputURLs  []string
				rawBaseURL string
			}{
				inputURLs:  []string{"/path/test", "https://www.google.com", "https://blog.boot.dev", "/temp?key=value"},
				rawBaseURL: "https://www.test.com",
			},
			outputURLs: []string{"https://www.test.com/path/test", "https://www.google.com", "https://blog.boot.dev", "https://www.test.com/temp?key=value"},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			absUrls, err := getAbsoluteURLs(tc.input.inputURLs, tc.input.rawBaseURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected errorL %v", i, tc.name, err)
			}

			if reflect.DeepEqual(absUrls, tc.outputURLs) == false {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.outputURLs, absUrls)
			}
		})
	}
}
