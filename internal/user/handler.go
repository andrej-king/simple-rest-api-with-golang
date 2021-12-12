package user

import (
	"github.com/cdo-pand/simple-rest-api-with-golang/internal/handlers"
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
}

// NewHandler get Handler interface
func NewHandler() handlers.Handler {
	return &handler{}
}

// Register realise Handler interface (from "internal/handlers/handler.go" file)
func (h *handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, usersURL, h.GetList)
	router.GET(userURL, h.GetUserByUUID)
	router.POST(usersURL, h.CreateUser)
	router.PUT(userURL, h.UpdateUser)
	router.PATCH(userURL, h.PartiallyUpdateUser)
	router.DELETE(userURL, h.DeleteUserByUUID)
}

// GetList get user list
func (h *handler) GetList(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("This is list of users"))
	if err != nil {
		panic(err)
	}
}

// GetUserByUUID get user by uuid
func (h *handler) GetUserByUUID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("Get user by uuid"))
	if err != nil {
		panic(err)
	}
}

// CreateUser create new user
func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(http.StatusCreated)
	_, err := w.Write([]byte("Create new user"))
	if err != nil {
		panic(err)
	}
}

// UpdateUser fully update user data
func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(http.StatusNoContent)
	_, err := w.Write([]byte(""))
	if err != nil {
		panic(err)
	}
}

// PartiallyUpdateUser partially update user data
func (h *handler) PartiallyUpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(http.StatusNoContent)
	_, err := w.Write([]byte(""))
	if err != nil {
		panic(err)
	}
}

// DeleteUserByUUID delete user by uuid
func (h *handler) DeleteUserByUUID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(http.StatusNoContent)
	_, err := w.Write([]byte(""))
	if err != nil {
		panic(err)
	}
}
