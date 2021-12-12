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
	router.GET(usersURL, h.GetList)
	router.GET(userURL, h.GetUserByUUID)
	router.POST(usersURL, h.CreateUser)
	router.PUT(userURL, h.UpdateUser)
	router.PATCH(userURL, h.PartiallyUpdateUser)
	router.DELETE(userURL, h.DeleteUserByUUID)
}

// GetList get user list
func (h *handler) GetList(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	_, err := w.Write([]byte("This is list of users"))
	w.WriteHeader(http.StatusOK)
	if err != nil {
		panic(err)
	}
}

// GetUserByUUID get user by uuid
func (h *handler) GetUserByUUID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	_, err := w.Write([]byte("Get user by uuid"))
	w.WriteHeader(http.StatusOK)
	if err != nil {
		panic(err)
	}
}

// CreateUser create new user
func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	_, err := w.Write([]byte("Create new user"))
	w.WriteHeader(http.StatusCreated)
	if err != nil {
		panic(err)
	}
}

// UpdateUser fully update user data
func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	_, err := w.Write([]byte("Fully update user data"))
	w.WriteHeader(http.StatusNoContent)
	if err != nil {
		panic(err)
	}
}

// PartiallyUpdateUser partially update user data
func (h *handler) PartiallyUpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	_, err := w.Write([]byte("Partially update user data"))
	w.WriteHeader(http.StatusNoContent)
	if err != nil {
		panic(err)
	}
}

// DeleteUserByUUID delete user by uuid
func (h *handler) DeleteUserByUUID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	_, err := w.Write([]byte("Delete user by uuid"))
	w.WriteHeader(http.StatusNoContent)
	if err != nil {
		panic(err)
	}
}
