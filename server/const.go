package server

import "time"

const (
	DefaultHostName = "0.0.0.0" // The hostname of the server.
	DefaultPort     = 8080      // Port to receive requests: see IANA Port Numbers.

	// http: routes.
	httpRouteV1Health = "/v1.0/health"
	httpRouteV1Hello  = "/v1.0/hello"

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
)
