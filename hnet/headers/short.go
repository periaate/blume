package headers

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
