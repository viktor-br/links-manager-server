package mocks

import (
	"github.com/viktor-br/links-manager-server/core/entities"
)

// SessionRepositoryMock mocks session repository implementation.
type SessionRepositoryMock struct {
	FindByIDImpl func(id string) (*entities.Session, error)
	StoreImpl    func(session *entities.Session) error
}

// FindByID mocks search session by ID.
func (sessionRepositoryMock *SessionRepositoryMock) FindByID(id string) (*entities.Session, error) {
	return sessionRepositoryMock.FindByIDImpl(id)
}

// Store mocks SessionRepository Store method.
func (sessionRepositoryMock *SessionRepositoryMock) Store(session *entities.Session) error {
	return sessionRepositoryMock.StoreImpl(session)
}
