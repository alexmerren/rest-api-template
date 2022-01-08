package handler

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"golang-api-template/internal/logger"
	"golang-api-template/internal/store"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	Context context.Context
	db      *store.Store
	Logger  logger.LoggerInterface
}

// NewHandler creates a new handler used to handle incoming requests
func NewHandler(context context.Context, logger logger.LoggerInterface, store *store.Store) *Handler {
	return &Handler{
		Context: context,
		db:      store,
		Logger:  logger,
	}
}

// Test is responsible for /api/test/
func Test(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// CreateContact is responsible for /api/create/, and taking in parameters as request body
func (h *Handler) CreateContact(w http.ResponseWriter, r *http.Request) {
	newContact := &store.Contact{}
	err := json.NewDecoder(r.Body).Decode(&newContact)
	if err != nil {
		h.errorResponse(err, http.StatusBadRequest)
		return
	}

	err = h.db.InsertContact(h.Context, newContact)
	if err != nil {
		h.errorResponse(err, http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(newContact)
	if err != nil {
		h.errorResponse(err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	h.Logger.Debug(fmt.Sprintf("%d\t%s\t%s", http.StatusOK, r.Method, r.URL.Path))
}

// GetAllContacts is responsible for /api/read/{id}, returning a single contact
func (h *Handler) GetContact(w http.ResponseWriter, r *http.Request) {
	contact, err := h.db.GetContact(h.Context, mux.Vars(r)["id"])
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			h.errorResponse(err, http.StatusNotFound)
		default:
			h.errorResponse(err, http.StatusInternalServerError)
		}
		return
	}

	err = json.NewEncoder(w).Encode(contact)
	if err != nil {
		h.errorResponse(err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	h.Logger.Debug(fmt.Sprintf("%d\t%s\t%s", http.StatusOK, r.Method, r.URL.Path))
}

// GetContact is responsible for /api/read/, returning all contacts
func (h *Handler) GetAllContacts(w http.ResponseWriter, r *http.Request) {
	contacts, err := h.db.GetAllContacts(h.Context)
	if err != nil {
		h.errorResponse(err, http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(contacts)
	if err != nil {
		h.errorResponse(err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	h.Logger.Debug(fmt.Sprintf("%d\t%s\t%s", http.StatusOK, r.Method, r.URL.Path))
}

// UpdateContact is responsible for /api/update/{id}, updating a contact with the request body
func (h *Handler) UpdateContact(w http.ResponseWriter, r *http.Request) {
	contact, err := h.db.GetContact(h.Context, mux.Vars(r)["id"])
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			h.errorResponse(err, http.StatusNotFound)
		default:
			h.errorResponse(err, http.StatusInternalServerError)
		}
		return
	}

	err = json.NewDecoder(r.Body).Decode(&contact)
	if err != nil {
		h.errorResponse(err, http.StatusBadRequest)
		return
	}

	err = h.db.UpdateContact(h.Context, contact)
	if err != nil {
		h.errorResponse(err, http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(contact)
	if err != nil {
		h.errorResponse(err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	h.Logger.Debug(fmt.Sprintf("%d\t%s\t%s", http.StatusOK, r.Method, r.URL.Path))
}

// DeleteContact is responsible for /api/delete/{id}, deleting a contact with the specific id
func (h *Handler) DeleteContact(w http.ResponseWriter, r *http.Request) {
	contact, err := h.db.GetContact(h.Context, mux.Vars(r)["id"])
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			h.errorResponse(err, http.StatusNotFound)
		default:
			h.errorResponse(err, http.StatusInternalServerError)
		}
		return
	}

	if err := h.db.DeleteContact(h.Context, mux.Vars(r)["id"]); err != nil {
		h.errorResponse(err, http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(contact)
	if err != nil {
		h.errorResponse(err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	h.Logger.Debug(fmt.Sprintf("%d\t%s\t%s", http.StatusOK, r.Method, r.URL.Path))
}

func (h *Handler) errorResponse(err error, status int) {
	h.Logger.Error(err)
	w.WriteHeader(status)
}
