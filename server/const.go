package server

import "time"

const (
	version                = "0.0.1"     // Application and server version.
	DefaultHostName        = "localhost" // The hostname of the server.
	DefaultPort            = 8080        // Port to receive requests: see IANA Port Numbers.
	DefaultProfPort        = 0           // Profiler port to receive requests.*
	DefaultMaxProcs        = 0           // Maximum number of computer processors to utilize.*
	DefaultPollingInterval = 300         // Polling interval in seconds to check artifactory (5 min).

	// * zeros = no change or no limitations or not enabled.

	// http: routes.
	httpRouteV1Health = "/v1.0/health"
	httpRouteV1Hello  = "/v1.0/"

	// Connections.
	TCPReadTimeout  = 10 * time.Second
	TCPWriteTimeout = 10 * time.Second

	httpGet    = "GET"
	httpPost   = "POST"
	httpPut    = "PUT"
	httpDelete = "DELETE"
	httpHead   = "HEAD"
	httpTrace  = "TRACE"
	httpPatch  = "PATCH"

	// Error messages.
	InvalidMediaType     = "Invalid Content-Type or Accept header value."
	InvalidMethod        = "Invalid Method for this route."
	InvalidBody          = "Invalid body of text in request."
	InvalidJSONText      = "Invalid JSON format in text of body in request."
	InvalidJSONAttribute = "Invalid - 'text' attribute in JSON not found."
	InvalidAuthorization = "Invalid authorization."
)
