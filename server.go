package social_network_for_programmers

import (
	"net/http"
	"social_network_for_programmers/internal/config"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(cfg config.HttpServer, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + cfg.Port,
		Handler:        handler,
		MaxHeaderBytes: cfg.MaxHeaderBytes,
		ReadTimeout:    cfg.ReadTimeout,
		WriteTimeout:   cfg.WriteTimeout,
	}

	return s.httpServer.ListenAndServe()
}
