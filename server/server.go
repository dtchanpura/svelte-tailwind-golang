package server

import (
	"log"
	"net/http"
)

// Server struct holds all server related details
type Server struct {
	mux   *http.ServeMux
	debug bool
}

// NewServer returns a new pointer to Server struct with debug and new ServeMux.
func NewServer(debug bool) *Server {
	return &Server{
		debug: debug,
		mux:   http.NewServeMux(),
	}
}

// RunServer for starting the listener at given address.
func (s *Server) RunServer(addr string) error {
	err := s.routes()
	if err != nil {
		return err
	}
	log.Printf("Starting Server on %s", addr)
	err = http.ListenAndServe(addr, s.mux)
	if err != nil {
		return err
	}
	return nil
}
