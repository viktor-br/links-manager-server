package implementation

import (
	_ "github.com/lib/pq"
	"github.com/satori/go.uuid"
	"github.com/viktor-br/links-manager-server/core/entities"
	"github.com/viktor-br/links-manager-server/core/security"
	reform "gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/postgresql"
	"testing"
	"time"
)

func TestUserCreate(t *testing.T) {
	conn, config, err := setUp()
	if err != nil {
		t.Errorf("connection init failed: %s", err.Error())
		return
	}
	DB := reform.NewDB(conn, postgresql.Dialect, nil)

	userRepository := NewUserRepository(config, DB)

	id := uuid.NewV4().String()
	username := "test@test.com"
	password := "test"
	t1 := time.Now()
	user := &entities.User{
		ID:        id,
		Username:  username,
		Password:  security.Hash(password, config.Secret()),
		CreatedAt: time.Now().AddDate(0, 0, -1),
		UpdatedAt: &t1,
		Role:      entities.RoleAdminUser,
	}
	err = userRepository.Store(user)

	if err != nil {
		t.Errorf("Store user failed: %s", err.Error())
		return
	}

	savedUser, err := userRepository.FindByUsername(username)
	if err != nil {
		t.Errorf("FindByUsername() failed: %s", err.Error())
		return
	}
	if savedUser == nil {
		t.Errorf("FindByUsername() return empty result: %s")
		return
	}

	if !CompareUserEntities(user, savedUser) {
		t.Errorf("Expected user %v, actual user %v", user, savedUser)
	}
}

func CompareUserEntities(user1, user2 *entities.User) bool {
	if user1.ID != user2.ID || user1.Username != user2.Username || user1.Password != user2.Password {
		return false
	}

	if user1.CreatedAt.Sub(user2.CreatedAt).Seconds() >= 1 {
		return false
	}

	if user1.UpdatedAt == nil && user2.UpdatedAt != nil {
		return false
	}

	if user1.UpdatedAt != nil && user2.UpdatedAt == nil {
		return false
	}

	if user1.UpdatedAt != user2.UpdatedAt && user1.UpdatedAt.Sub(*user2.UpdatedAt).Seconds() >= 1 {
		return false
	}

	if user1.Role != user2.Role {
		return false
	}

	return true
}
