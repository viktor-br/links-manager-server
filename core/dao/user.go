package dao

import (
	"time"
)

const (
	// UserFieldNameEmail email field name
	UserFieldNameEmail string = "email"
)

// User implement user DAO
//go:generate reform
//reform:users
type User struct {
	ID        string     `reform:"id,pk"`
	Email     string     `reform:"email"`
	Password  string     `reform:"password"`
	CreatedAt time.Time  `reform:"created_at"`
	UpdatedAt *time.Time `reform:"updated_at"`
	Role      int        `reform:"role"`
}
