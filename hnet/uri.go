package hnet

import (
	"strings"
	"unicode"
)

// Thanks, ChatGPT

// HasProtocol checks whether the given string s starts with a valid protocol scheme followed by "://"
// According to RFC 3986, the scheme must start with a letter, followed by letters, digits, "+", "-", or "."
func HasProtocol[S ~string](str S) bool {
	s := string(str)
	// Define the separator
	sep := "://"

	// Find the index of the separator
	sepIndex := strings.Index(s, sep)
	if sepIndex == -1 {
		// Separator not found
		return false
	}

	// Extract the scheme part
	scheme := s[:sepIndex]
	if len(scheme) == 0 {
		// Empty scheme
		return false
	}

	// Validate the first character: must be a letter
	firstRune, _ := utf8DecodeRuneInString(scheme, 0)
	if !unicode.IsLetter(firstRune) {
		return false
	}

	// Validate the remaining characters
	for _, r := range scheme[1:] {
		if !(unicode.IsLetter(r) || unicode.IsDigit(r) || r == '+' || r == '-' || r == '.') {
			return false
		}
	}

	// All checks passed
	return true
}

// utf8DecodeRuneInString is a helper function to decode the first rune in a string.
// It returns the rune and its size in bytes.
// If the string is empty, it returns utf8.RuneError and 0.
func utf8DecodeRuneInString(s string, index int) (r rune, size int) {
	if index >= len(s) {
		return unicode.ReplacementChar, 0
	}
	return utf8DecodeRune([]byte(s[index:]))
}

// utf8DecodeRune decodes the first rune in the given byte slice.
func utf8DecodeRune(b []byte) (r rune, size int) {
	if len(b) == 0 {
		return unicode.ReplacementChar, 0
	}
	c := b[0]
	// Simple ASCII check
	if c < 0x80 {
		return rune(c), 1
	}
	// For simplicity, assume invalid for non-ASCII
	return unicode.ReplacementChar, 0
}
