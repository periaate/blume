package hnet

import (
	"strings"
	"testing"

	"github.com/periaate/blume/gen"
)

// Test Opt_HTTPS behavior
func TestOpt_HTTPS(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"https://example.com", "https://example.com"},
		{"http://example.com", "https://example.com"},
		{"example.com", "https://example.com"},
		{"", "https://"},
	}

	for _, tt := range tests {
		result := Opt_HTTPS(tt.input)
		if result != tt.expected {
			t.Errorf("Opt_HTTPS(%q) = %q; want %q", tt.input, result, tt.expected)
		}
	}
}

// Test Opt_HTTP behavior
func TestOpt_HTTP(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"http://example.com", "http://example.com"},
		{"https://example.com", "http://example.com"},
		{"example.com", "http://example.com"},
		{"", "http://"},
	}

	for _, tt := range tests {
		result := Opt_HTTP(tt.input)
		if result != tt.expected {
			t.Errorf("Opt_HTTP(%q) = %q; want %q", tt.input, result, tt.expected)
		}
	}
}

// Test URL with default option
func TestURL_Default(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"http://example.com", "http://example.com"},
		{"example.com", "http://example.com"},
		{"https://example.com", "http://example.com"},
		{"", "http://"},
	}

	for _, tt := range tests {
		result := URL(tt.input)
		if result != tt.expected {
			t.Errorf("URL(%q) = %q; want %q", tt.input, result, tt.expected)
		}
	}
}

// Test URL with custom Transformer
func TestURL_Custom(t *testing.T) {
	uppercaseTransformer := func(s string) string {
		return strings.ToUpper(s)
	}

	tests := []struct {
		input    string
		options  []gen.Transformer[string]
		expected string
	}{
		{"example.com", []gen.Transformer[string]{uppercaseTransformer}, "EXAMPLE.COM"},
		{"example.com", []gen.Transformer[string]{Opt_HTTP, uppercaseTransformer}, "HTTP://EXAMPLE.COM"},
		{"example.com", []gen.Transformer[string]{Opt_HTTPS}, "https://example.com"},
	}

	for _, tt := range tests {
		result := URL(tt.input, tt.options...)
		if result != tt.expected {
			t.Errorf("URL(%q, options...) = %q; want %q", tt.input, result, tt.expected)
		}
	}
}
