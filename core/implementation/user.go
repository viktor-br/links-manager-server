package implementation

import (
	"github.com/satori/go.uuid"
	"github.com/viktor-br/links-manager-server/core/entities"
)

// UserRepository represent storage for user entities.
type UserRepository interface {
	FindByUsername(string) (*entities.User, error)
}

// UserRepositoryImpl implements UserRepository.
type UserRepositoryImpl struct {
}

// NewUserRepository create UserRepository instance.
func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

// FindByUsername search user by username.
func (userRepository *UserRepositoryImpl) FindByUsername(username string) (*entities.User, error) {
	return &entities.User{ID: uuid.NewV4().String(), Username: "test", Password: "test"}, nil
}
