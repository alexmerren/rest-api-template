package server

import (
	"fmt"
	"golang-api-template/pkg/config"
	"golang-api-template/pkg/handler"
	"golang-api-template/pkg/logger"
	"golang-api-template/pkg/store"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	Logger logger.LoggerInterface
	Server *http.Server
	Port   int
}

func NewServer(logger logger.LoggerInterface, store *store.Store) (*Server, error) {
	config := config.GetConfig()
	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", config.Host, config.Port),
		Handler: newRouter(logger, store),
	}

	return &Server{
		Logger: logger,
		Server: server,
		Port:   config.Port,
	}, nil
}

func newRouter(logger logger.LoggerInterface, store *store.Store) *mux.Router {
	r := mux.NewRouter()
	h := handler.NewHandler(logger, store)
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
