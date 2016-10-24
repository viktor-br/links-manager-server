package interactors

import "github.com/viktor-br/links-manager-server/core/entities"

// UserInteractor combines different implementations to process external requests.
type UserInteractor interface {
	Authenticate(entities.User) (string, error)
	Authorize(string) (entities.User, error)
	Create(entities.User) (entities.User, error)
}

// UserInteractorImpl implements UserInteractor.
type UserInteractorImpl struct {
}

// NewUserInteractor contructs UserInteractor.
func NewUserInteractor() (UserInteractor, error) {
	return UserInteractorImpl{}, nil
}

// Authenticate generates access token.
func (userInteractor UserInteractorImpl) Authenticate(user entities.User) (string, error) {
	return "123", nil
}

// Create implements new user creation.
func (userInteractor UserInteractorImpl) Create(user entities.User) (entities.User, error) {
	return user, nil
}

// Authorize checks if user authorized on system.
func (userInteractor UserInteractorImpl) Authorize(token string) (entities.User, error) {
	return entities.User{Username: "test", Password: "test"}, nil
}