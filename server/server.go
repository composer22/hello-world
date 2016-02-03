// Package server implements a server for artifactory monitoring.
package server

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/composer22/hello-world/logger"
)

// Server is the main structure that represents a server instance.
type Server struct {
	mu sync.RWMutex // For locking access to server attributes.

	running bool           // Is the server running?
	srvr    *http.Server   // HTTP server.
	log     *logger.Logger // Log instance for recording error and other messages.
}

// New is a factory function that returns a new server instance.
func New() *Server {
	s := &Server{
		running: false,
		log:     logger.New(logger.UseDefault, false),
	}

	// Setup the routes and server.
	mux := http.NewServeMux()
	mux.HandleFunc(httpRouteV1Health, s.healthHandler)
	mux.HandleFunc(httpRouteV1Hello, s.helloHandler)
	s.srvr = &http.Server{
		Addr:         fmt.Sprintf("%s:%d", DefaultHostName, DefaultPort),
		Handler:      &Middleware{serv: s, handler: mux},
		ReadTimeout:  TCPReadTimeout,
		WriteTimeout: TCPWriteTimeout,
	}
	return s
}

// Start spins up the server to accept incoming requests.
func (s *Server) Start() error {
	if s.isRunning() {
		return errors.New("Server already started.")
	}

	s.log.Infof("Starting hello world\n")
	s.handleSignals()

	s.mu.Lock()
	s.running = true
	s.mu.Unlock()
	err := s.srvr.ListenAndServe()
	if err != nil {
		s.log.Emergencyf("Listen and Server Error: %s", err.Error())
	}

	// Done.
	s.mu.Lock()
	s.running = false
	s.mu.Unlock()
	return nil
}

// Shutdown takes down the server gracefully back to an initialize state.
func (s *Server) Shutdown() {
	if !s.isRunning() {
		return
	}
	s.log.Infof("BEGIN server service stop.")
	s.mu.Lock()
	s.srvr.SetKeepAlivesEnabled(false)
	s.running = false
	s.mu.Unlock()
	s.log.Infof("END server service stop.")
}

// handleSignals responds to operating system interrupts such as application kills.
func (s *Server) handleSignals() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			s.log.Infof("Server received signal: %v\n", sig)
			s.Shutdown()
			s.log.Infof("Server exiting.")
			os.Exit(0)
		}
	}()
}

// The following methods handle server routes.

// healthHandler handles a client "is the server alive?" request.
func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

// helloHandler returns "Hello World!".
func (s *Server) helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

// initResponseHeader sets up the common http response headers for the return of all json calls.
func (s *Server) initResponseHeader(w http.ResponseWriter) {
	h := w.Header()
	h.Add("Content-Type", "application/json;charset=utf-8")
	h.Add("Date", time.Now().UTC().Format(time.RFC1123Z))
	h.Add("X-Request-ID", createV4UUID())
}

// isRunning returns a boolean representing whether the server is running or not.
func (s *Server) isRunning() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.running
}

// requestLogEntry is a datastructure of a log entry for recording server access requests.
type requestLogEntry struct {
	Method        string      `json:"method"`
	URL           *url.URL    `json:"url"`
	Proto         string      `json:"proto"`
	Header        http.Header `json:"header"`
	Body          string      `json:"body"`
	ContentLength int64       `json:"contentLength"`
	Host          string      `json:"host"`
	RemoteAddr    string      `json:"remoteAddr"`
	RequestURI    string      `json:"requestURI"`
	Trailer       http.Header `json:"trailer"`
}

// LogRequest logs the http request information into the logger.
func (s *Server) LogRequest(r *http.Request) {
	var cl int64

	if r.ContentLength > 0 {
		cl = r.ContentLength
	}

	bd, err := ioutil.ReadAll(r.Body)
	if err != nil {
		bd = []byte("Could not parse body")
	}
	r.Body = ioutil.NopCloser(bytes.NewBuffer(bd)) // We need to set the body back after we read it.

	b, _ := json.Marshal(&requestLogEntry{
		Method:        r.Method,
		URL:           r.URL,
		Proto:         r.Proto,
		Header:        r.Header,
		Body:          string(bd),
		ContentLength: cl,
		Host:          r.Host,
		RemoteAddr:    r.RemoteAddr,
		RequestURI:    r.RequestURI,
		Trailer:       r.Trailer,
	})
	s.log.Infof(`{"request":%s}`, string(b))
}
