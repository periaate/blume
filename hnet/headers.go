package hnet

import "net/http"

type Header string

func (h Header) String() string                        { return string(h) }
func (h Header) Set(w http.ResponseWriter, val string) { w.Header().Set(string(h), val) }
func (h Header) Tuple(val string) [2]string            { return [2]string{string(h), val} }

// Authentication
const (
	// WWW_Authenticate: Defines the authentication method that should be used to access a resource.
	WWW_Authenticate Header = "WWW-Authenticate"

	// Authorization: Contains the credentials to authenticate a user-agent with a server.
	Authorization Header = "Authorization"

	// Proxy_Authenticate: Defines the authentication method that should be used to access a resource behind a proxy server.
	Proxy_Authenticate Header = "Proxy-Authenticate"

	// Proxy_Authorization: Contains the credentials to authenticate a user agent with a proxy server.
	Proxy_Authorization Header = "Proxy-Authorization"
)

// Caching
const (
	// Age: The time, in seconds, that the object has been in a proxy cache.
	Age Header = "Age"

	// Cache_Control: Directives for caching mechanisms in both requests and responses.
	Cache_Control Header = "Cache-Control"

	// Clear_Site_Data: Clears browsing data (e.g. cookies, storage, cache) associated with the requesting website.
	Clear_Site_Data Header = "Clear-Site-Data"

	// Expires: The date/time after which the response is considered stale.
	Expires Header = "Expires"

	// No_Vary_Search: Specifies a set of rules that define how a URL's query parameters will affect cache matching. These rules dictate whether the same URL with different URL parameters should be saved as separate browser cache entries.
	No_Vary_Search Header = "No-Vary-Search"
)

// Conditionals
const (
	// Last_Modified: The last modification date of the resource, used to compare several versions of the same resource. It is less accurate than ETag, but easier to calculate in some environments. Conditional requests using If-Modified-Since and If-Unmodified-Since use this value to change the behavior of the request.
	Last_Modified Header = "Last-Modified"

	// ETag: A unique string identifying the version of the resource. Conditional requests using If-Match and If-None-Match use this value to change the behavior of the request.
	ETag Header = "ETag"

	// If_Match: Makes the request conditional, and applies the method only if the stored resource matches one of the given ETags.
	If_Match Header = "If-Match"

	// If_None_Match: Makes the request conditional, and applies the method only if the stored resource doesn't match any of the given ETags. This is used to update caches (for safe requests), or to prevent uploading a new resource when one already exists.
	If_None_Match Header = "If-None-Match"

	// If_Modified_Since: Makes the request conditional, and expects the resource to be transmitted only if it has been modified after the given date. This is used to transmit data only when the cache is out of date.
	If_Modified_Since Header = "If-Modified-Since"

	// If_Unmodified_Since: Makes the request conditional, and expects the resource to be transmitted only if it has not been modified after the given date. This ensures the coherence of a new fragment of a specific range with previous ones, or to implement an optimistic concurrency control system when modifying existing documents.
	If_Unmodified_Since Header = "If-Unmodified-Since"

	// Vary: Determines how to match request headers to decide whether a cached response can be used rather than requesting a fresh one from the origin server.
	Vary Header = "Vary"
)

// Connection management
const (
	// Connection: Controls whether the network connection stays open after the current transaction finishes.
	Connection Header = "Connection"

	// Keep_Alive: Controls how long a persistent connection should stay open.
	Keep_Alive Header = "Keep-Alive"
)

// Content negotiation
const (
	// Accept: Informs the server about the types of data that can be sent back.
	Accept Header = "Accept"

	// Accept_Encoding: The encoding algorithm, usually a compression algorithm, that can be used on the resource sent back.
	Accept_Encoding Header = "Accept-Encoding"

	// Accept_Language: Informs the server about the human language the server is expected to send back. This is a hint and is not necessarily under the full control of the user: the server should always pay attention not to override an explicit user choice (like selecting a language from a dropdown).
	Accept_Language Header = "Accept-Language"

	// Accept_Patch: A request content negotiation response header that advertises which media type the server is able to understand in a PATCH request.
	Accept_Patch Header = "Accept-Patch"

	// Accept_Post: A request content negotiation response header that advertises which media type the server is able to understand in a POST request.
	Accept_Post Header = "Accept-Post"
)

