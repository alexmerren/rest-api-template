package server

import (
	"fmt"
	"golang-api-template/internal/config"
	"golang-api-template/internal/handler"
	"golang-api-template/internal/logger"
	"golang-api-template/internal/store"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	Logger *logger.ZapLogger
	Server *http.Server
	Port   int
}

func NewServer(logger *logger.ZapLogger, store *store.Store) (*Server, error) {
	config := config.GetConfig()
	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", config.Host, config.Port),
		Handler: newRouter(store),
	}

	return &Server{
		Logger: logger,
		Server: server,
		Port:   config.Port,
	}, nil
}

func newRouter(store *store.Store) *mux.Router {
	r := mux.NewRouter()
	h := handler.NewHandler(store)
	r.HandleFunc("/api/test/", handler.Test)
	r.HandleFunc("/api/create/", h.CreateContact)
	r.HandleFunc("/api/read/{id}", h.GetContact)
	r.HandleFunc("/api/update/", h.UpdateContact)
	r.HandleFunc("/api/delete/", h.DeleteContact)
	return r
}

func (s *Server) Run() {
	s.Logger.Info(fmt.Sprintf("Running API at %s", s.Server.Addr))
	if err := http.ListenAndServe(s.Server.Addr, s.Server.Handler); err != nil {
		s.Logger.Error("API Listen Error")
	}
}
