package entities

const (
	// RoleRegularUser code of regular user
	RoleRegularUser = iota
	// RoleAdminUser code of admin user
	RoleAdminUser
)

// User represent user entity
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     int
}

// IsAllowedCreateUser checks if user allowed to create another users.
func (user *User) IsAllowedCreateUser() (bool, error) {
	return user.Role == RoleAdminUser, nil
}
