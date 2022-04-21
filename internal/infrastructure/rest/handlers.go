package rest

import (
	"context"
	"encoding/json"
	"net/http"
	"rest-api-template/internal/domain/entities"

	"github.com/gorilla/mux"
)

type CreateRequestBody struct {
	Contacts []*entities.Contact `json:"contacts"`
}

func (s *RESTServer) Create(w http.ResponseWriter, r *http.Request) {
	var requestBody CreateRequestBody
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		HandleError(w, r, err)
		return
	}

	err := s.usecases.CreateContacts(context.Background(), requestBody.Contacts)
	if err != nil {
		HandleError(w, r, err)
		return
	}
}

func (s *RESTServer) ReadOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	requestID := vars["id"]
	contact, err := s.usecases.GetContactByID(context.Background(), requestID)
	if err != nil {
		HandleError(w, r, err)
		return
	}
	err = json.NewEncoder(w).Encode(contact)
	if err != nil {
		s.logger.Info(err)
		HandleError(w, r, err)
		return
	}
}

func (s *RESTServer) ReadMany(w http.ResponseWriter, r *http.Request) {
	contacts, err := s.usecases.ListContacts(context.Background())
	if err != nil {
		HandleError(w, r, err)
		return
	}
	err = json.NewEncoder(w).Encode(contacts)
	if err != nil {
		HandleError(w, r, err)
		return
	}
}

type UpdateRequestBody struct {
	Contact *entities.Contact `json:"contact"`
}

func (s *RESTServer) Update(w http.ResponseWriter, r *http.Request) {
	var requestBody UpdateRequestBody
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		HandleError(w, r, err)
		return
	}

	vars := mux.Vars(r)
	requestID := vars["id"]
	contact, err := s.usecases.UpdateContactByID(context.Background(), requestID, requestBody.Contact)
	if err != nil {
		HandleError(w, r, err)
		return
	}
	err = json.NewEncoder(w).Encode(contact)
	if err != nil {
		HandleError(w, r, err)
		return
	}
}

func (s *RESTServer) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	requestID := vars["id"]
	contact, err := s.usecases.DeleteContactByID(context.Background(), requestID)
	if err != nil {
		HandleError(w, r, err)
		return
	}
	err = json.NewEncoder(w).Encode(contact)
	if err != nil {
		HandleError(w, r, err)
		return
	}
}
