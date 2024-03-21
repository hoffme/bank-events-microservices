package http

import (
	"net/http"
	"sync"

	"github.com/hoffme/backend-transactions/internal/shared/logger"
)

type Config struct {
	Addr string `json:"addr"`
}

type Server struct {
	config Config
	server *http.Server
}

func New(config Config) *Server {
	result := &Server{config: config}
	result.server = &http.Server{Addr: config.Addr}
	return result
}

func (s *Server) SetHandler(handler http.Handler) {
	s.server.Handler = handler
}

func (s *Server) Run(wg *sync.WaitGroup) {
	defer wg.Done()

	logger.Info("server running on %s", s.config.Addr)

	err := s.server.ListenAndServe()
	if err != nil {
		logger.Warn("server %s stopped: %w", s.config.Addr, err)
	}
}

func (s *Server) Close() {
	err := s.server.Close()
	if err != nil {
		logger.Warn("server %s closed: %w", s.config.Addr, err)
	}
}
