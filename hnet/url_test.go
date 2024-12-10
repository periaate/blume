package hnet

import (
	"strings"
	"testing"

	. "github.com/periaate/blume/core"
)

func TestURL_Custom(t *testing.T) {
	uppercaseTransformer := func(s URL) URL { return URL(strings.ToUpper(string(s))) }

	tests := []struct {
		input    URL
		options  []Monadic[URL, URL]
		expected URL
	}{
		{"example.com", []Monadic[URL, URL]{uppercaseTransformer}, "EXAMPLE.COM"},
		{"example.com", []Monadic[URL, URL]{AsProtocol(HTTP), uppercaseTransformer}, "HTTP://EXAMPLE.COM"},
		{"example.com", []Monadic[URL, URL]{AsProtocol(HTTPS)}, "https://example.com"},
		{"example.com", []Monadic[URL, URL]{AsProtocol(WS), uppercaseTransformer}, "WS://EXAMPLE.COM"},
		{"example.com", []Monadic[URL, URL]{AsProtocol(WSS)}, "wss://example.com"},
	}

	for _, tt := range tests {
		result := tt.input.Format(tt.options...)
		if result != tt.expected {
			t.Errorf("URL(%q, options...) = %q; want %q", tt.input, result, tt.expected)
		}
	}
}
