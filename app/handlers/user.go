// Package handlers contains only matching URL to appropriate controllers and its actions.
package handlers

import (
	"fmt"
	"github.com/viktor-br/links-manager-server/app/controllers"
	"github.com/viktor-br/links-manager-server/app/log"
	"github.com/viktor-br/links-manager-server/core/interactors"
	"net/http"
)

// UserHandler routes user CRUD request to user controller.
type UserHandler struct {
	Controller controllers.UserController
	Logger     log.Logger
}

// UserAuthenticateHandler routes authentication request to user controller.
type UserAuthenticateHandler struct {
	Controller controllers.UserController
	Logger     log.Logger
}

// NewUserHandler is UserHandler constructor.
func NewUserHandler(userController controllers.UserController, userInteractor interactors.UserInteractor, logger log.Logger) *UserHandler {
	return &UserHandler{
		Controller: userController,
		Logger:     logger,
	}
}

// NewUserAuthenticateHandler is UserAuthenticateHandler constructor.
func NewUserAuthenticateHandler(userController controllers.UserController, userInteractor interactors.UserInteractor, logger log.Logger) *UserAuthenticateHandler {
	return &UserAuthenticateHandler{
		Controller: userController,
		Logger:     logger,
	}
}

func (userHandler *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		userHandler.Controller.Create(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		if userHandler.Logger != nil {
			userHandler.Logger.Log(
				log.LogRequestURI, r.RequestURI,
				log.LogRemoteAddr, r.RemoteAddr,
				log.LogHTTPStatus, http.StatusMethodNotAllowed,
				log.LogController, "user",
				log.LogMessage, fmt.Sprintf("Method %s Not Allowed", r.Method),
			)
		}
	}
}

func (userAuthenticateHandler *UserAuthenticateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		userAuthenticateHandler.Controller.Authenticate(w, r)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		if userAuthenticateHandler.Logger != nil {
			userAuthenticateHandler.Logger.Log(
				log.LogRequestURI, r.RequestURI,
				log.LogRemoteAddr, r.RemoteAddr,
				log.LogHTTPStatus, http.StatusMethodNotAllowed,
				log.LogController, "user",
				log.LogMessage, fmt.Sprintf("Method %s Not Allowed", r.Method),
			)
		}
	}
}
