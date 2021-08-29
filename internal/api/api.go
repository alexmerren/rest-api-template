package api

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Server struct {
	Port   int
	Router *mux.Router
}

func NewServer(port int) *Server {
	r := mux.NewRouter()
	server := &Server{
		Port:   port,
		Router: r,
	}
	// This is where we add the routes, and their handlers.
	r.HandleFunc("/api/v1/{id}", handler).Methods("GET")
	return server
}

func handler(w http.ResponseWriter, r *http.Request) {
	return
}
