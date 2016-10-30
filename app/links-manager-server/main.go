package main

import (
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/viktor-br/links-manager-server/app/controllers"
	"github.com/viktor-br/links-manager-server/app/handlers"
	l "github.com/viktor-br/links-manager-server/app/log"
	"github.com/viktor-br/links-manager-server/core/interactors"
	"net/http"
	"os"
)

func main() {
	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))
	logger = log.NewContext(logger).With("ts", log.DefaultTimestampUTC)
	logger = log.NewContext(logger).With("instance_id", 123)

	userInteractor, err := interactors.NewUserInteractor()
	if err != nil {
		logger.Log(l.LogMessage, fmt.Sprintf("Faile to start, unable to create interactor %s", err.Error()))
		return
	}
	userController := controllers.NewUserController(userInteractor, logger)
	userHandler := handlers.NewUserHandler(userController, userInteractor, logger)
	userAuthenticateHandler := handlers.NewUserAuthenticateHandler(userController, userInteractor, logger)

	http.Handle("/api/v1/user", userHandler)
	http.Handle("/api/v1/user/login", userAuthenticateHandler)
	http.ListenAndServe(":8080", nil)
}
