package hnet

import (
	"fmt"
	"testing"
)

func TestHasProtocol(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"http://example.com", true},
		{"ftp://files.server.com", true},
		{"https://secure.example.com", true},
		{"mailto:user@example.com", true},
		{"file:///path/to/file", true},
		{"custom-protocol+1.0://data", true},
		{"123protocol://invalid.com", false},      // Scheme starts with a digit
		{"://missing.scheme.com", false},          // Empty scheme
		{"http//missing.colon.com", false},        // Missing slash
		{"http:/missing.slash.com", false},        // Only one slash
		{"ht@tp://invalid.characters.com", false}, // Invalid character '@'
		{"http://", true},                         // Valid scheme with no authority
		{"", false},                               // Empty string
		{"no-protocol-here", false},               // No separator
		{"http:/example.com", false},              // Incorrect separator
		{"ht tp://space.in.scheme.com", false},    // Space in scheme
		{"HTTP://uppercase.scheme.com", true},     // Uppercase letters in scheme
		{"HtTp://MixedCase.scheme.com", true},     // Mixed case
		{"ht+tp://plus.in.scheme.com", true},      // Plus in scheme
		{"ht-tp://hyphen.in.scheme.com", true},    // Hyphen in scheme
		{"ht.tp://dot.in.scheme.com", true},       // Dot in scheme
		{"ht@tp://invalid@.com", false},           // Invalid '@' in scheme
		{"http://valid-scheme.com", true},         // Typo but valid characters
	}

	for _, tc := range testCases {
		result := HasProtocol(tc.input)
		status := "PASS"
		if result != tc.expected {
			status = "FAIL"
		}
		fmt.Printf("%s: Input: %-30s Expected: %-5v Got: %-5v\n", status, fmt.Sprintf("%q", tc.input), tc.expected, result)
	}
}
