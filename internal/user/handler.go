package user

import (
	"fmt"
	"github.com/cdo-pand/simple-rest-api-with-golang/internal/apperror"
	"github.com/cdo-pand/simple-rest-api-with-golang/internal/handlers"
	"github.com/cdo-pand/simple-rest-api-with-golang/pkg/logging"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// hint check if correct realize interface
var _ handlers.Handler = &handler{}

const (
	usersURL = "/users"
	userURL  = "/users/:uuid"
)

type handler struct {
	logger *logging.Logger
}

// NewHandler get Handler interface
func NewHandler(logger *logging.Logger) handlers.Handler {
	return &handler{
		logger: logger,
	}
}

// Register realise Handler interface (from "internal/handlers/handler.go" file)
func (h *handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, usersURL, apperror.Middleware(h.GetList))
	router.HandlerFunc(http.MethodGet, userURL, apperror.Middleware(h.GetUserByUUID))
	router.HandlerFunc(http.MethodPost, usersURL, apperror.Middleware(h.CreateUser))
	router.HandlerFunc(http.MethodPut, userURL, apperror.Middleware(h.UpdateUser))
	router.HandlerFunc(http.MethodPatch, userURL, apperror.Middleware(h.PartiallyUpdateUser))
	router.HandlerFunc(http.MethodDelete, userURL, apperror.Middleware(h.DeleteUserByUUID))
}

// GetList get user list
func (h *handler) GetList(w http.ResponseWriter, r *http.Request) error {
	//w.WriteHeader(http.StatusOK)
	//_, err := w.Write([]byte("This is list of users"))
	//if err != nil {
	//	panic(err)
	//}
	//
	//return nil
	return apperror.ErrNotFound
}

// GetUserByUUID get user by uuid
func (h *handler) GetUserByUUID(w http.ResponseWriter, r *http.Request) error {
	//w.WriteHeader(http.StatusOK)
	//_, err := w.Write([]byte("Get user by uuid"))
	//if err != nil {
	//	panic(err)
	//}
	//
	//return nil
	return apperror.NewAppError(nil, "test", "test", "some code status")
}

// CreateUser create new user
func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) error {
	//w.WriteHeader(http.StatusCreated)
	//_, err := w.Write([]byte("Create new user"))
	//if err != nil {
	//	panic(err)
	//}
	//
	//return nil
	return fmt.Errorf("this is API error")
}

// UpdateUser fully update user data
func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusNoContent)
	_, err := w.Write([]byte(""))
	if err != nil {
		panic(err)
	}

	return nil
}

// PartiallyUpdateUser partially update user data
func (h *handler) PartiallyUpdateUser(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusNoContent)
	_, err := w.Write([]byte(""))
	if err != nil {
		panic(err)
	}

	return nil
}

// DeleteUserByUUID delete user by uuid
func (h *handler) DeleteUserByUUID(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusNoContent)
	_, err := w.Write([]byte(""))
	if err != nil {
		panic(err)
	}

	return nil
}
