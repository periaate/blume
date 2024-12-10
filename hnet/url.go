package hnet

import (
	"net/url"

	. "github.com/periaate/blume/core"
	"github.com/periaate/blume/gen"
)

//-blume:derive String
type URL string

func (u URL) ToURL() (*url.URL, error) { return url.Parse(string(u.Format())) }

func (u URL) Format(options ...Monadic[URL, URL]) URL {
	if len(options) == 0 { options = append(options, AsProtocol(HTTP)) }
	return Pipe[URL](options...)(u)
}

type Protocol string

const (
	HTTP  Protocol = "http"
	HTTPS Protocol = "https"
	WS    Protocol = "ws"
	WSS   Protocol = "wss"
)

func (u URL) AsProtocol(protocol Protocol) URL {
	if len(u) == 0 { return URL(protocol + "://") }
	if HasProtocol(u) { return URL(gen.ReplaceRegex[URL](".*://", string(protocol+"://"))(u)) }
	return URL(protocol) + "://" + u
}

func AsProtocol(protocol Protocol) Monadic[URL, URL] {
	return func(u URL) URL { return u.AsProtocol(protocol) }
}
