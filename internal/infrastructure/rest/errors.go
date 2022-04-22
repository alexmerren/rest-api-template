package rest

import (
	"encoding/json"
	"errors"
	"net/http"
	"rest-api-template/internal/domain/entities"
)

const (
	publicInternalErrorMessage   = "An internal error occurred"
	publicBadRequestErrorMessage = "A bad request was given"
	publicNotFoundErrorMessage   = "A resource could not be found"
)

type ErrorResponse struct {
	StatusCode  int    `json:"status_code"`
	ErrorString string `json:"error_string"`
}

// nolint:errcheck // No point in trying to check for encoding errors here.
func HandleError(w http.ResponseWriter, r *http.Request, err error) {
	notFoundErrorType := &entities.NotFoundError{}
	if errors.As(err, &notFoundErrorType) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(&ErrorResponse{
			StatusCode:  http.StatusNotFound,
			ErrorString: publicNotFoundErrorMessage,
		})
		return
	}

	badRequestErrorType := &entities.BadRequestError{}
	if errors.As(err, &badRequestErrorType) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&ErrorResponse{
			StatusCode:  http.StatusBadRequest,
			ErrorString: publicBadRequestErrorMessage,
		})
		return
	}

	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(&ErrorResponse{
		StatusCode:  http.StatusInternalServerError,
		ErrorString: publicInternalErrorMessage,
	})
}
