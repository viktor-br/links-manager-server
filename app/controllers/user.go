package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/viktor-br/links-manager-server/app/log"
	"github.com/viktor-br/links-manager-server/core/entities"
	"github.com/viktor-br/links-manager-server/core/interactors"
	"io/ioutil"
	"net"
	"net/http"
)

// XAuthToken represents HTTP header for authentication token.
const XAuthToken = "X-Auth-Token"

// UserAuth structure to parse login request into.
type UserAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// UserController provide contract to interact with controller.
type UserController interface {
	Create(w http.ResponseWriter, r *http.Request)
	Authenticate(w http.ResponseWriter, r *http.Request)
	Log(args ...interface{})
}

// UserControllerImpl contains user CRUD actions.
type UserControllerImpl struct {
	Logger     log.Logger
	Interactor interactors.UserInteractor
}

// NewUserController constructs UserController.
func NewUserController(userInteractor interactors.UserInteractor, logger log.Logger) *UserControllerImpl {
	return &UserControllerImpl{
		Logger:     logger,
		Interactor: userInteractor,
	}
}

// Log checks if attached logger and uses it.
func (userCtrl *UserControllerImpl) Log(args ...interface{}) {
	if userCtrl.Logger != nil {
		userCtrl.Logger.Log(args...)
	}
}

// Create validates parameters, invokes interactor and return results.
func (userCtrl *UserControllerImpl) Create(w http.ResponseWriter, r *http.Request) {
	var user entities.User

	method := "user::create"

	// Read user info
	userIdentifier := r.Header.Get(XAuthToken)
	if userIdentifier == "" {
		w.WriteHeader(http.StatusForbidden)
		userCtrl.Log(
			log.LogRequestURI, r.RequestURI,
			log.LogRemoteAddr, r.RemoteAddr,
			log.LogHTTPStatus, http.StatusForbidden,
			log.LogController, method,
			log.LogMessage, "empty token",
		)
		return
	}

	currentUser, err := userCtrl.Interactor.Authorize(userIdentifier)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		userCtrl.Log(
			log.LogRequestURI, r.RequestURI,
			log.LogRemoteAddr, r.RemoteAddr,
			log.LogHTTPStatus, http.StatusForbidden,
			log.LogController, method,
			log.LogToken, userIdentifier,
			log.LogMessage, err.Error(),
		)
		return
	}

	isAllowedCreateUser, err := currentUser.IsAllowedCreateUser()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		userCtrl.Log(
			log.LogRequestURI, r.RequestURI,
			log.LogRemoteAddr, r.RemoteAddr,
			log.LogHTTPStatus, http.StatusInternalServerError,
			log.LogController, method,
			log.LogToken, userIdentifier,
			log.LogUserID, currentUser.ID,
			log.LogMessage, err.Error(),
		)
		return
	}

	if !isAllowedCreateUser {
		w.WriteHeader(http.StatusForbidden)
		userCtrl.Log(
			log.LogRequestURI, r.RequestURI,
			log.LogRemoteAddr, r.RemoteAddr,
			log.LogHTTPStatus, http.StatusForbidden,
			log.LogController, method,
			log.LogToken, userIdentifier,
			log.LogUserID, currentUser.ID,
			log.LogMessage, "Not allowed to create new user",
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	input, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		userCtrl.Log(
			log.LogRequestURI, r.RequestURI,
			log.LogRemoteAddr, r.RemoteAddr,
			log.LogHTTPStatus, http.StatusBadRequest,
			log.LogController, method,
			log.LogToken, userIdentifier,
			log.LogUserID, currentUser.ID,
			log.LogMessage, err.Error(),
		)
		return
	}

	err = json.Unmarshal(input, &user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		userCtrl.Log(
			log.LogRequestURI, r.RequestURI,
			log.LogRemoteAddr, r.RemoteAddr,
			log.LogHTTPStatus, http.StatusBadRequest,
			log.LogController, method,
			log.LogToken, userIdentifier,
			log.LogUserID, currentUser.ID,
			log.LogMessage, fmt.Sprintf("json parse failed: %s", err.Error()),
		)
		fmt.Println(err.Error())
		return
	}

	// Create user
	err = userCtrl.Interactor.Create(&user)
	w.WriteHeader(http.StatusConflict)
	if err == interactors.ErrUserAlreadyExists {
		userCtrl.Log(
			log.LogRequestURI, r.RequestURI,
			log.LogRemoteAddr, r.RemoteAddr,
			log.LogHTTPStatus, http.StatusConflict,
			log.LogController, method,
			log.LogToken, userIdentifier,
			log.LogUserID, currentUser.ID,
			log.LogMessage, fmt.Sprintf("unable to create user %s: %s", user.Username, err.Error()),
		)
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		userCtrl.Log(
			log.LogRequestURI, r.RequestURI,
			log.LogRemoteAddr, r.RemoteAddr,
			log.LogHTTPStatus, http.StatusNotFound,
			log.LogController, method,
			log.LogToken, userIdentifier,
			log.LogUserID, currentUser.ID,
			log.LogMessage, fmt.Sprintf("unable to create user %s: %s", user.Username, err.Error()),
		)
		return
	}

	userJSON, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		userCtrl.Log(
			log.LogRequestURI, r.RequestURI,
			log.LogRemoteAddr, r.RemoteAddr,
			log.LogHTTPStatus, http.StatusInternalServerError,
			log.LogController, method,
			log.LogToken, userIdentifier,
			log.LogUserID, currentUser.ID,
			log.LogMessage, fmt.Sprintf("json encode failed: %s", err.Error()),
		)
		return
	}

	_, err = w.Write(userJSON)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		userCtrl.Log(
			log.LogRequestURI, r.RequestURI,
			log.LogRemoteAddr, r.RemoteAddr,
			log.LogHTTPStatus, http.StatusInternalServerError,
			log.LogController, method,
			log.LogToken, userIdentifier,
			log.LogUserID, currentUser.ID,
			log.LogMessage, fmt.Sprintf("unable to write response for %s: %s", user.Username, err.Error()),
		)
		return
	}

	w.WriteHeader(http.StatusOK)
	userCtrl.Log(
		log.LogRequestURI, r.RequestURI,
		log.LogRemoteAddr, r.RemoteAddr,
		log.LogHTTPStatus, http.StatusOK,
		log.LogController, method,
		log.LogToken, userIdentifier,
		log.LogUserID, currentUser.ID,
		log.LogMessage, fmt.Sprintf("user %s created successfully", user.Username),
	)
}

