package router

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"golang-api-template/internal/datastore"
	"golang-api-template/internal/logger"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	log logger.Logger
	ctx context.Context
	db  *datastore.Datastore
)

// NewHandler creates a new handler used to handle incoming requests
func ProvideRouter(context context.Context, logger logger.Logger, datastore *datastore.Datastore) *mux.Router {
	ctx = context
	log = logger
	db = datastore

	router := mux.NewRouter()
	router.Use(commonMiddleware)
	router.HandleFunc("/api/test/", Test).Methods("GET")
	router.HandleFunc("/api/read/", GetAllContacts).Methods("GET")
	router.HandleFunc("/api/read/{id}/", GetContact).Methods("GET")
	router.HandleFunc("/api/create/", CreateContact).Methods("POST")
	router.HandleFunc("/api/update/{id}/", UpdateContact).Methods("POST")
	router.HandleFunc("/api/delete/{id}/", DeleteContact).Methods("POST")
	return router
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		log.Debug(fmt.Sprintf("%s\t%s", r.Method, r.URL.Path))
		next.ServeHTTP(w, r)
	})
}

// Test is responsible for /api/test/
func Test(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// CreateContact is responsible for /api/create/, and taking in parameters as request body
func CreateContact(w http.ResponseWriter, r *http.Request) {
	newContact := &datastore.Contact{}
	err := json.NewDecoder(r.Body).Decode(&newContact)
	if err != nil {
		errorResponse(w, r, err, http.StatusBadRequest)
		return
	}

	err = db.InsertContact(ctx, newContact)
	if err != nil {
		errorResponse(w, r, err, http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(newContact)
	if err != nil {
		errorResponse(w, r, err, http.StatusInternalServerError)
		return
	}

}

// GetAllContacts is responsible for /api/read/{id}, returning a single contact
func GetContact(w http.ResponseWriter, r *http.Request) {
	contact, err := db.GetContact(ctx, mux.Vars(r)["id"])
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			errorResponse(w, r, err, http.StatusNotFound)
		default:
			errorResponse(w, r, err, http.StatusInternalServerError)
		}
		return
	}

	err = json.NewEncoder(w).Encode(contact)
	if err != nil {
		errorResponse(w, r, err, http.StatusInternalServerError)
		return
	}
}

// GetContact is responsible for /api/read/, returning all contacts
func GetAllContacts(w http.ResponseWriter, r *http.Request) {
	contacts, err := db.GetAllContacts(ctx)
	if err != nil {
		errorResponse(w, r, err, http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(contacts)
	if err != nil {
		errorResponse(w, r, err, http.StatusInternalServerError)
		return
	}
}

// UpdateContact is responsible for /api/update/{id}, updating a contact with the request body
func UpdateContact(w http.ResponseWriter, r *http.Request) {
	contact, err := db.GetContact(ctx, mux.Vars(r)["id"])
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			errorResponse(w, r, err, http.StatusNotFound)
		default:
			errorResponse(w, r, err, http.StatusInternalServerError)
		}
		return
	}

	err = json.NewDecoder(r.Body).Decode(&contact)
	if err != nil {
		errorResponse(w, r, err, http.StatusBadRequest)
		return
	}

	err = db.UpdateContact(ctx, contact)
	if err != nil {
		errorResponse(w, r, err, http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(contact)
	if err != nil {
		errorResponse(w, r, err, http.StatusInternalServerError)
		return
	}
}

// DeleteContact is responsible for /api/delete/{id}, deleting a contact with the specific id
func DeleteContact(w http.ResponseWriter, r *http.Request) {
	contact, err := db.GetContact(ctx, mux.Vars(r)["id"])
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			errorResponse(w, r, err, http.StatusNotFound)
		default:
			errorResponse(w, r, err, http.StatusInternalServerError)
		}
		return
	}

	if err := db.DeleteContact(ctx, mux.Vars(r)["id"]); err != nil {
		errorResponse(w, r, err, http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(contact)
	if err != nil {
		errorResponse(w, r, err, http.StatusInternalServerError)
		return
	}
}

func errorResponse(w http.ResponseWriter, r *http.Request, err error, status int) {
	log.Error(err)
	w.WriteHeader(status)
}
