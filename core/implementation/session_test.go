package implementation

import (
	"github.com/satori/go.uuid"
	"github.com/viktor-br/links-manager-server/core/config"
	"github.com/viktor-br/links-manager-server/core/entities"
	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/postgresql"
	"testing"
	"time"
)

func TestSessionCreate(t *testing.T) {
	conn, err := setUpConnection()
	if err != nil {
		t.Errorf("connection init failed: %s", err.Error())
		return
	}
	config := &config.AppConfigImpl{
		SecretVal: "123",
	}
	DB := reform.NewDB(conn, postgresql.Dialect, nil)

	sessionRepository := NewSessionRepository(DB)

	userRepository := NewUserRepository(config, DB)

	userID := uuid.NewV4().String()
	email := "test@test.com"
	t1 := time.Now()
	user := &entities.User{
		ID:        userID,
		Username:  email,
		Password:  "test",
		CreatedAt: time.Now().AddDate(0, 0, -1),
		UpdatedAt: &t1,
		Role:      entities.RoleAdminUser,
	}
	err = userRepository.Store(user)

	if err != nil {
		t.Errorf("Store user failed: %s", err.Error())
		return
	}

	sessionID := uuid.NewV4().String()
	session := &entities.Session{
		ID:         sessionID,
		User:       user,
		RemoteAddr: "127.0.0.1",
		CreatedAt:  time.Now().AddDate(0, 0, -1),
		ExpiresAt:  time.Now(),
	}
	err = sessionRepository.Store(session)

	if err != nil {
		t.Errorf("Store session failed: %s", err.Error())
		return
	}

	savedSession, err := sessionRepository.FindByID(sessionID)
	if err != nil {
		t.Errorf("FindByID() failed: %s", err.Error())
		return
	}
	if savedSession == nil {
		t.Errorf("FindByID() return empty result: %s")
		return
	}

	if !CompareSessionEntities(session, savedSession) {
		t.Errorf("Expected session %v, actual session %v", session, savedSession)
	}
}

func CompareSessionEntities(session1, session2 *entities.Session) bool {
	if session1.ID != session2.ID || session1.RemoteAddr != session2.RemoteAddr {
		return false
	}

	if session1.CreatedAt.Sub(session2.CreatedAt).Seconds() == 0 {
		return false
	}

	if session1.ExpiresAt.Sub(session2.ExpiresAt).Seconds() == 0 {
		return false
	}

	return true
}