// Controls
const (
	// Expect: Indicates expectations that need to be fulfilled by the server to properly handle the request.
	Expect Header = "Expect"

	// Max_Forwards: When using TRACE, indicates the maximum number of hops the request can do before being reflected to the sender.
	Max_Forwards Header = "Max-Forwards"
)

// Cookies
const (
	// Cookie: Contains stored HTTP cookies previously sent by the server with the Set-Cookie header.
	Cookie Header = "Cookie"

	// Set_Cookie: Send cookies from the server to the user-agent.
	Set_Cookie Header = "Set-Cookie"
)

// CORS
const (
	// Access_Control_Allow_Credentials: Indicates whether the response to the request can be exposed when the credentials flag is true.
	Access_Control_Allow_Credentials Header = "Access-Control-Allow-Credentials"

	// Access_Control_Allow_Headers: Used in response to a preflight request to indicate which HTTP headers can be used when making the actual request.
	Access_Control_Allow_Headers Header = "Access-Control-Allow-Headers"

	// Access_Control_Allow_Methods: Specifies the methods allowed when accessing the resource in response to a preflight request.
	Access_Control_Allow_Methods Header = "Access-Control-Allow-Methods"

	// Access_Control_Allow_Origin: Indicates whether the response can be shared.
	Access_Control_Allow_Origin Header = "Access-Control-Allow-Origin"

	// Access_Control_Expose_Headers: Indicates which headers can be exposed as part of the response by listing their names.
	Access_Control_Expose_Headers Header = "Access-Control-Expose-Headers"

	// Access_Control_Max_Age: Indicates how long the results of a preflight request can be cached.
	Access_Control_Max_Age Header = "Access-Control-Max-Age"

	// Access_Control_Request_Headers: Used when issuing a preflight request to let the server know which HTTP headers will be used when the actual request is made.
	Access_Control_Request_Headers Header = "Access-Control-Request-Headers"

	// Access_Control_Request_Method: Used when issuing a preflight request to let the server know which HTTP method will be used when the actual request is made.
	Access_Control_Request_Method Header = "Access-Control-Request-Method"

	// Origin: Indicates where a fetch originates from.
	Origin Header = "Origin"

	// Timing_Allow_Origin: Specifies origins that are allowed to see values of attributes retrieved via features of the Resource Timing API, which would otherwise be reported as zero due to cross-origin restrictions.
	Timing_Allow_Origin Header = "Timing-Allow-Origin"
)

// Downloads
const (
	// Content_Disposition: Indicates if the resource transmitted should be displayed inline (default behavior without the header), or if it should be handled like a download and the browser should present a "Save As" dialog.
	Content_Disposition Header = "Content-Disposition"
)

// Integrity digests
const (
	// Content_Digest: Provides a digest of the stream of octets framed in an HTTP message (the message content) dependent on Content-Encoding and Content-Range.
	Content_Digest Header = "Content-Digest"

	// Repr_Digest: Provides a digest of the selected representation of the target resource before transmission. Unlike the Content-Digest, the digest does not consider Content-Encoding or Content-Range.
	Repr_Digest Header = "Repr-Digest"

	// Want_Content_Digest: States the wish for a Content-Digest header. It is the Content- analogue of Want-Repr-Digest.
	Want_Content_Digest Header = "Want-Content-Digest"

	// Want_Repr_Digest: States the wish for a Repr-Digest header. It is the Repr- analogue of Want-Content-Digest.
	Want_Repr_Digest Header = "Want-Repr-Digest"
)

