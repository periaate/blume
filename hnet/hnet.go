package hnet

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	. "github.com/periaate/blume"
	"github.com/periaate/blume/yap"
)

type URL string

func (u URL) ToURL(options ...func(URL) URL) (*url.URL, error) {
	return url.Parse(string(u.Format(options...)))
}

func (u URL) Format(options ...func(URL) URL) URL {
	if len(options) == 0 {
		options = append(options, AsProtocol(HTTP))
	}
	return Pipe(options...)(u)
}

type Protocol string

const (
	HTTP  Protocol = "http"
	HTTPS Protocol = "https"
	WS    Protocol = "ws"
	WSS   Protocol = "wss"
)

// Use funny go type system haha why should we allow ~string as input to methods haha i love "type safety"
func (p Protocol) Use(h any) string {
	switch v := h.(type) {
	case string:
		return string(URL(v).AsProtocol(p))
	case URL:
		return string(v.AsProtocol(p))
	case String:
		return string(URL(v).AsProtocol(p))
	default:
		return string(URL(fmt.Sprint(h)).AsProtocol(p))
	}
}

func (u URL) AsProtocol(protocol Protocol) URL {
	if len(u) == 0 {
		return URL(protocol + "://")
	}
	if HasPrefix(HTTP, HTTPS, WS, WSS)(Protocol(u)) {
		return URL(ReplaceRegex[URL](".*://", string(protocol+"://"))(u))
	}
	return URL(protocol) + "://" + u
}

func AsProtocol(protocol Protocol) func(URL) URL {
	return func(u URL) URL { return u.AsProtocol(protocol) }
}

// Short hands for http headers.
const (
	Allow_Credentials = Access_Control_Allow_Credentials
	Allow_Headers     = Access_Control_Allow_Headers
	Allow_Methods     = Access_Control_Allow_Methods
	Allow_Origin      = Access_Control_Allow_Origin
	Expose_Headers    = Access_Control_Expose_Headers
	Max_Age           = Access_Control_Max_Age
)

type Header string

func (h Header) String() string                        { return string(h) }
func (h Header) Set(w http.ResponseWriter, val string) { w.Header().Set(string(h), val) }
func (h Header) Tuple(val any) [2]string {
	switch v := val.(type) {
	case string:
		return [2]string{string(h), v}
	case fmt.Stringer:
		return [2]string{string(h), v.String()}
	default:
		return [2]string{string(h), fmt.Sprint(val)}
	}
}

func NewError(code Status, message string) NetError { return NetError{Code: code, Message: message} }

type NetError struct {
	Code    Status
	Message string
}

func (e NetError) Error() string { return fmt.Sprintf("%d: %s", e.Code, e.Message) }

// Status represents an HTTP status code.
type Status int

func (s Status) ToError(message string) NetError { return NewError(s, message) }
func (s Status) ToErrorf(format string, args ...any) NetError {
	return NewError(s, fmt.Sprintf(format, args...))
}
func (s Status) Error() string { return s.Explanation() }
func (s Status) Respond(w http.ResponseWriter, r io.Reader) {
	io.Copy(w, r)
	w.WriteHeader(int(s))
}
func (s Status) Respondf(w http.ResponseWriter, format string, args ...any) {
	res := fmt.Sprintf(format, args...)
	if len(res) == 0 {
		res = s.Explanation()
	}
	if isErr(s) {
		http.Error(w, res, int(s))
		return
	}
	w.Write([]byte(res))
	w.WriteHeader(int(s))
}

func (s Status) AsError(w http.ResponseWriter) error {
	if isErr(s) {
		http.Error(w, s.Explanation(), int(s))
		return nil
	}
	return fmt.Errorf("status code [%d] is not an error", s)
}

func (s Status) AsErrorf(w http.ResponseWriter, format string, args ...any) error {
	if isErr(s) {
		http.Error(w, fmt.Sprintf(format, args...), int(s))
		return nil
	}
	return fmt.Errorf("status code [%d] is not an error", s)
}

func isErr(s Status) bool { return int(s) >= 400 }

