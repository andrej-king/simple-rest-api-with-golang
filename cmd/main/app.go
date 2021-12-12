package main

import (
	"github.com/cdo-pand/simple-rest-api-with-golang/internal/user"
	"github.com/cdo-pand/simple-rest-api-with-golang/pkg/logging"
	"github.com/julienschmidt/httprouter"
	"net"
	"net/http"
	"time"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("Create router")
	router := httprouter.New()

	logger.Info("Register user handler")
	handler := user.NewHandler(logger)
	handler.Register(router)

	appStart(router)
}

func appStart(router *httprouter.Router) {
	logger := logging.GetLogger()

	logger.Info("Start application")
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Info("Server is listening 0.0.0.0:8080")
	logger.Fatal(server.Serve(listen))
}