// Message body information
const (
	// Content_Length: The size of the resource, in decimal number of bytes.
	Content_Length Header = "Content-Length"

	// Content_Type: Indicates the media type of the resource.
	Content_Type Header = "Content-Type"

	// Content_Encoding: Used to specify the compression algorithm.
	Content_Encoding Header = "Content-Encoding"

	// Content_Language: Describes the human language(s) intended for the audience, so that it allows a user to differentiate according to the users' own preferred language.
	Content_Language Header = "Content-Language"

	// Content_Location: Indicates an alternate location for the returned data.
	Content_Location Header = "Content-Location"
)

// Proxies
const (
	// Forwarded: Contains information from the client-facing side of proxy servers that is altered or lost when a proxy is involved in the path of the request.
	Forwarded Header = "Forwarded"

	// Via: Added by proxies, both forward and reverse proxies, and can appear in the request headers and the response headers.
	Via Header = "Via"
)

// Range requests
const (
	// Accept_Ranges: Indicates if the server supports range requests, and if so in which unit the range can be expressed.
	Accept_Ranges Header = "Accept-Ranges"

	// Range: Indicates the part of a document that the server should return.
	Range Header = "Range"

	// If_Range: Creates a conditional range request that is only fulfilled if the given etag or date matches the remote resource. Used to prevent downloading two ranges from incompatible version of the resource.
	If_Range Header = "If-Range"

	// Content_Range: Indicates where in a full body message a partial message belongs.
	Content_Range Header = "Content-Range"
)

// Redirects
const (
	// Location: Indicates the URL to redirect a page to.
	Location Header = "Location"

	// Refresh: Directs the browser to reload the page or redirect to another. Takes the same value as the meta element with http-equiv="refresh".
	Refresh Header = "Refresh"
)

// Request context
const (
	// From: Contains an Internet email address for a human user who controls the requesting user agent.
	From Header = "From"

	// Host: Specifies the domain name of the server (for virtual hosting), and (optionally) the TCP port number on which the server is listening.
	Host Header = "Host"

	// Referer: The address of the previous web page from which a link to the currently requested page was followed.
	Referer Header = "Referer"

	// Referrer_Policy: Governs which referrer information sent in the Referer header should be included with requests made.
	Referrer_Policy Header = "Referrer-Policy"

	// User_Agent: Contains a characteristic string that allows the network protocol peers to identify the application type, operating system, software vendor or software version of the requesting software user agent.
	User_Agent Header = "User-Agent"
)

// Response context
const (
	// Allow: Lists the set of HTTP request methods supported by a resource.
	Allow Header = "Allow"

	// Server: Contains information about the software used by the origin server to handle the request.
	Server Header = "Server"
)