func (s Status) Explanation() string {
	switch s {
	case Continue:
		return "Continue: The client should continue with its request."
	case Switching_Protocols:
		return "Switching Protocols: Server is switching protocols as requested."
	case Processing:
		return "Processing: Request received, no response available yet."
	case Early_Hints:
		return "Early Hints: Preload resources while server prepares the response."
	case OK:
		return "OK: The request was successful."
	case Created:
		return "Created: A new resource was successfully created."
	case Accepted:
		return "Accepted: Request received but not yet acted upon."
	case Non_Authoritative_Information:
		return "Non-Authoritative Information: Metadata from a third-party copy."
	case No_Content:
		return "No Content: No content to send for this request."
	case Reset_Content:
		return "Reset Content: Reset the document that made this request."
	case Partial_Content:
		return "Partial Content: Partial data response to a range request."
	case Multi_Status:
		return "Multi-Status: Information about multiple resources."
	case Already_Reported:
		return "Already Reported: Binding members already enumerated."
	case IM_Used:
		return "IM Used: Instance manipulations applied to current instance."
	case Multiple_Choices:
		return "Multiple Choices: Multiple possible responses. User or client must choose."
	case Moved_Permanently:
		return "Moved Permanently: Resource permanently moved to a new URL."
	case Found:
		return "Found: Resource temporarily moved to a new URL."
	case See_Other:
		return "See Other: Redirect to another URI using a GET request."
	case Not_Modified:
		return "Not Modified: Cached response is still valid."
	case Use_Proxy:
		return "Use Proxy: Deprecated; resource must be accessed through a proxy."
	case Unused:
		return "Unused: Reserved for future use."
	case Temporary_Redirect:
		return "Temporary Redirect: Temporarily redirect to a new URI, same method."
	case Permanent_Redirect:
		return "Permanent Redirect: Permanently redirect to a new URI, same method."
	case Bad_Request:
		return "Bad Request: The server cannot process the request."
	case Unauthorized:
		return "Unauthorized: Authentication is required."
	case Payment_Required:
		return "Payment Required: Reserved for future use."
	case Forbidden:
		return "Forbidden: The client does not have access rights."
	case Not_Found:
		return "Not Found: The requested resource could not be found."
	case Method_Not_Allowed:
		return "Method Not Allowed: Request method not supported."
	case Not_Acceptable:
		return "Not Acceptable: No content meets the criteria."
	case Proxy_Authentication_Required:
		return "Proxy Authentication Required: Authenticate with the proxy."
	case Request_Timeout:
		return "Request Timeout: Server timed out waiting for the request."
	case Conflict:
		return "Conflict: The request conflicts with server state."
	case Gone:
		return "Gone: The requested content is permanently unavailable."
	case Length_Required:
		return "Length Required: Content-Length header is missing."
	case Precondition_Failed:
		return "Precondition Failed: Preconditions in headers not met."
	case Content_Too_Large:
		return "Content Too Large: Request body exceeds server limits."
	case URI_Too_Long:
		return "URI Too Long: The URI is too long for the server to process."
	case Unsupported_Media_Type:
		return "Unsupported Media Type: The server cannot process the media format."
	case Range_Not_Satisfiable:
		return "Range Not Satisfiable: The requested range is invalid."
	case Expectation_Failed:
		return "Expectation Failed: The expectation cannot be met."
	case Im_a_teapot:
		return "I'm a teapot: The server refuses to brew coffee with a teapot."
	case Misdirected_Request:
		return "Misdirected Request: The request was directed to the wrong server."
	case Unprocessable_Content:
		return "Unprocessable Content: Request was well-formed but contains semantic errors."
	case Locked:
		return "Locked: The resource is locked."
	case Failed_Dependency:
		return "Failed Dependency: Request failed due to a prior failed request."
	case Too_Early:
		return "Too Early: Server is unwilling to process a request that might be replayed."
	case Upgrade_Required:
		return "Upgrade Required: Upgrade to a different protocol is required."
	case Precondition_Required:
		return "Precondition Required: Request must be conditional to prevent conflicts."
	case Too_Many_Requests:
		return "Too Many Requests: Rate limit exceeded."
	case Request_Header_Fields_Too_Large:
		return "Request Header Fields Too Large: Headers are too large to process."
	case Unavailable_For_Legal_Reasons:
		return "Unavailable For Legal Reasons: Resource cannot be provided for legal reasons."
	case Internal_Server_Error:
		return "Internal Server Error: Generic server error."
	case Not_Implemented:
		return "Not Implemented: The request method is not supported."
	case Bad_Gateway:
		return "Bad Gateway: Invalid response from an upstream server."
	case Service_Unavailable:
		return "Service Unavailable: Server cannot handle the request."
	case Gateway_Timeout:
		return "Gateway Timeout: No response from an upstream server."
	case HTTP_Version_Not_Supported:
		return "HTTP Version Not Supported: HTTP version not supported by the server."
	case Variant_Also_Negotiates:
		return "Variant Also Negotiates: Internal configuration error."
	case Insufficient_Storage:
		return "Insufficient Storage: Server unable to store representation."
	case Loop_Detected:
		return "Loop Detected: Infinite loop detected while processing request."
	case Not_Extended:
		return "Not Extended: HTTP Extension not supported."
	case Network_Authentication_Required:
		return "Network Authentication Required: Authentication needed to access network."
	default:
		panic(fmt.Sprintf("Unknown status code: %d", s))
	}
}

type LogHandler struct {
	http.ResponseWriter
	r     *http.Request
	start time.Time
}

func (h *LogHandler) WriteHeader(code int) {
	h.ResponseWriter.WriteHeader(code)
	if code >= 400 {
		yap.Error("request", "method", h.r.Method, "URL", h.r.RequestURI, "time", time.Since(h.start), "status", code)
	} else {
		yap.Info("request", "method", h.r.Method, "URL", h.r.RequestURI, "time", time.Since(h.start), "status", code)
	}
}

func LogRequests(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(&LogHandler{w, r, start}, r)
	})
}

func Pre(pref string) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return http.StripPrefix(pref, http.HandlerFunc(next)).ServeHTTP
	}
}
