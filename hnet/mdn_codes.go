// This file is generated directly from the from the MDN Web Docs.
// The script formats the licensed prose content to code comments.
// See the script responsible for this in the docs/js directory of this repository.
// The documentation in this file is available under the CC-BY-SA 2.5 license,
// as is all prose content on MDN Web Docs.
//
// Attribution:
// - Source: MDN Web Docs (https://developer.mozilla.org/)
// - License: CC-BY-SA 2.5 (https://creativecommons.org/licenses/by-sa/2.5/)
// - Contributors: https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/contributors.txt
// - Source file:  https://github.com/mdn/content/blob/main/files/en-us/web/http/status/index.md?plain=1
// - Source repo:  https://github.com/mdn/content
//
// At time of generation, the source file was last modified on Oct 18, 2024.
package hnet

// Informational responses
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status#informational_responses
const (
	// Continue: This interim response indicates that the client should continue
	// the request or ignore the response if the request is already finished.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/100
	Continue Status = 100
	// Switching_Protocols: This code is sent in response to an Upgrade request
	// header from the client and indicates the protocol the server is switching to.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/101
	Switching_Protocols Status = 101
	// Processing: This code was used in WebDAV contexts to indicate that a request
	// has been received by the server, but no status was available at the time of the
	// response.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/102
	Processing Status = 102
	// Early_Hints: This status code is primarily intended to be used with the Link
	// header, letting the user agent start preloading resources while the server
	// prepares a response or preconnect to an origin from which the page will need
	// resources.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/103
	Early_Hints Status = 103
)

// Successful responses
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status#successful_responses
const (
	// OK: The request succeeded. The result and meaning of "success" depends on
	// the HTTP method: GET: The resource has been fetched and transmitted in the
	// message body. HEAD: Representation headers are included in the response without
	// any message body. PUT or POST: The resource describing the result of the action
	// is transmitted in the message body. TRACE: The message body contains the request
	// as received by the server.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/200
	OK Status = 200
	// Created: The request succeeded, and a new resource was created as a result.
	// This is typically the response sent after POST requests, or some PUT requests.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/201
	Created Status = 201
	// Accepted: The request has been received but not yet acted upon. It is
	// noncommittal, since there is no way in HTTP to later send an asynchronous
	// response indicating the outcome of the request. It is intended for cases where
	// another process or server handles the request, or for batch processing.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/202
	Accepted Status = 202
	// Non_Authoritative_Information: This response code means the returned
	// metadata is not exactly the same as is available from the origin server, but is
	// collected from a local or a third-party copy. This is mostly used for mirrors or
	// backups of another resource. Except for that specific case, the 200 OK response
	// is preferred to this status.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/203
	Non_Authoritative_Information Status = 203
	// No_Content: There is no content to send for this request, but the headers
	// are useful. The user agent may update its cached headers for this resource with
	// the new ones.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/204
	No_Content Status = 204
	// Reset_Content: Tells the user agent to reset the document which sent this
	// request.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/205
	Reset_Content Status = 205
	// Partial_Content: This response code is used in response to a range request
	// when the client has requested a part or parts of a resource.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/206
	Partial_Content Status = 206
	// Multi_Status: Conveys information about multiple resources, for situations
	// where multiple status codes might be appropriate.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/207
	Multi_Status Status = 207
	// Already_Reported: Used inside a <dav:propstat> response element to avoid
	// repeatedly enumerating the internal members of multiple bindings to the same
	// collection.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/208
	Already_Reported Status = 208
	// IM_Used: The server has fulfilled a GET request for the resource, and the
	// response is a representation of the result of one or more instance-manipulations
	// applied to the current instance.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/226
	IM_Used Status = 226
)

