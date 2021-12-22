package server

import (
	"fmt"
	"golang-api-template/internal/config"
	"golang-api-template/internal/handler"
	"golang-api-template/internal/logger"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	Logger *logger.ZapLogger
	Server *http.Server
	Port   int
}

func NewServer(config *config.Configuration, logger *logger.ZapLogger) (*Server, error) {
	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", config.Host, config.Port),
		Handler: newRouter(),
	}

	return &Server{
		Logger: logger,
		Server: server,
		Port:   config.Port,
	}, nil
}

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/test", handler.TestHandlerFunc)
	return r
}

func (s *Server) Run() {
	s.Logger.Info("Running API")
	if err := http.ListenAndServe(s.Server.Addr, s.Server.Handler); err != nil {
		s.Logger.Error("API Listen Error")
	}
}
