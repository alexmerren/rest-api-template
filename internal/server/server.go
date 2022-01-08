package server

import (
	"context"
	"fmt"
	"golang-api-template/internal/config"
	"golang-api-template/internal/handler"
	"golang-api-template/internal/logger"
	"golang-api-template/internal/store"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	Logger logger.LoggerInterface
	Server *http.Server
	Port   int
}

// NewServer returns a server that can be run, with all the proper configurations
func NewServer(context context.Context, config config.ConfigInterface, logger logger.LoggerInterface, store *store.Store) (*Server, error) {
	host, err := config.GetString("Host")
	port, err := config.GetInt("Port")
	if err != nil {
		return nil, err
	}

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", host, port),
		Handler: newRouter(context, logger, store),
	}

	return &Server{
		Logger: logger,
		Server: server,
		Port:   port,
	}, nil
}

// newRouter creates a new router for the server to use
func newRouter(context context.Context, logger logger.LoggerInterface, store *store.Store) *mux.Router {
	r := mux.NewRouter()
	h := handler.NewHandler(context, logger, store)
	r.HandleFunc("/api/test/", handler.Test)
	r.HandleFunc("/api/create/", h.CreateContact)
	r.HandleFunc("/api/read/", h.GetAllContacts)
	r.HandleFunc("/api/read/{id}/", h.GetContact)
	r.HandleFunc("/api/update/{id}/", h.UpdateContact)
	r.HandleFunc("/api/delete/{id}/", h.DeleteContact)
	return r
}

func (s *Server) Run() {
	s.Logger.Info(fmt.Sprintf("Running API at %s", s.Server.Addr))
	if err := http.ListenAndServe(s.Server.Addr, s.Server.Handler); err != nil {
		s.Logger.Error("API Listen Error")
	}
}
