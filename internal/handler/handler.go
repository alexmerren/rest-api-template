package handler

import (
	"encoding/json"
	"fmt"
	"golang-api-template/internal/store"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	db *store.Store
}

func NewHandler(store *store.Store) *Handler {
	return &Handler{
		db: store,
	}
}

func Test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Api is working!")
}

func (h *Handler) CreateContact(w http.ResponseWriter, r *http.Request) {
	// Get the input and unmarshal into a struct
	var newContact store.Contact
	json.Unmarshal(r.Body, &newContact)

	// Store the struct in the database
	err := InsertContact(newContact)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Respond with the stored struct
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newContact)
}

func (h *Handler) GetContact(w http.ResponseWriter, r *http.Request) {
	if mux.Vars(r)["id"] != nil {
		contact, err := h.db.GetContact(mux.Vars(r)["id"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(contact)
		return
	}

	contacts, err := GetAllContacts()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(contacts)
	return
}

func (h *Handler) UpdateContact(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) DeleteContact(w http.ResponseWriter, r *http.Request) {
}
