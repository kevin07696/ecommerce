package auth

import (
	"net/http"

	"github.com/kevin07696/ecommerce/domain"
)

func errorHandler(w http.ResponseWriter, err *domain.CustomError) {
	switch {
	case err.Err == domain.ErrDuplicateKey:
		w.WriteHeader(http.StatusConflict)
	case err.Err == domain.ErrValidation:
		w.WriteHeader(http.StatusBadRequest)
	case err.Err == domain.ErrNotFound:
		w.WriteHeader(http.StatusNoContent)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
}
