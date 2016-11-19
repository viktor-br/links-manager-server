package dao

import (
	"time"
)

const (
	// UserFieldNameUsername username field name
	UserFieldNameUsername string = "username"
)

// User implement user DAO
//go:generate reform
//reform:users
type User struct {
	ID        string     `reform:"id,pk"`
	Username  string     `reform:"username"`
	Password  string     `reform:"password"`
	CreatedAt time.Time  `reform:"created_at"`
	UpdatedAt *time.Time `reform:"updated_at"`
	Role      int        `reform:"role"`
}
