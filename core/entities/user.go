package entities

import (
	"time"
)

const (
	// RoleRegularUser code of regular user
	RoleRegularUser = iota
	// RoleAdminUser code of admin user
	RoleAdminUser
)

// User represent user entity
type User struct {
	ID        string
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt *time.Time
	Role      int
}

// IsAllowedCreateUser checks if user allowed to create another users.
func (user *User) IsAllowedCreateUser() (bool, error) {
	return user.Role == RoleAdminUser, nil
}
