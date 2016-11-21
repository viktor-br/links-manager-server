package implementation

import (
	"database/sql"
	"github.com/satori/go.uuid"
	"github.com/viktor-br/links-manager-server/core/config"
	"github.com/viktor-br/links-manager-server/core/dao"
	"github.com/viktor-br/links-manager-server/core/entities"
	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/postgresql"
	"os"
	"time"
	"github.com/viktor-br/links-manager-server/core/security"
)

func setUp() (*sql.DB, config.AppConfig, error) {
	connectionStr := os.Getenv("LMS_TEST_MAIN_STORAGE_CONNECTION")
	storageType := os.Getenv("LMS_TEST_MAIN_STORAGE_TYPE")
	secret := os.Getenv("LMS_TEST_SECRET")

	config := &config.AppConfigImpl{
		SecretVal: secret,
	}

	conn, err := sql.Open(storageType, connectionStr)
	if err != nil {
		return nil, nil, err
	}
	DB := reform.NewDB(conn, postgresql.Dialect, nil)

	// Clear all users and sessions in testing database
	_, err = DB.DeleteFrom(dao.UserTable.NewStruct().View(), "")
	if err != nil {
		return nil, nil, err
	}
	_, err = DB.DeleteFrom(dao.SessionTable.NewStruct().View(), "")
	if err != nil {
		return nil, nil, err
	}
	// Create default admin account
	admin := &dao.User{
		ID:        uuid.NewV4().String(),
		Username:  "admin",
		Password:  security.Hash("admin", config.Secret()),
		CreatedAt: time.Now(),
		Role:      entities.RoleAdminUser,
	}
	err = DB.Save(admin)
	if err != nil {
		return nil, nil, err
	}

	return conn, config, err
}
