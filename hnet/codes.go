package hnet

// Status represents an HTTP status code.
type Status int

const (
	// Informational responses (100–199)
	Continue           Status = 100 // The server has received the request headers and the client should proceed to send the request body.
	SwitchingProtocols Status = 101 // The requester has asked the server to switch protocols and the server has agreed to do so.
	Processing         Status = 102 // WebDAV: The server has received and is processing the request, but no response is available yet.
	EarlyHints         Status = 103 // Used to return some response headers before final HTTP message.

	// Successful responses (200–299)
	OK                   Status = 200 // The request was successful.
	Created              Status = 201 // The request was successful and a resource was created.
	Accepted             Status = 202 // The request has been accepted for processing, but the processing is not complete.
	NonAuthoritativeInfo Status = 203 // The request was successful but the returned meta-information is not from the origin server.
	NoContent            Status = 204 // The server successfully processed the request, and is not returning any content.
	ResetContent         Status = 205 // The server successfully processed the request and is asking the client to reset the document view.
	PartialContent       Status = 206 // The server is delivering only part of the resource due to a range header sent by the client.
	MultiStatus          Status = 207 // WebDAV: The message body contains multiple status codes for different operations.
	AlreadyReported      Status = 208 // WebDAV: The members of a DAV binding have already been enumerated.
	IMUsed               Status = 226 // The server has fulfilled a GET request for the resource, and the response is a representation of the result of one or more instance manipulations.

	// Redirection messages (300–399)
	MultipleChoices   Status = 300 // The request has more than one possible response. User-agent or user should choose one of them.
	MovedPermanently  Status = 301 // The URL of the requested resource has been changed permanently.
	Found             Status = 302 // The requested resource resides temporarily under a different URL.
	SeeOther          Status = 303 // The server is redirecting the client to a different resource.
	NotModified       Status = 304 // Indicates that the resource has not been modified since the version specified by the request headers.
	UseProxy          Status = 305 // The requested resource is only available through a proxy.
	TemporaryRedirect Status = 307 // The request should be repeated with another URL; however, future requests should still use the original URL.
	PermanentRedirect Status = 308 // The request and all future requests should be repeated using another URL.

	// Client error responses (400–499)
	BadRequest                  Status = 400 // The server could not understand the request due to invalid syntax.
	Unauthorized                Status = 401 // The client must authenticate itself to get the requested response.
	PaymentRequired             Status = 402 // Reserved for future use.
	Forbidden                   Status = 403 // The client does not have access rights to the content.
	NotFound                    Status = 404 // The server cannot find the requested resource.
	MethodNotAllowed            Status = 405 // The method specified in the request is not allowed.
	NotAcceptable               Status = 406 // The server cannot produce a response matching the accept headers of the request.
	ProxyAuthRequired           Status = 407 // The client must first authenticate itself with the proxy.
	RequestTimeout              Status = 408 // The server timed out waiting for the request.
	Conflict                    Status = 409 // The request could not be completed due to a conflict with the current state of the resource.
	Gone                        Status = 410 // The resource requested is no longer available and will not be available again.
	LengthRequired              Status = 411 // The server rejected the request because the Content-Length header field is not defined.
	PreconditionFailed          Status = 412 // The client has indicated preconditions in its headers which the server does not meet.
	PayloadTooLarge             Status = 413 // The request entity is larger than the server is willing or able to process.
	URITooLong                  Status = 414 // The request URI is longer than the server is willing to interpret.
	UnsupportedMediaType        Status = 415 // The media format of the requested data is not supported by the server.
	RangeNotSatisfiable         Status = 416 // The range specified by the Range header field in the request cannot be fulfilled.
	ExpectationFailed           Status = 417 // The server cannot meet the requirements of the Expect request-header field.
	ImATeapot                   Status = 418 // The server refuses to brew coffee because it is, permanently, a teapot.
	MisdirectedRequest          Status = 421 // The request was directed at a server that is not able to produce a response.
	UnprocessableEntity         Status = 422 // WebDAV: The request was well-formed but could not be followed due to semantic errors.
	Locked                      Status = 423 // WebDAV: The resource being accessed is locked.
	FailedDependency            Status = 424 // WebDAV: The request failed because it depended on another request that failed.
	TooEarly                    Status = 425 // The server is unwilling to risk processing a request that might be replayed.
	UpgradeRequired             Status = 426 // The client should switch to a different protocol.
	PreconditionRequired        Status = 428 // The server requires the request to be conditional.
	TooManyRequests             Status = 429 // The client has sent too many requests in a given amount of time.
	RequestHeaderFieldsTooLarge Status = 431 // The server is unwilling to process the request because its header fields are too large.
	UnavailableForLegalReasons  Status = 451 // The requested resource is unavailable due to legal reasons.

	// Server error responses (500–599)
	InternalServerError           Status = 500 // The server has encountered a situation it doesn't know how to handle.
	NotImplemented                Status = 501 // The server does not support the functionality required to fulfill the request.
	BadGateway                    Status = 502 // The server received an invalid response from the upstream server.
	ServiceUnavailable            Status = 503 // The server is not ready to handle the request.
	GatewayTimeout                Status = 504 // The server did not get a response in time from an upstream server.
	HTTPVersionNotSupported       Status = 505 // The server does not support the HTTP protocol version used in the request.
	VariantAlsoNegotiates         Status = 506 // Transparent content negotiation for the request results in a circular reference.
	InsufficientStorage           Status = 507 // WebDAV: The server is unable to store the representation needed to complete the request.
	LoopDetected                  Status = 508 // WebDAV: The server detected an infinite loop while processing the request.
	NotExtended                   Status = 510 // Further extensions to the request are required for the server to fulfill it.
	NetworkAuthenticationRequired Status = 511 // The client needs to authenticate to gain network access.
)

