package main

import (
	"github.com/go-kit/kit/log"
	"github.com/viktor-br/links-manager-server/app/handlers"
	"net/http"
	"os"
)

func main() {
	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))
	logger = log.NewContext(logger).With("ts", log.DefaultTimestampUTC)
	logger = log.NewContext(logger).With("instance_id", 123)

	userHandler := handlers.NewUserHandler(logger)
	userAuthenticateHandler := handlers.NewUserAuthenticateHandler(logger)

	http.Handle("/api/v1/user", userHandler)
	http.Handle("/api/v1/user/login", userAuthenticateHandler)
	http.ListenAndServe(":8080", nil)
}
