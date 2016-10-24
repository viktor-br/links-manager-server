// Package handlers contains only matching URL to appropriate controllers and its actions.
package handlers

import (
	"fmt"
	"github.com/viktor-br/links-manager-server/app/controllers"
	"net/http"
)

// User handles resources for user CRUD operations, invokes appropriate controller actions.
func User(w http.ResponseWriter, r *http.Request) {
	ctrl := controllers.NewUserController()
	switch r.Method {
	case http.MethodPut:
		ctrl.Create(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Printf("userPut unsupported method %s\n", r.Method)
	}
}

// UserLogin handles resources for user authentication.
func UserLogin(w http.ResponseWriter, r *http.Request) {
	ctrl := controllers.NewUserController()
	if r.Method == http.MethodPost {
		ctrl.Authenticate(w, r)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Printf("userLogin unsupported method %s\n", r.Method)
	}
}
