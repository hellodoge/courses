package server

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

const (
	DefaultPort    = 8080
	DefaultTimeout = 10 * time.Second
)

type Server struct {
	httpServer *http.Server
}

type Config struct {
	Port    uint16
	Timeout time.Duration
}

func InitServer(config Config, handler http.Handler) *Server {
	var s = new(Server)
	s.httpServer = &http.Server{
		Addr:         fmt.Sprintf(":%d", config.Port),
		Handler:      handler,
		ReadTimeout:  config.Timeout,
		WriteTimeout: config.Timeout,
	}
	return s
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