// Redirection messages
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status#redirection_messages
const (
	// Multiple_Choices: In agent-driven content negotiation, the request has more
	// than one possible response and the user agent or user should choose one of them.
	// There is no standardized way for clients to automatically choose one of the
	// responses, so this is rarely used.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/300
	Multiple_Choices Status = 300
	// Moved_Permanently: The URL of the requested resource has been changed
	// permanently. The new URL is given in the response.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/301
	Moved_Permanently Status = 301
	// Found: This response code means that the URI of requested resource has been
	// changed temporarily. Further changes in the URI might be made in the future, so
	// the same URI should be used by the client in future requests.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/302
	Found Status = 302
	// See_Other: The server sent this response to direct the client to get the
	// requested resource at another URI with a GET request.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/303
	See_Other Status = 303
	// Not_Modified: This is used for caching purposes. It tells the client that
	// the response has not been modified, so the client can continue to use the same
	// cached version of the response.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/304
	Not_Modified Status = 304
	// Use_Proxy: Defined in a previous version of the HTTP specification to
	// indicate that a requested response must be accessed by a proxy. It has been
	// deprecated due to security concerns regarding in-band configuration of a proxy.
	// https://developer.mozilla.org#305_use_proxy
	Use_Proxy Status = 305
	// Unused: This response code is no longer used; but is reserved. It was used
	// in a previous version of the HTTP/1.1 specification.
	// https://developer.mozilla.org#306_unused
	Unused Status = 306
	// Temporary_Redirect: The server sends this response to direct the client to
	// get the requested resource at another URI with the same method that was used in
	// the prior request. This has the same semantics as the 302 Found response code,
	// with the exception that the user agent must not change the HTTP method used: if a
	// POST was used in the first request, a POST must be used in the redirected
	// request.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/307
	Temporary_Redirect Status = 307
	// Permanent_Redirect: This means that the resource is now permanently located
	// at another URI, specified by the Location response header. This has the same
	// semantics as the 301 Moved Permanently HTTP response code, with the exception
	// that the user agent must not change the HTTP method used: if a POST was used in
	// the first request, a POST must be used in the second request.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/308
	Permanent_Redirect Status = 308
)

// Client error responses
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status#client_error_responses
const (
	// Bad_Request: The server cannot or will not process the request due to
	// something that is perceived to be a client error (e.g., malformed request syntax,
	// invalid request message framing, or deceptive request routing).
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400
	Bad_Request Status = 400
	// Unauthorized: Although the HTTP standard specifies "unauthorized",
	// semantically this response means "unauthenticated". That is, the client must
	// authenticate itself to get the requested response.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/401
	Unauthorized Status = 401
	// Payment_Required: The initial purpose of this code was for digital payment
	// systems, however this status code is rarely used and no standard convention
	// exists.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/402
	Payment_Required Status = 402
	// Forbidden: The client does not have access rights to the content; that is,
	// it is unauthorized, so the server is refusing to give the requested resource.
	// Unlike 401 Unauthorized, the client's identity is known to the server.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/403
	Forbidden Status = 403
	// Not_Found: The server cannot find the requested resource. In the browser,
	// this means the URL is not recognized. In an API, this can also mean that the
	// endpoint is valid but the resource itself does not exist. Servers may also send
	// this response instead of 403 Forbidden to hide the existence of a resource from
	// an unauthorized client. This response code is probably the most well known due to
	// its frequent occurrence on the web.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404
	Not_Found Status = 404
	// Method_Not_Allowed: The request method is known by the server but is not
	// supported by the target resource. For example, an API may not allow DELETE on a
	// resource, or the TRACE method entirely.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/405
	Method_Not_Allowed Status = 405
	// Not_Acceptable: This response is sent when the web server, after performing
	// server-driven content negotiation, doesn't find any content that conforms to the
	// criteria given by the user agent.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/406
	Not_Acceptable Status = 406
	// Proxy_Authentication_Required: This is similar to 401 Unauthorized but
	// authentication is needed to be done by a proxy.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/407
	Proxy_Authentication_Required Status = 407
	// Request_Timeout: This response is sent on an idle connection by some
	// servers, even without any previous request by the client. It means that the
	// server would like to shut down this unused connection. This response is used much
	// more since some browsers use HTTP pre-connection mechanisms to speed up browsing.
	// Some servers may shut down a connection without sending this message.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/408
	Request_Timeout Status = 408
	// Conflict: This response is sent when a request conflicts with the current
	// state of the server. In WebDAV remote web authoring, 409 responses are errors
	// sent to the client so that a user might be able to resolve a conflict and
	// resubmit the request.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/409
	Conflict Status = 409
	// Gone: This response is sent when the requested content has been permanently
	// deleted from server, with no forwarding address. Clients are expected to remove
	// their caches and links to the resource. The HTTP specification intends this
	// status code to be used for "limited-time, promotional services". APIs should not
	// feel compelled to indicate resources that have been deleted with this status
	// code.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/410
	Gone Status = 410
	// Length_Required: Server rejected the request because the Content-Length
	// header field is not defined and the server requires it.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/411
	Length_Required Status = 411
	// Precondition_Failed: In conditional requests, the client has indicated
	// preconditions in its headers which the server does not meet.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/412
	Precondition_Failed Status = 412
	// Content_Too_Large: The request body is larger than limits defined by server.
	// The server might close the connection or return an Retry-After header field.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/413
	Content_Too_Large Status = 413
	// URI_Too_Long: The URI requested by the client is longer than the server is
	// willing to interpret.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/414
	URI_Too_Long Status = 414
	// Unsupported_Media_Type: The media format of the requested data is not
	// supported by the server, so the server is rejecting the request.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/415
	Unsupported_Media_Type Status = 415
	// Range_Not_Satisfiable: The ranges specified by the Range header field in the
	// request cannot be fulfilled. It's possible that the range is outside the size of
	// the target resource's data.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/416
	Range_Not_Satisfiable Status = 416
	// Expectation_Failed: This response code means the expectation indicated by
	// the Expect request header field cannot be met by the server.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/417
	Expectation_Failed Status = 417
	// Im_a_teapot: The server refuses the attempt to brew coffee with a teapot.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/418
	Im_a_teapot Status = 418
	// Misdirected_Request: The request was directed at a server that is not able
	// to produce a response. This can be sent by a server that is not configured to
	// produce responses for the combination of scheme and authority that are included
	// in the request URI.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/421
	Misdirected_Request Status = 421
	// Unprocessable_Content: The request was well-formed but was unable to be
	// followed due to semantic errors.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422
	Unprocessable_Content Status = 422
	// Locked: The resource that is being accessed is locked.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/423
	Locked Status = 423
	// Failed_Dependency: The request failed due to failure of a previous request.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/424
	Failed_Dependency Status = 424
	// Too_Early: Indicates that the server is unwilling to risk processing a
	// request that might be replayed.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/425
	Too_Early Status = 425
	// Upgrade_Required: The server refuses to perform the request using the
	// current protocol but might be willing to do so after the client upgrades to a
	// different protocol. The server sends an Upgrade header in a 426 response to
	// indicate the required protocol(s).
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/426
	Upgrade_Required Status = 426
	// Precondition_Required: The origin server requires the request to be
	// conditional. This response is intended to prevent the 'lost update' problem,
	// where a client GETs a resource's state, modifies it and PUTs it back to the
	// server, when meanwhile a third party has modified the state on the server,
	// leading to a conflict.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/428
	Precondition_Required Status = 428
	// Too_Many_Requests: The user has sent too many requests in a given amount of
	// time (rate limiting).
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/429
	Too_Many_Requests Status = 429
	// Request_Header_Fields_Too_Large: The server is unwilling to process the
	// request because its header fields are too large. The request may be resubmitted
	// after reducing the size of the request header fields.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/431
	Request_Header_Fields_Too_Large Status = 431
	// Unavailable_For_Legal_Reasons: The user agent requested a resource that
	// cannot legally be provided, such as a web page censored by a government.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/451
	Unavailable_For_Legal_Reasons Status = 451
)

