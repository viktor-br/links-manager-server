package implementation

import (
	"github.com/satori/go.uuid"
	"github.com/viktor-br/links-manager-server/core/entities"
)

// SessionRepository represents session storage.
type SessionRepository interface {
	FindByID(id string) (*entities.Session, error)
	Store(*entities.Session) error
}

// SessionRepositoryImpl implements session storage interface.
type SessionRepositoryImpl struct {
}

// NewSessionRepository creates session storage.
func NewSessionRepository() SessionRepository {
	return &SessionRepositoryImpl{}
}

// FindByID search session by ID.
func (sessionRepository *SessionRepositoryImpl) FindByID(id string) (*entities.Session, error) {
	return &entities.Session{
		ID:   id,
		User: &entities.User{ID: uuid.NewV4().String(), Username: "test", Password: "test"},
	}, nil
}

// Store saves session entity.
func (sessionRepository *SessionRepositoryImpl) Store(session *entities.Session) error {
	session.ID = uuid.NewV4().String()
	return nil
}
