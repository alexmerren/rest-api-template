package server

import (
	"context"
	"fmt"
	"golang-api-template/internal/config"
	"golang-api-template/internal/datastore"
	"golang-api-template/internal/logger"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	Logger logger.LoggerInterface
	Server *http.Server
	Port   int
}

// NewServer returns a server that can be run, with all the proper configurations
// nolint:ineffassign,staticcheck // This allows us to check if any of them have an error, and return that error
// https://go.dev/doc/effective_go#redeclaration
func ProvideServer(
	context context.Context,
	config config.ConfigInterface,
	logger logger.LoggerInterface,
	datastore *datastore.Datastore,
	router *mux.Router,
) (*Server, error) {
	host, err := config.GetString("Host")
	port, err := config.GetInt("Port")
	if err != nil {
		return nil, err
	}

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", host, port),
		Handler: router,
	}

	return &Server{
		Logger: logger,
		Server: server,
		Port:   port,
	}, nil
}

func (s *Server) Run() {
	s.Logger.Info(fmt.Sprintf("Running API at %s", s.Server.Addr))
	if err := http.ListenAndServe(s.Server.Addr, s.Server.Handler); err != nil {
		s.Logger.Error("API Listen Error")
	}
}
