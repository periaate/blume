package hnet

import (
	"github.com/periaate/blume/fsio"
	"github.com/periaate/blume/gen"
	"github.com/periaate/blume/str"
)

func URL(inp string, options ...gen.Transformer[string]) string {
	if len(options) == 0 {
		options = append(options, Opt_HTTP)
	}
	return gen.Pipe[string](options...)(inp)
}

func Opt_HTTP(s string) string {
	if len(s) == 0 {
		return "http://"
	}
	if HasProtocol(s) {
		return str.ReplaceRegex(".*://", "http://")(s)
	}
	return fsio.Join("http://", s)
}

func Opt_HTTPS(s string) string {
	if len(s) == 0 {
		return "https://"
	}
	if HasProtocol(s) {
		return str.ReplaceRegex(".*://", "https://")(s)
	}
	return fsio.Join("https://", s)
}
