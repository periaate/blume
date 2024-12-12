// This file is generated directly from the from the MDN Web Docs.
// The script formats the licensed prose content to code comments.
// See the script responsible for this in the docs/js directory of this repository.
// The documentation in this file is available under the CC-BY-SA 2.5 license,
// as is all prose content on MDN Web Docs.
//
// Attribution:
// - Source: MDN Web Docs (https://developer.mozilla.org/)
// - License: CC-BY-SA 2.5 (https://creativecommons.org/licenses/by-sa/2.5/)
// - Contributors: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/contributors.txt
// - Source file:  https://github.com/mdn/content/blob/main/files/en-us/web/http/headers/index.md?plain=1
// - Source repo:  https://github.com/mdn/content
//
// At time of generation, the source file was last modified on Nov 25, 2024.
package hnet

// Authentication
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#authentication
const (
	// WWW_Authenticate: Defines the authentication method that should be used to
	// access a resource.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/WWW-Authenticate
	WWW_Authenticate Header = "WWW-Authenticate"
	// Authorization: Contains the credentials to authenticate a user-agent with a
	// server.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Authorization
	Authorization Header = "Authorization"
	// Proxy_Authenticate: Defines the authentication method that should be used to
	// access a resource behind a proxy server.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Proxy-Authenticate
	Proxy_Authenticate Header = "Proxy-Authenticate"
	// Proxy_Authorization: Contains the credentials to authenticate a user agent
	// with a proxy server.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Proxy-Authorization
	Proxy_Authorization Header = "Proxy-Authorization"
)

// Caching
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#caching
const (
	// Age: The time, in seconds, that the object has been in a proxy cache.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Age
	Age Header = "Age"
	// Cache_Control: Directives for caching mechanisms in both requests and
	// responses.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Cache-Control
	Cache_Control Header = "Cache-Control"
	// Clear_Site_Data: Clears browsing data (e.g. cookies, storage, cache)
	// associated with the requesting website.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Clear-Site-Data
	Clear_Site_Data Header = "Clear-Site-Data"
	// Expires: The date/time after which the response is considered stale.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Expires
	Expires Header = "Expires"
	// No_Vary_Search: Specifies a set of rules that define how a URL's query
	// parameters will affect cache matching. These rules dictate whether the same URL
	// with different URL parameters should be saved as separate browser cache entries.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/No-Vary-Search
	No_Vary_Search Header = "No-Vary-Search"
)

// Conditionals
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#conditionals
const (
	// Last_Modified: The last modification date of the resource, used to compare
	// several versions of the same resource. It is less accurate than ETag, but easier
	// to calculate in some environments. Conditional requests using If-Modified-Since
	// and If-Unmodified-Since use this value to change the behavior of the request.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Last-Modified
	Last_Modified Header = "Last-Modified"
	// ETag: A unique string identifying the version of the resource. Conditional
	// requests using If-Match and If-None-Match use this value to change the behavior
	// of the request.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/ETag
	ETag Header = "ETag"
	// If_Match: Makes the request conditional, and applies the method only if the
	// stored resource matches one of the given ETags.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/If-Match
	If_Match Header = "If-Match"
	// If_None_Match: Makes the request conditional, and applies the method only if
	// the stored resource doesn't match any of the given ETags. This is used to update
	// caches (for safe requests), or to prevent uploading a new resource when one
	// already exists.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/If-None-Match
	If_None_Match Header = "If-None-Match"
	// If_Modified_Since: Makes the request conditional, and expects the resource
	// to be transmitted only if it has been modified after the given date. This is used
	// to transmit data only when the cache is out of date.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/If-Modified-Since
	If_Modified_Since Header = "If-Modified-Since"
	// If_Unmodified_Since: Makes the request conditional, and expects the resource
	// to be transmitted only if it has not been modified after the given date. This
	// ensures the coherence of a new fragment of a specific range with previous ones,
	// or to implement an optimistic concurrency control system when modifying existing
	// documents.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/If-Unmodified-Since
	If_Unmodified_Since Header = "If-Unmodified-Since"
	// Vary: Determines how to match request headers to decide whether a cached
	// response can be used rather than requesting a fresh one from the origin server.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Vary
	Vary Header = "Vary"
)

// Connection management
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#connection_management
const (
	// Connection: Controls whether the network connection stays open after the
	// current transaction finishes.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Connection
	Connection Header = "Connection"
	// Keep_Alive: Controls how long a persistent connection should stay open.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Keep-Alive
	Keep_Alive Header = "Keep-Alive"
)

