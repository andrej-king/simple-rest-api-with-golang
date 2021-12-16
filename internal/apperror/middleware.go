package apperror

import (
	"errors"
	"net/http"
)

// appHandler wrap with same signature as request handler.
type appHandler func(w http.ResponseWriter, r *http.Request) error

// Middleware for requests
func Middleware(h appHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var appErr *AppError
		err := h(w, r)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")

			if errors.As(err, &appErr) {

				// foreach in custom errors
				if errors.Is(err, ErrNotFound) {
					w.WriteHeader(http.StatusNotFound)
					w.Write(ErrNotFound.Marshal())
					return
					//} else if (err, NoAuthErr) {
					//	w.WriteHeader(http.StatusUnauthorized)
					//	w.Write(NoAuthErr.Marshal())
					//	return
				}

				// in not in list
				err = err.(*AppError)
				w.WriteHeader(http.StatusBadRequest)
				w.Write(appErr.Marshal())
				return
			}
			// wrap system errors
			w.WriteHeader(http.StatusTeapot)
			w.Write(systemError(err).Marshal())
		}
	}
}