// Authenticate checks parameters, invokes token generator and return in HTTP header.
func (userCtrl *UserControllerImpl) Authenticate(w http.ResponseWriter, r *http.Request) {
	var userAuth UserAuth

	method := "user::authenticate"

	input, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		userCtrl.Log(
			log.LogRequestURI, r.RequestURI,
			log.LogRemoteAddr, r.RemoteAddr,
			log.LogHTTPStatus, http.StatusBadRequest,
			log.LogController, method,
			log.LogMessage, err.Error(),
		)
		return
	}

	err = json.Unmarshal(input, &userAuth)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		userCtrl.Log(
			log.LogRequestURI, r.RequestURI,
			log.LogRemoteAddr, r.RemoteAddr,
			log.LogHTTPStatus, http.StatusBadRequest,
			log.LogController, method,
			log.LogMessage, fmt.Sprintf("json parse failed: %s", err.Error()),
		)
		return
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)

	if err != nil {
		// Just log this error, but it's not critical
		userCtrl.Log(
			log.LogRequestURI, r.RequestURI,
			log.LogRemoteAddr, r.RemoteAddr,
			log.LogController, method,
			log.LogMessage, fmt.Sprintf("split host failed: %s", err.Error()),
		)
	}
	user, session, err := userCtrl.Interactor.Authenticate(userAuth.Username, userAuth.Password, ip)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		userCtrl.Log(
			log.LogRequestURI, r.RequestURI,
			log.LogRemoteAddr, r.RemoteAddr,
			log.LogHTTPStatus, http.StatusForbidden,
			log.LogController, method,
			log.LogUserID, userAuth.Username,
			log.LogMessage, fmt.Sprintf("authentication failed: %s", err.Error()),
		)
		return
	}

	w.Header().Set(XAuthToken, session.ID)

	w.WriteHeader(http.StatusOK)

	userCtrl.Log(
		log.LogRequestURI, r.RequestURI,
		log.LogRemoteAddr, r.RemoteAddr,
		log.LogHTTPStatus, http.StatusOK,
		log.LogController, method,
		log.LogUserID, user.ID,
	)
}
