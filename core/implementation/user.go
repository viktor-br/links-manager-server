package implementation

import (
	"database/sql"
	"github.com/satori/go.uuid"
	"github.com/viktor-br/links-manager-server/core/config"
	"github.com/viktor-br/links-manager-server/core/dao"
	"github.com/viktor-br/links-manager-server/core/entities"
	reform "gopkg.in/reform.v1"
)

// UserRepository represent storage for user entities.
type UserRepository interface {
	FindByUsername(string) (*entities.User, error)
	Store(*entities.User) error
}

// UserRepositoryImpl implemen--ts UserRepository.
type UserRepositoryImpl struct {
	config config.AppConfig
	db     *reform.DB
}

// NewUserRepository create UserRepository instance.
func NewUserRepository(config config.AppConfig, db *reform.DB) UserRepository {
	return &UserRepositoryImpl{
		config: config,
		db:     db,
	}
}

// FindByUsername search user by username.
func (userRepository *UserRepositoryImpl) FindByUsername(username string) (*entities.User, error) {
	userStruct := dao.UserTable.NewStruct()
	err := userRepository.db.FindOneTo(userStruct, dao.UserFieldNameUsername, username)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	userRecord := userStruct.(*dao.User)
	return CreateUserEntityFromDAO(userRecord), nil
}

// Store saves user entity.
func (userRepository *UserRepositoryImpl) Store(user *entities.User) error {
	if user.ID == "" {
		user.ID = uuid.NewV4().String()
	}
	userRecord := CreateUserDAOFromEntity(user)

	return userRepository.db.Save(userRecord)
}

// CreateUserDAOFromEntity create DAO for core.entities.User
func CreateUserDAOFromEntity(user *entities.User) *dao.User {
	return &dao.User{
		ID:        user.ID,
		Username:  user.Username,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Role:      user.Role,
	}
}

// CreateUserEntityFromDAO create entities.User pointer from DAO
func CreateUserEntityFromDAO(userRecord *dao.User) *entities.User {
	return &entities.User{
		ID:        userRecord.ID,
		Username:  userRecord.Username,
		Password:  userRecord.Password,
		CreatedAt: userRecord.CreatedAt,
		UpdatedAt: userRecord.UpdatedAt,
		Role:      userRecord.Role,
	}
}
