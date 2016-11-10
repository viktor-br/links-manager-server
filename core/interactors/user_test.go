package interactors

import (
	"errors"
	"github.com/satori/go.uuid"
	"github.com/viktor-br/links-manager-server/app/mocks"
	"github.com/viktor-br/links-manager-server/core/config"
	"github.com/viktor-br/links-manager-server/core/entities"
	"github.com/viktor-br/links-manager-server/core/security"
	"testing"
)

func TestUserAuthenticate(t *testing.T) {
	username := "username"
	password := "password"

	config := &config.AppConfigImpl{
		SecretVal: "123",
	}

	userRepository := &mocks.UserRepositoryMock{
		FindByUsernameImpl: func(username string) (*entities.User, error) {
			return &entities.User{Username: username, Password: security.Hash(password, config.Secret())}, nil
		},
	}
	sessionRepository := &mocks.SessionRepositoryMock{
		StoreImpl: func(session *entities.Session) error {
			session.ID = uuid.NewV4().String()

			return nil
		},
	}
	userInteractor, err := NewUserInteractor(config, userRepository, sessionRepository)

	if err != nil {
		t.Error("Unable to create user interactor")
	}

	user, session, err := userInteractor.Authenticate(username, password)

	if err != nil {
		t.Errorf("Expect error nil, %s obtained", err.Error())
	}

	if user == nil {
		t.Error("Expect user entity")
	}

	if session == nil {
		t.Error("Expect session entity")
	}
}

func TestUserAuthenticateUserSearchFailed(t *testing.T) {
	username := "username"
	password := "password"

	config := &config.AppConfigImpl{
		SecretVal: "123",
	}

	userRepository := &mocks.UserRepositoryMock{
		FindByUsernameImpl: func(username string) (*entities.User, error) {
			return nil, errors.New("FindByUsername failed")
		},
	}
	sessionRepository := &mocks.SessionRepositoryMock{
		StoreImpl: func(session *entities.Session) error {
			session.ID = uuid.NewV4().String()

			return nil
		},
	}
	userInteractor, err := NewUserInteractor(config, userRepository, sessionRepository)

	if err != nil {
		t.Error("Unable to create user interactor")
	}

	user, session, err := userInteractor.Authenticate(username, password)

	if err == nil {
		t.Error("Expect user not found error, nil obtained")
	}

	if user != nil {
		t.Error("Expect user is nil")
	}

	if session != nil {
		t.Error("Expect session is nil")
	}
}

func TestUserAuthenticateUserNotFound(t *testing.T) {
	username := "username"
	password := "password"

	config := &config.AppConfigImpl{
		SecretVal: "123",
	}

	userRepository := &mocks.UserRepositoryMock{
		FindByUsernameImpl: func(username string) (*entities.User, error) {
			return nil, nil
		},
	}
	sessionRepository := &mocks.SessionRepositoryMock{
		StoreImpl: func(session *entities.Session) error {
			session.ID = uuid.NewV4().String()

			return nil
		},
	}
	userInteractor, err := NewUserInteractor(config, userRepository, sessionRepository)

	if err != nil {
		t.Error("Unable to create user interactor")
	}

	user, session, err := userInteractor.Authenticate(username, password)

	if err != ErrNotExists {
		t.Errorf("Expect ErrNotExists error, %s obtained", err.Error())
	}

	if user != nil {
		t.Error("Expect user is nil")
	}

	if session != nil {
		t.Error("Expect session is nil")
	}
}

func TestUserAuthenticateWrongPasswordHash(t *testing.T) {
	username := "username"
	password := "password"

	config := &config.AppConfigImpl{
		SecretVal: "123",
	}

	userRepository := &mocks.UserRepositoryMock{
		FindByUsernameImpl: func(username string) (*entities.User, error) {
			return &entities.User{Username: username, Password: "wrong_hash"}, nil
		},
	}
	sessionRepository := &mocks.SessionRepositoryMock{
		StoreImpl: func(session *entities.Session) error {
			session.ID = uuid.NewV4().String()

			return nil
		},
	}
	userInteractor, err := NewUserInteractor(config, userRepository, sessionRepository)

	if err != nil {
		t.Error("Unable to create user interactor")
	}

	user, session, err := userInteractor.Authenticate(username, password)

	if err != ErrWrongCredentials {
		t.Errorf("Expect ErrWrongCredentials error, %s obtained", err.Error())
	}

	if user != nil {
		t.Error("Expect user is nil")
	}

	if session != nil {
		t.Error("Expect session is nil")
	}
}

func TestUserAuthenticateSessionStoreFailed(t *testing.T) {
	username := "username"
	password := "password"

	config := &config.AppConfigImpl{
		SecretVal: "123",
	}

	ErrSessionStoreFailed := errors.New("Session store failed")

	userRepository := &mocks.UserRepositoryMock{
		FindByUsernameImpl: func(username string) (*entities.User, error) {
			return &entities.User{Username: username, Password: security.Hash(password, config.Secret())}, nil
		},
	}
	sessionRepository := &mocks.SessionRepositoryMock{
		StoreImpl: func(session *entities.Session) error {
			session.ID = uuid.NewV4().String()

			return ErrSessionStoreFailed
		},
	}
	userInteractor, err := NewUserInteractor(config, userRepository, sessionRepository)

	if err != nil {
		t.Error("Unable to create user interactor")
	}

	user, session, err := userInteractor.Authenticate(username, password)

	if err != ErrSessionStoreFailed {
		t.Errorf("Expect ErrSessionStoreFailed error, %s obtained", err.Error())
	}

	if user != nil {
		t.Error("Expect user is nil")
	}

	if session != nil {
		t.Error("Expect session is nil")
	}
}
