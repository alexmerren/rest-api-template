package server

import (
	"fmt"
	"net/http"

	"golang-api-template/internal/controllers"

	"github.com/gorilla/mux"
)

type Server struct {
	Router *mux.Router
}

func NewServer() *Server {
	return &Server{
		Router: makeRoutes(),
	}
}

func makeRoutes() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", controllers.TestAPI).Methods("GET")
	r.HandleFunc("/api", controllers.TestAPI).Methods("GET")
	return r
}

func (s *Server) StartServer(host string, port int) error {
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), s.Router)
	if err != nil {
		return err
	}
	return nil
}
