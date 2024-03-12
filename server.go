package social_network_for_programmers

import (
	"net/http"
	"social_network_for_programmers/internal/config"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(cfg config.HttpServer, handlers http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + cfg.Port,
		Handler:        handlers,
		MaxHeaderBytes: cfg.MaxHeaderBytes,
		ReadTimeout:    cfg.ReadTimeout,
		WriteTimeout:   cfg.WriteTimeout,
	}

	return s.httpServer.ListenAndServe()
}
