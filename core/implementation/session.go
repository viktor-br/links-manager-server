package implementation

import (
	"github.com/satori/go.uuid"
	"github.com/viktor-br/links-manager-server/core/entities"
)

// SessionRepository represents session storage.
type SessionRepository interface {
	Store(*entities.Session) error
}

// SessionRepositoryImpl implements session storage interface.
type SessionRepositoryImpl struct {
}

// NewSessionRepository creates session storage.
func NewSessionRepository() SessionRepository {
	return &SessionRepositoryImpl{}
}

// Store saves session entity.
func (sessionRepository *SessionRepositoryImpl) Store(session *entities.Session) error {
	session.ID = uuid.NewV4().String()
	return nil
}