func Def(status int) NetErr { return Status(status).Def() }

func (s Status) Err(msg string, pairs ...string) NetErr { return Free(int(s), msg, pairs...) }
func (s Status) Def(pairs ...string) NetErr {
	if s < 400 {
		return nil
	}
	switch s {
	case BadRequest:
		return s.Err("Bad Request", pairs...)
	case Unauthorized:
		return s.Err("Unauthorized - Authentication required", pairs...)
	case PaymentRequired:
		return s.Err("Payment Required", pairs...)
	case Forbidden:
		return s.Err("Forbidden - Access denied", pairs...)
	case NotFound:
		return s.Err("Not Found - Resource not available", pairs...)
	case MethodNotAllowed:
		return s.Err("Method Not Allowed", pairs...)
	case NotAcceptable:
		return s.Err("Not Acceptable - Unavailable in requested format", pairs...)
	case ProxyAuthRequired:
		return s.Err("Proxy Authentication Required", pairs...)
	case RequestTimeout:
		return s.Err("Request Timeout", pairs...)
	case Conflict:
		return s.Err("Conflict - Request could not be completed due to a conflict", pairs...)
	case Gone:
		return s.Err("Gone - Resource is no longer available", pairs...)
	case LengthRequired:
		return s.Err("Length Required - Content-Length header is missing", pairs...)
	case PreconditionFailed:
		return s.Err("Precondition Failed", pairs...)
	case PayloadTooLarge:
		return s.Err("Payload Too Large", pairs...)
	case URITooLong:
		return s.Err("URI Too Long", pairs...)
	case UnsupportedMediaType:
		return s.Err("Unsupported Media Type", pairs...)
	case RangeNotSatisfiable:
		return s.Err("Range Not Satisfiable", pairs...)
	case ExpectationFailed:
		return s.Err("Expectation Failed", pairs...)
	case ImATeapot:
		return s.Err("I'm a teapot - Fun Easter egg", pairs...)
	case UnprocessableEntity:
		return s.Err("Unprocessable Entity", pairs...)
	case Locked:
		return s.Err("Locked - Resource is locked", pairs...)
	case FailedDependency:
		return s.Err("Failed Dependency", pairs...)
	case UpgradeRequired:
		return s.Err("Upgrade Required - Switch to a different protocol", pairs...)
	case PreconditionRequired:
		return s.Err("Precondition Required", pairs...)
	case TooManyRequests:
		return s.Err("Too Many Requests - Rate limit exceeded", pairs...)
	case RequestHeaderFieldsTooLarge:
		return s.Err("Request Header Fields Too Large", pairs...)
	case UnavailableForLegalReasons:
		return s.Err("Unavailable For Legal Reasons", pairs...)

	// Server error responses
	case InternalServerError:
		return s.Err("Internal Server Error", pairs...)
	case NotImplemented:
		return s.Err("Not Implemented - Feature not supported", pairs...)
	case BadGateway:
		return s.Err("Bad Gateway - Invalid response from upstream server", pairs...)
	case ServiceUnavailable:
		return s.Err("Service Unavailable - Try again later", pairs...)
	case GatewayTimeout:
		return s.Err("Gateway Timeout - Upstream server failed to respond", pairs...)
	case HTTPVersionNotSupported:
		return s.Err("HTTP Version Not Supported", pairs...)
	case VariantAlsoNegotiates:
		return s.Err("Variant Also Negotiates - Circular reference detected", pairs...)
	case InsufficientStorage:
		return s.Err("Insufficient Storage - Unable to store data", pairs...)
	case LoopDetected:
		return s.Err("Loop Detected", pairs...)
	case NotExtended:
		return s.Err("Not Extended - Further extensions required", pairs...)
	case NetworkAuthenticationRequired:
		return s.Err("Network Authentication Required", pairs...)
	default:
		return s.Err("Unknown Error", pairs...)
	}
}
