package implementation

import (
	"github.com/satori/go.uuid"
	"github.com/viktor-br/links-manager-server/core/dao"
	"github.com/viktor-br/links-manager-server/core/entities"
	"gopkg.in/reform.v1"
)

// SessionRepository represents session storage.
type SessionRepository interface {
	FindByID(id string) (*entities.Session, error)
	Store(*entities.Session) error
}

// SessionRepositoryImpl implements session storage interface.
type SessionRepositoryImpl struct {
	db *reform.DB
}

// NewSessionRepository creates session storage.
func NewSessionRepository(db *reform.DB) SessionRepository {
	return &SessionRepositoryImpl{
		db: db,
	}
}

// FindByID search session by ID.
func (sessionRepository *SessionRepositoryImpl) FindByID(id string) (*entities.Session, error) {
	session, err := sessionRepository.db.FindByPrimaryKeyFrom(dao.SessionTable, id)
	if err != nil {
		return nil, err
	}
	sessionRecord := session.(*dao.Session)

	user, err := sessionRepository.db.FindByPrimaryKeyFrom(dao.UserTable, sessionRecord.UserID)
	if err != nil {
		return nil, err
	}
	userRecord := user.(*dao.User)
	sessionEntity := CreateSessionEntityFromDAO(sessionRecord)
	sessionEntity.User = CreateUserEntityFromDAO(userRecord)

	return sessionEntity, nil
}

// Store saves session entity.
func (sessionRepository *SessionRepositoryImpl) Store(session *entities.Session) error {
	if session.ID == "" {
		session.ID = uuid.NewV4().String()
	}
	userRecord := CreateSessionDAOFromEntity(session)

	return sessionRepository.db.Save(userRecord)
}

// CreateSessionEntityFromDAO create entities.Session instance from session DAO.
func CreateSessionEntityFromDAO(sessionRecord *dao.Session) *entities.Session {
	return &entities.Session{
		ID:         sessionRecord.ID,
		RemoteAddr: sessionRecord.RemoteAddr,
		CreatedAt:  sessionRecord.CreatedAt,
		ExpiresAt:  sessionRecord.ExpiresAt,
	}
}

// CreateSessionDAOFromEntity create DAO instance from session entities.Session.
func CreateSessionDAOFromEntity(session *entities.Session) *dao.Session {
	userID := ""
	if session.User != nil {
		userID = session.User.ID
	}
	return &dao.Session{
		ID:         session.ID,
		UserID:     userID,
		RemoteAddr: session.RemoteAddr,
		CreatedAt:  session.CreatedAt,
		ExpiresAt:  session.ExpiresAt,
	}
}
