package server

import (
	"io/fs"
	"net/http"

	"github.com/dtchanpura/svelte-tailwind-golang/frontend"
)

func (s *Server) routes() error {
	// Frontend Code as http.FileSystem
	var ffs http.FileSystem = http.Dir("frontend/public")
	if !s.debug {
		if f, err := fs.Sub(frontend.Content, "public"); err == nil {
			ffs = http.FileSystem(http.FS(f))
		} else {
			return err
		}
	}

	// Depending upon s.debug we serve either the embedded part or static files
	// directly.
	s.mux.Handle("/", http.FileServer(ffs))

	s.mux.HandleFunc("/version", handleVersion())
	return nil
}
