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
	Logger logger.Logger
	Server *http.Server
	Port   int
}

// NewServer returns a server that can be run, with all the proper configurations
func ProvideServer(
	context context.Context,
	config config.Config,
	logger logger.Logger,
	datastore *datastore.Datastore,
	router *mux.Router,
) (*Server, error) {
	host, hostErr := config.GetString("Host")
	if hostErr != nil {
		return nil, hostErr
	}

	port, portErr := config.GetInt("Port")
	if portErr != nil {
		return nil, portErr
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
