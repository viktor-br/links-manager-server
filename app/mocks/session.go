package mocks

import (
	"github.com/viktor-br/links-manager-server/core/entities"
)

// SessionRepositoryMock mocks session repository implementation.
type SessionRepositoryMock struct {
	StoreImpl func(session *entities.Session) error
}

// Store mocks SessionRepository Store method.
func (sessionRepositoryMock *SessionRepositoryMock) Store(session *entities.Session) error {
	return sessionRepositoryMock.StoreImpl(session)
}
