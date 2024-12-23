package hnet

import (
	"strings"
	"testing"
)

func TestURL_Custom(t *testing.T) {
	uppercaseTransformer := func(s URL) URL { return URL(strings.ToUpper(string(s))) }

	tests := []struct {
		input    URL
		options  []func(URL) URL
		expected URL
	}{
		{"example.com", []func(URL) URL{uppercaseTransformer}, "EXAMPLE.COM"},
		{"example.com", []func(URL) URL{AsProtocol(HTTP), uppercaseTransformer}, "HTTP://EXAMPLE.COM"},
		{"example.com", []func(URL) URL{AsProtocol(HTTPS)}, "https://example.com"},
		{"example.com", []func(URL) URL{AsProtocol(WS), uppercaseTransformer}, "WS://EXAMPLE.COM"},
		{"example.com", []func(URL) URL{AsProtocol(WSS)}, "wss://example.com"},
	}

	for _, tt := range tests {
		result := tt.input.Format(tt.options...)
		if result != tt.expected {
			t.Errorf("URL(%q, options...) = %q; want %q", tt.input, result, tt.expected)
		}
	}
}
