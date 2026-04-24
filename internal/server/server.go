package server

import (
	"log/slog"
	"net/http"
)

type Server struct {
	serv *http.Server
}

func New(router http.Handler, port string) *Server {
	return &Server{
		serv: &http.Server{
			Addr:    ":" + port,
			Handler: router,
		},
	}
}

func (s *Server) Run() error {
	slog.Info("server is started", "ADDR", s.serv.Addr)
	return s.serv.ListenAndServe()
}
