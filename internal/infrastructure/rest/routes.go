package rest

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (s *RESTServer) mapRoutes() error {
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/create/", s.HandleCreate).Methods(http.MethodPost)
	router.HandleFunc("/api/v1/read/", s.HandleRead).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/update/", s.HandleUpdate).Methods(http.MethodPut)
	router.HandleFunc("/api/v1/delete/", s.HandleDelete).Methods(http.MethodPost)
	router.Use(s.loggingMiddleware)
	s.httpServer.Handler = router
	return nil
}