// Security
const (
	// Cross_Origin_Embedder_Policy: Allows a server to declare an embedder policy for a given document.
	Cross_Origin_Embedder_Policy Header = "Cross-Origin-Embedder-Policy"

	// Cross_Origin_Opener_Policy: Prevents other domains from opening/controlling a window.
	Cross_Origin_Opener_Policy Header = "Cross-Origin-Opener-Policy"

	// Cross_Origin_Resource_Policy: Prevents other domains from reading the response of the resources to which this header is applied. See also CORP explainer article.
	Cross_Origin_Resource_Policy Header = "Cross-Origin-Resource-Policy"

	// Content_Security_Policy: Controls resources the user agent is allowed to load for a given page.
	Content_Security_Policy Header = "Content-Security-Policy"

	// Content_Security_Policy_Report_Only: Allows web developers to experiment with policies by monitoring, but not enforcing, their effects. These violation reports consist of JSON documents sent via an HTTP POST request to the specified URI.
	Content_Security_Policy_Report_Only Header = "Content-Security-Policy-Report-Only"

	// Expect_CT: Lets sites opt in to reporting and enforcement of Certificate Transparency to detect use of misissued certificates for that site.
	Expect_CT Header = "Expect-CT"

	// Permissions_Policy: Provides a mechanism to allow and deny the use of browser features in a website's own frame, and in <iframe>s that it embeds.
	Permissions_Policy Header = "Permissions-Policy"

	// Reporting_Endpoints: Response header that allows website owners to specify one or more endpoints used to receive errors such as CSP violation reports, Cross-Origin-Opener-Policy reports, or other generic violations. Or: Response header used to specify server endpoints where the browser should send warning and error reports when using the Reporting API.
	Reporting_Endpoints Header = "Reporting-Endpoints"

	// Strict_Transport_Security: Force communication using HTTPS instead of HTTP.
	Strict_Transport_Security Header = "Strict-Transport-Security"

	// Upgrade_Insecure_Requests: Sends a signal to the server expressing the client's preference for an encrypted and authenticated response, and that it can successfully handle the upgrade-insecure-requests directive.
	Upgrade_Insecure_Requests Header = "Upgrade-Insecure-Requests"

	// X_Content_Type_Options: Disables MIME sniffing and forces browser to use the type given in Content-Type.
	X_Content_Type_Options Header = "X-Content-Type-Options"

	// X_Frame_Options: Indicates whether a browser should be allowed to render a page in a <frame>, <iframe>, <embed> or <object>.
	X_Frame_Options Header = "X-Frame-Options"

	// X_Permitted_Cross_Domain_Policies: Specifies if a cross-domain policy file (crossdomain.xml) is allowed. The file may define a policy to grant clients, such as Adobe's Flash Player (now obsolete), Adobe Acrobat, Microsoft Silverlight (now obsolete), or Apache Flex, permission to handle data across domains that would otherwise be restricted due to the Same-Origin Policy. See the Cross-domain Policy File Specification for more information.
	X_Permitted_Cross_Domain_Policies Header = "X-Permitted-Cross-Domain-Policies"

	// X_Powered_By: May be set by hosting environments or other frameworks and contains information about them while not providing any usefulness to the application or its visitors. Unset this header to avoid exposing potential vulnerabilities.
	X_Powered_By Header = "X-Powered-By"

	// X_XSS_Protection: Enables cross-site scripting filtering.
	X_XSS_Protection Header = "X-XSS-Protection"
)

// Server-sent events
const (
	// Report_To: Response header used to specify server endpoints where the browser should send warning and error reports when using the Reporting API.
	Report_To Header = "Report-To"
)

// Transfer coding
const (
	// Transfer_Encoding: Specifies the form of encoding used to safely transfer the resource to the user.
	Transfer_Encoding Header = "Transfer-Encoding"

	// TE: Specifies the transfer encodings the user agent is willing to accept.
	TE Header = "TE"

	// Trailer: Allows the sender to include additional fields at the end of chunked message.
	Trailer Header = "Trailer"
)

// WebSockets
const (
	// Sec_WebSocket_Accept: Response header that indicates that the server is willing to upgrade to a WebSocket connection.
	Sec_WebSocket_Accept Header = "Sec-WebSocket-Accept"

	// Sec_WebSocket_Extensions: In requests, this header indicates the WebSocket extensions supported by the client in preferred order. In responses, it indicates the extension selected by the server from the client's preferences.
	Sec_WebSocket_Extensions Header = "Sec-WebSocket-Extensions"

	// Sec_WebSocket_Key: Request header containing a key that verifies that the client explicitly intends to open a WebSocket.
	Sec_WebSocket_Key Header = "Sec-WebSocket-Key"

	// Sec_WebSocket_Protocol: In requests, this header indicates the sub-protocols supported by the client in preferred order. In responses, it indicates the the sub-protocol selected by the server from the client's preferences.
	Sec_WebSocket_Protocol Header = "Sec-WebSocket-Protocol"

	// Sec_WebSocket_Version: In requests, this header indicates the version of the WebSocket protocol used by the client. In responses, it is sent only if the requested protocol version is not supported by the server, and lists the versions that the server supports.
	Sec_WebSocket_Version Header = "Sec-WebSocket-Version"
)