// Content negotiation
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#content_negotiation
const (
	// Accept: Informs the server about the types of data that can be sent back.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Accept
	Accept Header = "Accept"
	// Accept_Encoding: The encoding algorithm, usually a compression algorithm,
	// that can be used on the resource sent back.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Accept-Encoding
	Accept_Encoding Header = "Accept-Encoding"
	// Accept_Language: Informs the server about the human language the server is
	// expected to send back. This is a hint and is not necessarily under the full
	// control of the user: the server should always pay attention not to override an
	// explicit user choice (like selecting a language from a dropdown).
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Accept-Language
	Accept_Language Header = "Accept-Language"
	// Accept_Patch: A request content negotiation response header that advertises
	// which media type the server is able to understand in a PATCH request.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Accept-Patch
	Accept_Patch Header = "Accept-Patch"
	// Accept_Post: A request content negotiation response header that advertises
	// which media type the server is able to understand in a POST request.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Accept-Post
	Accept_Post Header = "Accept-Post"
)

// Controls
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#controls
const (
	// Expect: Indicates expectations that need to be fulfilled by the server to
	// properly handle the request.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Expect
	Expect Header = "Expect"
	// Max_Forwards: When using TRACE, indicates the maximum number of hops the
	// request can do before being reflected to the sender.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Max-Forwards
	Max_Forwards Header = "Max-Forwards"
)

// Cookies
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#cookies
const (
	// Cookie: Contains stored HTTP cookies previously sent by the server with the
	// Set-Cookie header.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Cookie
	Cookie Header = "Cookie"
	// Set_Cookie: Send cookies from the server to the user-agent.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Set-Cookie
	Set_Cookie Header = "Set-Cookie"
)

// CORS
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#cors
const (
	// Access_Control_Allow_Credentials: Indicates whether the response to the
	// request can be exposed when the credentials flag is true.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Credentials
	Access_Control_Allow_Credentials Header = "Access-Control-Allow-Credentials"
	// Access_Control_Allow_Headers: Used in response to a preflight request to
	// indicate which HTTP headers can be used when making the actual request.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Headers
	Access_Control_Allow_Headers Header = "Access-Control-Allow-Headers"
	// Access_Control_Allow_Methods: Specifies the methods allowed when accessing
	// the resource in response to a preflight request.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Methods
	Access_Control_Allow_Methods Header = "Access-Control-Allow-Methods"
	// Access_Control_Allow_Origin: Indicates whether the response can be shared.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Origin
	Access_Control_Allow_Origin Header = "Access-Control-Allow-Origin"
	// Access_Control_Expose_Headers: Indicates which headers can be exposed as
	// part of the response by listing their names.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Expose-Headers
	Access_Control_Expose_Headers Header = "Access-Control-Expose-Headers"
	// Access_Control_Max_Age: Indicates how long the results of a preflight
	// request can be cached.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Max-Age
	Access_Control_Max_Age Header = "Access-Control-Max-Age"
	// Access_Control_Request_Headers: Used when issuing a preflight request to let
	// the server know which HTTP headers will be used when the actual request is made.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Request-Headers
	Access_Control_Request_Headers Header = "Access-Control-Request-Headers"
	// Access_Control_Request_Method: Used when issuing a preflight request to let
	// the server know which HTTP method will be used when the actual request is made.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Request-Method
	Access_Control_Request_Method Header = "Access-Control-Request-Method"
	// Origin: Indicates where a fetch originates from.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Origin
	Origin Header = "Origin"
	// Timing_Allow_Origin: Specifies origins that are allowed to see values of
	// attributes retrieved via features of the Resource Timing API, which would
	// otherwise be reported as zero due to cross-origin restrictions.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Timing-Allow-Origin
	Timing_Allow_Origin Header = "Timing-Allow-Origin"
)

// Downloads
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#downloads
const (
	// Content_Disposition: Indicates if the resource transmitted should be
	// displayed inline (default behavior without the header), or if it should be
	// handled like a download and the browser should present a "Save As" dialog.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Disposition
	Content_Disposition Header = "Content-Disposition"
)

// Integrity digests
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#integrity_digests
const (
	// Content_Digest: Provides a digest of the stream of octets framed in an HTTP
	// message (the message content) dependent on Content-Encoding and Content-Range.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Digest
	Content_Digest Header = "Content-Digest"
	// Repr_Digest: Provides a digest of the selected representation of the target
	// resource before transmission. Unlike the Content-Digest, the digest does not
	// consider Content-Encoding or Content-Range.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Repr-Digest
	Repr_Digest Header = "Repr-Digest"
	// Want_Content_Digest: States the wish for a Content-Digest header. It is the
	// Content- analogue of Want-Repr-Digest.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Want-Content-Digest
	Want_Content_Digest Header = "Want-Content-Digest"
	// Want_Repr_Digest: States the wish for a Repr-Digest header. It is the Repr-
	// analogue of Want-Content-Digest.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Want-Repr-Digest
	Want_Repr_Digest Header = "Want-Repr-Digest"
)

