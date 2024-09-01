package main

import "testing"

func TestNormalizeTest(t *testing.T) {
	tests := []struct {
		name     string
		inputURL string
		expected string
	}{
		{
			name:     "Remove Scheme",
			inputURL: "https://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "Remove Scheme",
			inputURL: "https://blog.boot.dev/path?key=value",
			expected: "blog.boot.dev/path?key=value",
		},
		{
			name:     "Remove Scheme",
			inputURL: "https://www.blog.boot.dev/path",
			expected: "www.blog.boot.dev/path",
		},
		{
			name:     "Remove Scheme",
			inputURL: "https://blog.boot.dev",
			expected: "blog.boot.dev/",
		},
		{
			name:     "Remove Scheme",
			inputURL: "https://blog.boot.dev:433",
			expected: "blog.boot.dev/",
		},
		{
			name:     "Remove Scheme",
			inputURL: "http://blog.boot.dev/path/",
			expected: "blog.boot.dev/path",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := normalize(tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
			}

			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
