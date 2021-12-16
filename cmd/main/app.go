package main

import (
	"context"
	"fmt"
	"github.com/cdo-pand/simple-rest-api-with-golang/internal/config"
	"github.com/cdo-pand/simple-rest-api-with-golang/internal/user"
	"github.com/cdo-pand/simple-rest-api-with-golang/internal/user/db"
	"github.com/cdo-pand/simple-rest-api-with-golang/pkg/client/mongodb"
	"github.com/cdo-pand/simple-rest-api-with-golang/pkg/logging"
	"github.com/julienschmidt/httprouter"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("Create router")
	router := httprouter.New()

	cfg := config.GetConfig()

	cfgMongo := cfg.MongoDB
	mongoDBClient, err := mongodb.NewClient(context.Background(),
		cfgMongo.Host, cfgMongo.Port, cfgMongo.Username, cfgMongo.Password, cfgMongo.Database, cfgMongo.AuthDB)
	if err != nil {
		panic(err)
	}
	storage := db.NewStorage(mongoDBClient, cfg.MongoDB.Collection, logger)

	user1 := user.User{
		ID:           "",
		Email:        "app@test.com",
		Username:     "John",
		PasswordHash: "password",
	}
	user1ID, err := storage.Create(context.Background(), user1)
	if err != nil {
		panic(err)
	}
	logger.Infof("created user id: %s", user1ID)

	user2 := user.User{
		ID:           "",
		Email:        "app2@test.com",
		Username:     "Jack",
		PasswordHash: "otherPassword",
	}
	user2ID, err := storage.Create(context.Background(), user2)
	if err != nil {
		panic(err)
	}
	logger.Infof("created user id: %s", user2ID)

	user2Found, err := storage.FindOne(context.Background(), user2ID)
	if err != nil {
		panic(err)
	}
	fmt.Println(user2Found)

	user2Found.Email = "newEmail@test.com"
	err = storage.Update(context.Background(), user2Found)
	if err != nil {
		panic(err)
	}
	user2Updated, err := storage.FindOne(context.Background(), user2ID)
	if err != nil {
		panic(err)
	}
	fmt.Println(user2Updated)

	err = storage.Delete(context.Background(), user2ID)
	if err != nil {
		panic(err)
	}

	_, err = storage.FindOne(context.Background(), user2ID)
	if err != nil {
		fmt.Errorf("user not found. error: %v", err)
	}

	users, err := storage.FindAll(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println("users", users)

	logger.Info("Register user handler")
	handler := user.NewHandler(logger)
	handler.Register(router)

	appStart(router, cfg)
}

func appStart(router *httprouter.Router, cfg *config.Config) {
	logger := logging.GetLogger()
	logger.Info("Start application")

	var listener net.Listener
	var listenErr error

	// if socket type
	if cfg.Listen.Type == "sock" {
		logger.Info("Detect app path")
		appDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			logger.Fatal(err)
		}

		logger.Info("Create socket")
		socketPath := path.Join(appDir, "app.sock")

		logger.Info("Listen unix socket")
		listener, listenErr = net.Listen("unix", socketPath)
		logger.Infof("Server is unix socket %s", socketPath)
	} else {
		logger.Info("Listen tcp")
		listener, listenErr = net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port))
		logger.Infof("Server is listening %s:%s", cfg.Listen.BindIP, cfg.Listen.Port)
	}

	if listenErr != nil {
		logger.Fatal(listenErr)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Fatal(server.Serve(listener))
}
