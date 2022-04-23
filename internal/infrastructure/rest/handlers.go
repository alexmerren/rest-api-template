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
		HandleError(w, r, entities.NewBadRequestError("could not decode JSON from Body", err))
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
		HandleError(w, r, entities.NewInternalError("could not encode to JSON", err))
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
		HandleError(w, r, entities.NewInternalError("could not encode to JSON", err))
		return
	}
}

func (s *RESTServer) Update(w http.ResponseWriter, r *http.Request) {
	var newContact *entities.Contact
	if err := json.NewDecoder(r.Body).Decode(&newContact); err != nil {
		HandleError(w, r, entities.NewBadRequestError("could not decode JSON from Body", err))
		return
	}

	vars := mux.Vars(r)
	requestID := vars["id"]
	contact, err := s.usecases.UpdateContactByID(context.Background(), requestID, newContact)
	if err != nil {
		HandleError(w, r, err)
		return
	}
	err = json.NewEncoder(w).Encode(contact)
	if err != nil {
		HandleError(w, r, entities.NewInternalError("could not encode to JSON", err))
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
		HandleError(w, r, entities.NewInternalError("could not encode to JSON", err))
		return
	}
}

func (s *RESTServer) Health(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte(`{"message":"healthy!"}`))
	if err != nil {
		HandleError(w, r, entities.NewInternalError("could not write to response", err))
		return
	}
}
