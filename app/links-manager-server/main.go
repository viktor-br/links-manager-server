package main

import (
	"net/http"
	"github.com/viktor-br/links-manager-server/app/handlers"
)

func main() {
	http.HandleFunc("/api/v1/user", handlers.User)
	http.HandleFunc("/api/v1/user/login", handlers.UserLogin)
	//http.HandleFunc("/api/v1/item/link", userPut)
	//http.HandleFunc("/api/v1/item/keyphrase", userPut)
	http.ListenAndServe(":8080", nil)
}


