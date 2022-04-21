package rest

import (
	"net/http"
)

func (s *RESTServer) HandleCreate(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (s *RESTServer) HandleRead(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (s *RESTServer) HandleUpdate(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (s *RESTServer) HandleDelete(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func HandleError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusTeapot)
}
