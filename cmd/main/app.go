package main

import (
	"fmt"
	"github.com/cdo-pand/simple-rest-api-with-golang/internal/user"
	"github.com/julienschmidt/httprouter"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Create router")
	router := httprouter.New()

	fmt.Println("Register user handler")
	handler := user.NewHandler()
	handler.Register(router)

	appStart(router)
}

func appStart(router *httprouter.Router) {
	fmt.Println("Start application")
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Server is listening 0.0.0.0:8080")
	log.Fatalln(server.Serve(listen))
}
