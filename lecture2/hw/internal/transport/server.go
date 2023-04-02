package transport

import (
	"fmt"
	"net/http"
	"time"
)

const (
	readTimeout  = 5 * time.Second
	writeTimeout = 5 * time.Second
)

type Server struct {
	httpServer *http.Server
}

func NewServer(port int) *Server {
	return &Server{httpServer: &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}}
}

func (s *Server) ListenAndServe() error {
	return s.httpServer.ListenAndServe()
}