// Message body information
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#message_body_information
const (
	// Content_Length: The size of the resource, in decimal number of bytes.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Length
	Content_Length Header = "Content-Length"
	// Content_Type: Indicates the media type of the resource.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Type
	Content_Type Header = "Content-Type"
	// Content_Encoding: Used to specify the compression algorithm.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Encoding
	Content_Encoding Header = "Content-Encoding"
	// Content_Language: Describes the human language(s) intended for the audience,
	// so that it allows a user to differentiate according to the users' own preferred
	// language.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Language
	Content_Language Header = "Content-Language"
	// Content_Location: Indicates an alternate location for the returned data.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Location
	Content_Location Header = "Content-Location"
)

// Proxies
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#proxies
const (
	// Forwarded: Contains information from the client-facing side of proxy servers
	// that is altered or lost when a proxy is involved in the path of the request.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Forwarded
	Forwarded Header = "Forwarded"
	// Via: Added by proxies, both forward and reverse proxies, and can appear in
	// the request headers and the response headers.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Via
	Via Header = "Via"
)

// Range requests
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#range_requests
const (
	// Accept_Ranges: Indicates if the server supports range requests, and if so in
	// which unit the range can be expressed.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Accept-Ranges
	Accept_Ranges Header = "Accept-Ranges"
	// Range: Indicates the part of a document that the server should return.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Range
	Range Header = "Range"
	// If_Range: Creates a conditional range request that is only fulfilled if the
	// given etag or date matches the remote resource. Used to prevent downloading two
	// ranges from incompatible version of the resource.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/If-Range
	If_Range Header = "If-Range"
	// Content_Range: Indicates where in a full body message a partial message
	// belongs.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Range
	Content_Range Header = "Content-Range"
)

// Redirects
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#redirects
const (
	// Location: Indicates the URL to redirect a page to.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Location
	Location Header = "Location"
	// Refresh: Directs the browser to reload the page or redirect to another.
	// Takes the same value as the meta element with http-equiv="refresh".
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Refresh
	Refresh Header = "Refresh"
)

// Request context
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#request_context
const (
	// From: Contains an Internet email address for a human user who controls the
	// requesting user agent.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/From
	From Header = "From"
	// Host: Specifies the domain name of the server (for virtual hosting), and
	// (optionally) the TCP port number on which the server is listening.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Host
	Host Header = "Host"
	// Referer: The address of the previous web page from which a link to the
	// currently requested page was followed.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Referer
	Referer Header = "Referer"
	// Referrer_Policy: Governs which referrer information sent in the Referer
	// header should be included with requests made.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Referrer-Policy
	Referrer_Policy Header = "Referrer-Policy"
	// User_Agent: Contains a characteristic string that allows the network
	// protocol peers to identify the application type, operating system, software
	// vendor or software version of the requesting software user agent.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/User-Agent
	User_Agent Header = "User-Agent"
)

// Response context
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#response_context
const (
	// Allow: Lists the set of HTTP request methods supported by a resource.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Allow
	Allow Header = "Allow"
	// Server: Contains information about the software used by the origin server to
	// handle the request.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Server
	Server Header = "Server"
)

