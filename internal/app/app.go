package app

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Start(handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":8080",
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1 mb
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
