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
	Logger *logger.ZapLogger
}

func NewHandler(store *store.Store, logger *logger.ZapLogger) *Handler {
	return &Handler{
		db:     store,
		Logger: logger,
	}
}

func Test(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) CreateContact(w http.ResponseWriter, r *http.Request) {
	// Get the input and unmarshal into a struct
	newContact := &store.Contact{}
	json.NewDecoder(r.Body).Decode(&newContact)

	// Store the struct in the database
	err := h.db.InsertContact(newContact)
	if err != nil {
		h.Logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Respond with the stored struct
	w.WriteHeader(http.StatusCreated)
	h.Logger.Info(fmt.Sprintf("%d\t%s\t%s", http.StatusCreated, r.Method, r.URL.Path))
	json.NewEncoder(w).Encode(newContact)
}

func (h *Handler) GetContact(w http.ResponseWriter, r *http.Request) {
	// Call the GetContact with the id taken from the url
	contact, err := h.db.GetContact(mux.Vars(r)["id"])
	// Change response text based on error type
	if err != nil {
		h.Logger.Error(err)
		switch err {
		case sql.ErrNoRows:
			w.WriteHeader(http.StatusNotFound)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	w.WriteHeader(http.StatusOK)
	h.Logger.Info(fmt.Sprintf("%d\t%s\t%s", http.StatusOK, r.Method, r.URL.Path))
	json.NewEncoder(w).Encode(contact)
}

func (h *Handler) GetAllContacts(w http.ResponseWriter, r *http.Request) {
	contacts, err := h.db.GetAllContacts()
	if err != nil {
		h.Logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	h.Logger.Info(fmt.Sprintf("%d\t%s\t%s", http.StatusOK, r.Method, r.URL.Path))
	json.NewEncoder(w).Encode(contacts)
}

func (h *Handler) UpdateContact(w http.ResponseWriter, r *http.Request) {
	contact, err := h.db.GetContact(mux.Vars(r)["id"])
	if err != nil {
		h.Logger.Error(err)
		switch err {
		case sql.ErrNoRows:
			w.WriteHeader(http.StatusNotFound)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	json.NewDecoder(r.Body).Decode(&contact)
	err = h.db.UpdateContact(contact)
	if err != nil {
		h.Logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	h.Logger.Info(fmt.Sprintf("%d\t%s\t%s", http.StatusOK, r.Method, r.URL.Path))
	json.NewEncoder(w).Encode(contact)
}

func (h *Handler) DeleteContact(w http.ResponseWriter, r *http.Request) {
	contact := &store.Contact{}
	contact, err := h.db.GetContact(mux.Vars(r)["id"])
	if err != nil {
		h.Logger.Error(err)
		switch err {
		case sql.ErrNoRows:
			w.WriteHeader(http.StatusNotFound)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	if err := h.db.DeleteContact(mux.Vars(r)["id"]); err != nil {
		h.Logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	h.Logger.Info(fmt.Sprintf("%d\t%s\t%s", http.StatusOK, r.Method, r.URL.Path))
	json.NewEncoder(w).Encode(contact)
}
