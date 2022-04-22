package rest

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (s *RESTServer) mapRoutes() error {
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/health/", s.Health).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/contacts/create/", s.Create).Methods(http.MethodPost)
	router.HandleFunc("/api/v1/contacts/read/", s.ReadMany).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/contacts/read/{id}/", s.ReadOne).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/contacts/update/{id}/", s.Update).Methods(http.MethodPut)
	router.HandleFunc("/api/v1/contacts/delete/{id}/", s.Delete).Methods(http.MethodPost)
	router.Use(s.loggingMiddleware)
	router.Use(s.formatMiddleware)
	s.httpServer.Handler = router
	return nil
}
