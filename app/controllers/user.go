package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/viktor-br/links-manager-server/core/entities"
	"github.com/viktor-br/links-manager-server/core/interactors"
	"io/ioutil"
	"net/http"
)

// XAuthToken represents HTTP header for authentication token.
const XAuthToken = "X-Auth-Token"

// NewUserController constructs UserController.
func NewUserController() *UserController {
	return &UserController{
		interactors.UserInteractorImpl{},
	}
}

// UserController contains user CRUD actions.
type UserController struct {
	Interactor interactors.UserInteractor
}

// Create validates parameters, invokes interactor and return results.
func (userCtrl *UserController) Create(w http.ResponseWriter, r *http.Request) {
	var b []byte
	var user entities.User

	// Read user info
	userIdentifier := r.Header.Get(XAuthToken)
	if userIdentifier == "" {
		w.WriteHeader(http.StatusForbidden)
		fmt.Println("userPut empty token")
		return
	}

	currentUser, err := userCtrl.Interactor.Authorize(userIdentifier)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		fmt.Printf("userPut access for token %s forbidden: %s\n", userIdentifier, err.Error())
		return
	}

	isAllowedCreateUser, err := currentUser.IsAllowedCreateUser()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Printf("userPut failed to check user rights %s: %s\n", currentUser.Username, err.Error())
		return
	}

	if !isAllowedCreateUser {
		w.WriteHeader(http.StatusForbidden)
		fmt.Printf("userPut forbidden for user %s\n", currentUser.Username)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	input, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Printf("userPut unable to read input: %s\n", err.Error())
		return
	}

	err = json.Unmarshal(input, &user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Printf("userPut json parse failed: %s, %v\n", err.Error(), b)
		return
	}

	// Create user
	user, err = userCtrl.Interactor.Create(user)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Printf("userLogin unable to create user: %s\n", user.Username)
		return
	}

	userJSON, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Printf("userPut json encode failed: %s\n", err.Error())
		return
	}

	w.Write(userJSON)

	w.WriteHeader(http.StatusOK)

	fmt.Printf("userLogin user created in successfully: %s\n", user.Username)
}

// Authenticate checks parameters, invokes token generator and return in HTTP header.
func (userCtrl *UserController) Authenticate(w http.ResponseWriter, r *http.Request) {
	var b []byte
	var user entities.User

	input, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Printf("userLogin unable to read input: %s\n", err.Error())
		return
	}

	err = json.Unmarshal(input, &user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Printf("userLogin json parse failed: %s, %v\n", err.Error(), b)
		return
	}

	authToken, err := userCtrl.Interactor.Authenticate(user)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		fmt.Printf("userLogin wrong credentials: %s\n", user.Username)
		return
	}

	w.Header().Set(XAuthToken, authToken)

	w.WriteHeader(http.StatusOK)
	fmt.Printf("userLogin user logged in successfully: %s\n", user.Username)
}
