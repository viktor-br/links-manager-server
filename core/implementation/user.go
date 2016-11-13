package implementation

import (
	"github.com/satori/go.uuid"
	"github.com/viktor-br/links-manager-server/core/entities"
	"github.com/viktor-br/links-manager-server/core/security"
	"github.com/viktor-br/links-manager-server/core/config"
)

// UserRepository represent storage for user entities.
type UserRepository interface {
	FindByUsername(string) (*entities.User, error)
	Store(*entities.User) error
}

// UserRepositoryImpl implements UserRepository.
type UserRepositoryImpl struct {
	config config.AppConfig
}

// NewUserRepository create UserRepository instance.
func NewUserRepository(config config.AppConfig) UserRepository {
	return &UserRepositoryImpl{
		config: config,
	}
}

// FindByUsername search user by username.
func (userRepository *UserRepositoryImpl) FindByUsername(username string) (*entities.User, error) {
	return &entities.User{ID: uuid.NewV4().String(), Username: "test", Password: security.Hash("test", userRepository.config.Secret())}, nil
}

// Store saves user entity.
func (userRepository *UserRepositoryImpl) Store(user *entities.User) error {
	return nil
}
