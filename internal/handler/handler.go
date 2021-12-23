package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"golang-api-template/internal/logger"
	"golang-api-template/internal/store"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	db     *store.Store
	logger *logger.ZapLogger
}

func NewHandler(store *store.Store, logger *logger.ZapLogger) *Handler {
	return &Handler{
		db:     store,
		logger: logger,
	}
}

func Test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Api is working!")
}

func (h *Handler) CreateContact(w http.ResponseWriter, r *http.Request) {
	// Get the input and unmarshal into a struct
	var newContact store.Contact
	json.NewDecoder(r.Body).Decode(&newContact)

	// Store the struct in the database
	err := h.db.InsertContact(newContact)
	if err != nil {
		h.logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Respond with the stored struct
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newContact)
}

func (h *Handler) GetContact(w http.ResponseWriter, r *http.Request) {
	contact, err := h.db.GetContact(mux.Vars(r)["id"])
	if err != nil {
		h.logger.Error(err)
		switch err {
		case sql.ErrNoRows:
			w.WriteHeader(http.StatusNotFound)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(contact)
}

func (h *Handler) GetAllContacts(w http.ResponseWriter, r *http.Request) {
	contacts, err := h.db.GetAllContacts()
	if err != nil {
		h.logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(contacts)
}

//func (h *Handler) UpdateContact(w http.ResponseWriter, r *http.Request) {
//err := h.db.UpdateContact()
//}

func (h *Handler) DeleteContact(w http.ResponseWriter, r *http.Request) {
	value := mux.Vars(r)["id"]
	if value != "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := h.db.DeleteContact(value); err != nil {
		h.logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	contact := store.Contact{
		ID: value,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(contact.ID)
}
