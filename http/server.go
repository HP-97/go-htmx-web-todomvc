package http

import (
	"fmt"
	"net/http"
)

// Server represents an HTTP server.
type Server struct {
	server *http.Server
	Addr string
	router *http.ServeMux
}

// NewServer returns a new instance of Server.
func NewServer() *Server {
	// Create a new server that wraps the net/http server.
	s := &Server {
		server: &http.Server{},
		router: http.NewServeMux(),
	}

	s.router.HandleFunc("GET /", s.handleIndex)

	return s
}

func (s *Server) serveHTTP(w http.ResponseWriter, r *http.Request) {
	// If required, we can write overriding code

	// Delegate remaining HTTP handling to the gorilla router.
	s.router.ServeHTTP(w, r)
}

func (s *Server) handleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, this is the index page!")
}
