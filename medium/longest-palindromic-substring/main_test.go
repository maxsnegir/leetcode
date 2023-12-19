package main

import (
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "babad",
			expected: "bab",
		},
		{
			input:    "cbbd",
			expected: "bb",
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if res := longestPalindrome(tt.input); res != tt.expected {
				t.Errorf("%s != %s", res, tt.expected)
			}
		})
	}
}
