package rest

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (s *RESTServer) mapRoutes() error {
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/health/", s.Health).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/create/", s.Create).Methods(http.MethodPost)
	router.HandleFunc("/api/v1/read/", s.ReadMany).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/read/{id}/", s.ReadOne).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/update/{id}/", s.Update).Methods(http.MethodPut)
	router.HandleFunc("/api/v1/delete/{id}/", s.Delete).Methods(http.MethodPost)
	router.Use(s.loggingMiddleware)
	router.Use(s.formatMiddleware)
	s.httpServer.Handler = router
	return nil
}