// Security
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#security
const (
	// Cross_Origin_Embedder_Policy: Allows a server to declare an embedder policy
	// for a given document.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Cross-Origin-Embedder-Policy
	Cross_Origin_Embedder_Policy Header = "Cross-Origin-Embedder-Policy"
	// Cross_Origin_Opener_Policy: Prevents other domains from opening/controlling
	// a window.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Cross-Origin-Opener-Policy
	Cross_Origin_Opener_Policy Header = "Cross-Origin-Opener-Policy"
	// Cross_Origin_Resource_Policy: Prevents other domains from reading the
	// response of the resources to which this header is applied. See also CORP
	// explainer article.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Cross-Origin-Resource-Policy
	Cross_Origin_Resource_Policy Header = "Cross-Origin-Resource-Policy"
	// Content_Security_Policy: Controls resources the user agent is allowed to
	// load for a given page.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy
	Content_Security_Policy Header = "Content-Security-Policy"
	// Content_Security_Policy_Report_Only: Allows web developers to experiment
	// with policies by monitoring, but not enforcing, their effects. These violation
	// reports consist of JSON documents sent via an HTTP POST request to the specified
	// URI.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy-Report-Only
	Content_Security_Policy_Report_Only Header = "Content-Security-Policy-Report-Only"
	// Expect_CT: Lets sites opt in to reporting and enforcement of Certificate
	// Transparency to detect use of misissued certificates for that site.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Expect-CT
	Expect_CT Header = "Expect-CT"
	// Permissions_Policy: Provides a mechanism to allow and deny the use of
	// browser features in a website's own frame, and in <iframe>s that it embeds.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Permissions-Policy
	Permissions_Policy Header = "Permissions-Policy"
	// Reporting_Endpoints: Response header that allows website owners to specify
	// one or more endpoints used to receive errors such as CSP violation reports,
	// Cross-Origin-Opener-Policy reports, or other generic violations.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Reporting-Endpoints
	SSE_Reporting_Endpoints Header = "Reporting-Endpoints"
	// Strict_Transport_Security: Force communication using HTTPS instead of HTTP.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Strict-Transport-Security
	Strict_Transport_Security Header = "Strict-Transport-Security"
	// Upgrade_Insecure_Requests: Sends a signal to the server expressing the
	// client's preference for an encrypted and authenticated response, and that it can
	// successfully handle the upgrade-insecure-requests directive.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Upgrade-Insecure-Requests
	Upgrade_Insecure_Requests Header = "Upgrade-Insecure-Requests"
	// X_Content_Type_Options: Disables MIME sniffing and forces browser to use the
	// type given in Content-Type.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Content-Type-Options
	X_Content_Type_Options Header = "X-Content-Type-Options"
	// X_Frame_Options: Indicates whether a browser should be allowed to render a
	// page in a <frame>, <iframe>, <embed> or <object>.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Frame-Options
	X_Frame_Options Header = "X-Frame-Options"
	// X_Permitted_Cross_Domain_Policies: Specifies if a cross-domain policy file
	// (crossdomain.xml) is allowed. The file may define a policy to grant clients, such
	// as Adobe's Flash Player (now obsolete), Adobe Acrobat, Microsoft Silverlight (now
	// obsolete), or Apache Flex, permission to handle data across domains that would
	// otherwise be restricted due to the Same-Origin Policy. See the Cross-domain
	// Policy File Specification for more information.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#X-Permitted-Cross-Domain-Policies
	X_Permitted_Cross_Domain_Policies Header = "X-Permitted-Cross-Domain-Policies"
	// X_Powered_By: May be set by hosting environments or other frameworks and
	// contains information about them while not providing any usefulness to the
	// application or its visitors. Unset this header to avoid exposing potential
	// vulnerabilities.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#X-Powered-By
	X_Powered_By Header = "X-Powered-By"
	// X_XSS_Protection: Enables cross-site scripting filtering.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-XSS-Protection
	X_XSS_Protection Header = "X-XSS-Protection"
)

// Server-sent events
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#server-sent_events
const (
	// Reporting_Endpoints: Response header used to specify server endpoints where
	// the browser should send warning and error reports when using the Reporting API.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Reporting-Endpoints
	Reporting_Endpoints Header = "Reporting-Endpoints"
	// Report_To: Response header used to specify server endpoints where the
	// browser should send warning and error reports when using the Reporting API.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Report-To
	Report_To Header = "Report-To"
)

// Transfer coding
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#transfer_coding
const (
	// Transfer_Encoding: Specifies the form of encoding used to safely transfer
	// the resource to the user.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Transfer-Encoding
	Transfer_Encoding Header = "Transfer-Encoding"
	// TE: Specifies the transfer encodings the user agent is willing to accept.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/TE
	TE Header = "TE"
	// Trailer: Allows the sender to include additional fields at the end of
	// chunked message.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Trailer
	Trailer Header = "Trailer"
)

