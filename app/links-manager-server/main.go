package main

import (
	"database/sql"
	"fmt"
	"github.com/go-kit/kit/log"
	_ "github.com/lib/pq"
	"github.com/viktor-br/links-manager-server/app/controllers"
	"github.com/viktor-br/links-manager-server/app/handlers"
	l "github.com/viktor-br/links-manager-server/app/log"
	"github.com/viktor-br/links-manager-server/core/config"
	"github.com/viktor-br/links-manager-server/core/implementation"
	"github.com/viktor-br/links-manager-server/core/interactors"
	reform "gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/postgresql"
	dbLog "log"
	"net/http"
	"os"
)

func main() {
	port := "8080"
	if val := os.Getenv("LMS_PORT"); val != "" {
		port = val
	}
	connectionStr := os.Getenv("LMS_MAIN_STORAGE_CONNECTION")
	storageType := os.Getenv("LMS_MAIN_STORAGE_TYPE")
	secret := os.Getenv("LMS_SECRET")
	instanceID := os.Getenv("LMS_INSTANCE_ID")

	var defaultLlogger log.Logger
	defaultLlogger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))
	defaultLlogger = log.NewContext(defaultLlogger).With("ts", log.DefaultTimestampUTC)
	if instanceID != "" {
		defaultLlogger = log.NewContext(defaultLlogger).With("instance_id", instanceID)
	}

	conn, err := sql.Open(storageType, connectionStr)
	if err != nil {
		defaultLlogger.Log(l.LogMessage, fmt.Sprintf("Faile to create conn: %s", err.Error()))
		return
	}
	dbLogger := dbLog.New(os.Stderr, "SQL: ", dbLog.Flags())
	DB := reform.NewDB(conn, postgresql.Dialect, reform.NewPrintfLogger(dbLogger.Printf))

	// TODO Move setting values to proper place
	config := &config.AppConfigImpl{
		SecretVal: secret,
	}

	userRepository := implementation.NewUserRepository(config, DB)
	sessionRepository := implementation.NewSessionRepository(DB)
	userInteractor, err := interactors.NewUserInteractor(config, userRepository, sessionRepository)
	if err != nil {
		defaultLlogger.Log(l.LogMessage, fmt.Sprintf("Faile to start, unable to create interactor %s", err.Error()))
		return
	}
	userController := controllers.NewUserController(userInteractor, defaultLlogger)
	userHandler := handlers.NewUserHandler(userController, userInteractor, defaultLlogger)
	userAuthenticateHandler := handlers.NewUserAuthenticateHandler(userController, userInteractor, defaultLlogger)

	http.Handle("/api/v1/user", userHandler)
	http.Handle("/api/v1/user/login", userAuthenticateHandler)
	http.ListenAndServe(":"+port, nil)
}
