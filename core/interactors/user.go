package interactors

import (
	"errors"
	"github.com/satori/go.uuid"
	"github.com/viktor-br/links-manager-server/core/config"
	"github.com/viktor-br/links-manager-server/core/entities"
	"github.com/viktor-br/links-manager-server/core/implementation"
	"github.com/viktor-br/links-manager-server/core/security"
	"time"
)

var (
	// ErrNotExists user not exists
	ErrNotExists = errors.New("User doesn't exist")
	// ErrWrongCredentials wrong credentials provided
	ErrWrongCredentials = errors.New("Wrong credentials")
	// ErrTokenExpired - token expired
	ErrTokenExpired = errors.New("Token expired")
)

// UserInteractor combines different implementations to process external requests.
type UserInteractor interface {
	Authenticate(username, password string) (*entities.User, *entities.Session, error)
	Authorize(string) (*entities.User, error)
	Create(*entities.User) error
}

// UserInteractorImpl implements UserInteractor.
type UserInteractorImpl struct {
	config            config.AppConfig
	userRepository    implementation.UserRepository
	sessionRepository implementation.SessionRepository
}

// NewUserInteractor constructs UserInteractor.
func NewUserInteractor(
	config config.AppConfig,
	userRepository implementation.UserRepository,
	sessionRepository implementation.SessionRepository,
) (UserInteractor, error) {
	return UserInteractorImpl{
		config:            config,
		userRepository:    userRepository,
		sessionRepository: sessionRepository,
	}, nil
}

// Authenticate generates access token.
func (userInteractor UserInteractorImpl) Authenticate(username, password string) (*entities.User, *entities.Session, error) {
	userEntity, err := userInteractor.userRepository.FindByUsername(username)
	if err != nil {
		return nil, nil, err
	}

	if userEntity == nil {
		return nil, nil, ErrNotExists
	}
	passwordHash := security.Hash(password, userInteractor.config.Secret())

	if userEntity.Password != passwordHash {
		return nil, nil, ErrWrongCredentials
	}

	session := new(entities.Session)
	session.ExpiresOn = time.Now().Add(24 * time.Hour)
	// TODO save RemoteIP and UserAgent
	err = userInteractor.sessionRepository.Store(session)
	if err != nil {
		return nil, nil, err
	}

	return userEntity, session, nil
}

// Create implements new user creation.
func (userInteractor UserInteractorImpl) Create(user *entities.User) error {
	user.ID = uuid.NewV4().String()
	return userInteractor.userRepository.Store(user)
}

// Authorize checks if user authorized on system.
func (userInteractor UserInteractorImpl) Authorize(token string) (*entities.User, error) {
	session, err := userInteractor.sessionRepository.FindByID(token)
	if err != nil {
		return nil, err
	}
	if session == nil {
		return nil, errors.New("Session is nil")
	}
	if session.ExpiresOn.Before(time.Now()) {
		return nil, ErrTokenExpired
	}
	return session.User, nil
}