// WebSockets
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#websockets
const (
	// Sec_WebSocket_Accept: Response header that indicates that the server is
	// willing to upgrade to a WebSocket connection.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Sec-WebSocket-Accept
	Sec_WebSocket_Accept Header = "Sec-WebSocket-Accept"
	// Sec_WebSocket_Extensions: In requests, this header indicates the WebSocket
	// extensions supported by the client in preferred order. In responses, it indicates
	// the extension selected by the server from the client's preferences.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Sec-WebSocket-Extensions
	Sec_WebSocket_Extensions Header = "Sec-WebSocket-Extensions"
	// Sec_WebSocket_Key: Request header containing a key that verifies that the
	// client explicitly intends to open a WebSocket.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Sec-WebSocket-Key
	Sec_WebSocket_Key Header = "Sec-WebSocket-Key"
	// Sec_WebSocket_Protocol: In requests, this header indicates the sub-protocols
	// supported by the client in preferred order. In responses, it indicates the the
	// sub-protocol selected by the server from the client's preferences.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Sec-WebSocket-Protocol
	Sec_WebSocket_Protocol Header = "Sec-WebSocket-Protocol"
	// Sec_WebSocket_Version: In requests, this header indicates the version of the
	// WebSocket protocol used by the client. In responses, it is sent only if the
	// requested protocol version is not supported by the server, and lists the versions
	// that the server supports.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Sec-WebSocket-Version
	Sec_WebSocket_Version Header = "Sec-WebSocket-Version"
)

// Other
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#other
const (
	// Alt_Svc: Used to list alternate ways to reach this service.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Alt-Svc
	Alt_Svc Header = "Alt-Svc"
	// Alt_Used: Used to identify the alternative service in use.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Alt-Used
	Alt_Used Header = "Alt-Used"
	// Date: Contains the date and time at which the message was originated.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Date
	Date Header = "Date"
	// Link: This entity-header field provides a means for serializing one or more
	// links in HTTP headers. It is semantically equivalent to the HTML <link> element.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Link
	Link Header = "Link"
	// Retry_After: Indicates how long the user agent should wait before making a
	// follow-up request.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Retry-After
	Retry_After Header = "Retry-After"
	// Server_Timing: Communicates one or more metrics and descriptions for the
	// given request-response cycle.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Server-Timing
	Server_Timing Header = "Server-Timing"
	// Service_Worker_Allowed: Used to remove the path restriction by including
	// this header in the response of the Service Worker script.
	// https://developer.mozilla.org#service-worker-allowed
	Service_Worker_Allowed Header = "Service-Worker-Allowed"
	// SourceMap: Links generated code to a source map.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/SourceMap
	SourceMap Header = "SourceMap"
	// Upgrade: This HTTP/1.1 (only) header can be used to upgrade an already
	// established client/server connection to a different protocol (over the same
	// transport protocol). For example, it can be used by a client to upgrade a
	// connection from HTTP 1.1 to HTTP 2.0, or an HTTP or HTTPS connection into a
	// WebSocket.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Upgrade
	Upgrade Header = "Upgrade"
	// Priority: Provides a hint from about the priority of a particular resource
	// request on a particular connection. The value can be sent in a request to
	// indicate the client priority, or in a response if the server chooses to
	// reprioritize the request.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Priority
	Priority Header = "Priority"
)

// Non-standard headers
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#non-standard_headers
const (
	// X_Forwarded_For: Identifies the originating IP addresses of a client
	// connecting to a web server through an HTTP proxy or a load balancer.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Forwarded-For
	X_Forwarded_For Header = "X-Forwarded-For"
	// X_Forwarded_Host: Identifies the original host requested that a client used
	// to connect to your proxy or load balancer.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Forwarded-Host
	X_Forwarded_Host Header = "X-Forwarded-Host"
	// X_Forwarded_Proto: Identifies the protocol (HTTP or HTTPS) that a client
	// used to connect to your proxy or load balancer.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Forwarded-Proto
	X_Forwarded_Proto Header = "X-Forwarded-Proto"
	// X_DNS_Prefetch_Control: Controls DNS prefetching, a feature by which
	// browsers proactively perform domain name resolution on both links that the user
	// may choose to follow as well as URLs for items referenced by the document,
	// including images, CSS, JavaScript, and so forth.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-DNS-Prefetch-Control
	X_DNS_Prefetch_Control Header = "X-DNS-Prefetch-Control"
	// X_Robots_Tag: The X-Robots-Tag HTTP header is used to indicate how a web
	// page is to be indexed within public search engine results. The header is
	// effectively equivalent to <meta name="robots" content="…">.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#X-Robots-Tag
	X_Robots_Tag Header = "X-Robots-Tag"
)

// Deprecated headers
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#deprecated_headers
const (
	// Pragma: Implementation-specific header that may have various effects
	// anywhere along the request-response chain. Used for backwards compatibility with
	// HTTP/1.0 caches where the Cache-Control header is not yet present.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Pragma
	Pragma Header = "Pragma"
	// Warning: General warning information about possible problems.
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Warning
	Warning Header = "Warning"
)