// Other
const (
	// Alt_Svc: Used to list alternate ways to reach this service.
	Alt_Svc Header = "Alt-Svc"

	// Alt_Used: Used to identify the alternative service in use.
	Alt_Used Header = "Alt-Used"

	// Date: Contains the date and time at which the message was originated.
	Date Header = "Date"

	// Link: This entity-header field provides a means for serializing one or more links in HTTP headers. It is semantically equivalent to the HTML <link> element.
	Link Header = "Link"

	// Retry_After: Indicates how long the user agent should wait before making a follow-up request.
	Retry_After Header = "Retry-After"

	// Server_Timing: Communicates one or more metrics and descriptions for the given request-response cycle.
	Server_Timing Header = "Server-Timing"

	// Service_Worker_Allowed: Used to remove the path restriction by including this header in the response of the Service Worker script.
	Service_Worker_Allowed Header = "Service-Worker-Allowed"

	// SourceMap: Links generated code to a source map.
	SourceMap Header = "SourceMap"

	// Upgrade: This HTTP/1.1 (only) header can be used to upgrade an already established client/server connection to a different protocol (over the same transport protocol). For example, it can be used by a client to upgrade a connection from HTTP 1.1 to HTTP 2.0, or an HTTP or HTTPS connection into a WebSocket.
	Upgrade Header = "Upgrade"

	// Priority: Provides a hint from about the priority of a particular resource request on a particular connection. The value can be sent in a request to indicate the client priority, or in a response if the server chooses to reprioritize the request.
	Priority Header = "Priority"
)

// Non-standard headers
const (
	// X_Forwarded_For: Identifies the originating IP addresses of a client connecting to a web server through an HTTP proxy or a load balancer.
	X_Forwarded_For Header = "X-Forwarded-For"

	// X_Forwarded_Host: Identifies the original host requested that a client used to connect to your proxy or load balancer.
	X_Forwarded_Host Header = "X-Forwarded-Host"

	// X_Forwarded_Proto: Identifies the protocol (HTTP or HTTPS) that a client used to connect to your proxy or load balancer.
	X_Forwarded_Proto Header = "X-Forwarded-Proto"

	// X_DNS_Prefetch_Control: Controls DNS prefetching, a feature by which browsers proactively perform domain name resolution on both links that the user may choose to follow as well as URLs for items referenced by the document, including images, CSS, JavaScript, and so forth.
	X_DNS_Prefetch_Control Header = "X-DNS-Prefetch-Control"

	// X_Robots_Tag: The X-Robots-Tag HTTP header is used to indicate how a web page is to be indexed within public search engine results. The header is effectively equivalent to <meta name="robots" content="…">.
	X_Robots_Tag Header = "X-Robots-Tag"
)

// Deprecated headers
const (
	// Pragma: Implementation-specific header that may have various effects anywhere along the request-response chain. Used for backwards compatibility with HTTP/1.0 caches where the Cache-Control header is not yet present.
	Pragma Header = "Pragma"

	// Warning: General warning information about possible problems.
	Warning Header = "Warning"
)

// Aliases and short hands

// CORS headers
const (
	// Allow_Origin: Indicates whether the response can be shared with requesting code from the given origin.
	Allow_Origin Header = "Access-Control-Allow-Origin"
	// Allow_Methods: Specifies the methods allowed when accessing the resource in response to a preflight request.
	Allow_Methods Header = "Access-Control-Allow-Methods"
	// Allow_Headers: Used in response to a preflight request to indicate which HTTP headers can be used when making the actual request.
	Allow_Headers Header = "Access-Control-Allow-Headers"
	// Allow_Credentials: Indicates whether the response to the request can be exposed when the credentials flag is true.
	Allow_Credentials Header = "Access-Control-Allow-Credentials"
	// Expose_Headers: Indicates which headers can be exposed as part of the response by listing their names.
	Expose_Headers Header = "Access-Control-Expose-Headers"
	// Max_Age: Indicates how long the results of a preflight request can be cached.
	Max_Age Header = "Access-Control-Max-Age"
	// Request_Headers: Used when issuing a preflight request to let the server know which HTTP headers will be used when the actual request is made.
	Request_Headers Header = "Access-Control-Request-Headers"
	// Request_Method: Used when issuing a preflight request to let the server know which HTTP method will be used when the actual request is made.
	Request_Method Header = "Access-Control-Request-Method"
)
