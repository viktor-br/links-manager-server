package implementation

import (
	"database/sql"
	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/postgresql"
	"github.com/viktor-br/links-manager-server/core/dao"
)

func setUpConnection() (*sql.DB, error) {
	conn, err := sql.Open("postgres", "postgres://localhost:5432/test?sslmode=disable")
	if err != nil {
		return nil, err
	}
	DB := reform.NewDB(conn, postgresql.Dialect, nil)

	// Clear all users and sessions in testing database
	DB.DeleteFrom(dao.UserTable.NewStruct().View(), "")
	DB.DeleteFrom(dao.SessionTable.NewStruct().View(), "")

	return conn, err
}