// Server error responses
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status#server_error_responses
const (
	// Internal_Server_Error: The server has encountered a situation it does not
	// know how to handle. This error is generic, indicating that the server cannot find
	// a more appropriate 5XX status code to respond with.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500
	Internal_Server_Error Status = 500
	// Not_Implemented: The request method is not supported by the server and
	// cannot be handled. The only methods that servers are required to support (and
	// therefore that must not return this code) are GET and HEAD.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/501
	Not_Implemented Status = 501
	// Bad_Gateway: This error response means that the server, while working as a
	// gateway to get a response needed to handle the request, got an invalid response.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/502
	Bad_Gateway Status = 502
	// Service_Unavailable: The server is not ready to handle the request. Common
	// causes are a server that is down for maintenance or that is overloaded. Note that
	// together with this response, a user-friendly page explaining the problem should
	// be sent. This response should be used for temporary conditions and the
	// Retry-After HTTP header should, if possible, contain the estimated time before
	// the recovery of the service. The webmaster must also take care about the
	// caching-related headers that are sent along with this response, as these
	// temporary condition responses should usually not be cached.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/503
	Service_Unavailable Status = 503
	// Gateway_Timeout: This error response is given when the server is acting as a
	// gateway and cannot get a response in time.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/504
	Gateway_Timeout Status = 504
	// HTTP_Version_Not_Supported: The HTTP version used in the request is not
	// supported by the server.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/505
	HTTP_Version_Not_Supported Status = 505
	// Variant_Also_Negotiates: The server has an internal configuration error:
	// during content negotiation, the chosen variant is configured to engage in content
	// negotiation itself, which results in circular references when creating responses.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/506
	Variant_Also_Negotiates Status = 506
	// Insufficient_Storage: The method could not be performed on the resource
	// because the server is unable to store the representation needed to successfully
	// complete the request.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/507
	Insufficient_Storage Status = 507
	// Loop_Detected: The server detected an infinite loop while processing the
	// request.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/508
	Loop_Detected Status = 508
	// Not_Extended: The client request declares an HTTP Extension (RFC 2774) that
	// should be used to process the request, but the extension is not supported.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/510
	Not_Extended Status = 510
	// Network_Authentication_Required: Indicates that the client needs to
	// authenticate to gain network access.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/511
	Network_Authentication_Required Status = 511
)
