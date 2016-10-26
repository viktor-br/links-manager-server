package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/satori/go.uuid"
	"github.com/viktor-br/links-manager-server/core/entities"
	"net/http"
	"testing"
)

// UserInteractorMock mocks UserInteractorImpl
type UserInteractorMock struct {
	AuthenticateImpl func(username, password string) (entities.User, string, error)
	AuthorizeImpl    func(string) (entities.User, error)
	CreateImpl       func(entities.User) (entities.User, error)
}

// Authenticate mocks method via implementation method.
func (userInteractorMock UserInteractorMock) Authenticate(username, password string) (entities.User, string, error) {
	return userInteractorMock.AuthenticateImpl(username, password)
}

func (userInteractorMock UserInteractorMock) Authorize(token string) (entities.User, error) {
	return userInteractorMock.AuthorizeImpl(token)
}

func (userInteractorMock UserInteractorMock) Create(user entities.User) (entities.User, error) {
	return userInteractorMock.CreateImpl(user)
}

func createRegularUser() entities.User {
	return entities.User{
		ID:       uuid.NewV4().String(),
		Username: "test",
		Password: "password",
		Role:     entities.RoleRegularUser,
	}
}

func createAdminUser() entities.User {
	return entities.User{
		ID:       uuid.NewV4().String(),
		Username: "test",
		Password: "password",
		Role:     entities.RoleAdminUser,
	}
}

func TestUserCreateEmptyToken(t *testing.T) {
	ctrl := &UserController{
		Interactor: &UserInteractorMock{
			AuthorizeImpl: func(string) (entities.User, error) {
				return entities.User{}, nil
			},
		},
	}

	r := NewHTTPRequestMock("", []byte{})

	w := NewResponseWriterMock()

	ctrl.Create(w, r)

	if w.WrittenHeader != http.StatusForbidden {
		t.Errorf("Expect %d status code, received %d", http.StatusForbidden, w.WrittenHeader)
	}
}

func TestUserCreateAuthorizeFailed(t *testing.T) {
	ctrl := &UserController{
		Interactor: &UserInteractorMock{
			AuthorizeImpl: func(string) (entities.User, error) {
				return entities.User{}, fmt.Errorf("Authorization failed")
			},
		},
	}

	r := NewHTTPRequestMock("124", []byte{})

	w := NewResponseWriterMock()

	ctrl.Create(w, r)

	if w.WrittenHeader != http.StatusForbidden {
		t.Errorf("Expect %d status code, received %d", http.StatusForbidden, w.WrittenHeader)
	}
}

func TestUserCreateByRegularUserFailed(t *testing.T) {
	u := createRegularUser()
	ctrl := &UserController{
		Interactor: &UserInteractorMock{
			AuthorizeImpl: func(string) (entities.User, error) {
				return u, nil
			},
			CreateImpl: func(entities.User) (entities.User, error) {
				return u, nil
			},
		},
	}

	userJSON, _ := json.Marshal(u)

	r := NewHTTPRequestMock("124", userJSON)

	w := NewResponseWriterMock()

	ctrl.Create(w, r)

	if w.WrittenHeader != http.StatusForbidden {
		t.Errorf("Expect %d status code, received %d", http.StatusForbidden, w.WrittenHeader)
	}
}

func TestUserCreateCorruptedJSON(t *testing.T) {
	u := createAdminUser()
	ctrl := &UserController{
		Interactor: &UserInteractorMock{
			AuthorizeImpl: func(string) (entities.User, error) {
				return u, nil
			},
			CreateImpl: func(entities.User) (entities.User, error) {
				return u, nil
			},
		},
	}

	r := NewHTTPRequestMock("124", []byte("corrupted json"))

	w := NewResponseWriterMock()

	ctrl.Create(w, r)

	if w.WrittenHeader != http.StatusBadRequest {
		t.Errorf("Expect %d status code, received %d", http.StatusBadRequest, w.WrittenHeader)
	}
}

func TestUserCreateFailed(t *testing.T) {
	u := createAdminUser()
	ctrl := &UserController{
		Interactor: &UserInteractorMock{
			AuthorizeImpl: func(string) (entities.User, error) {
				return u, nil
			},
			CreateImpl: func(entities.User) (entities.User, error) {
				return u, fmt.Errorf("Unable to create user")
			},
		},
	}
	userJSON, _ := json.Marshal(u)

	r := NewHTTPRequestMock("124", userJSON)

	w := NewResponseWriterMock()

	ctrl.Create(w, r)

	if w.WrittenHeader != http.StatusNotFound {
		t.Errorf("Expect %d status code, received %d", http.StatusNotFound, w.WrittenHeader)
	}
}

func TestUserCreateSuccess(t *testing.T) {
	u := createAdminUser()
	ctrl := &UserController{
		Interactor: &UserInteractorMock{
			AuthorizeImpl: func(string) (entities.User, error) {
				return u, nil
			},
			CreateImpl: func(entities.User) (entities.User, error) {
				return u, nil
			},
		},
	}
	userJSON, _ := json.Marshal(u)

	r := NewHTTPRequestMock("124", userJSON)

	w := NewResponseWriterMock()

	ctrl.Create(w, r)

	if w.WrittenHeader != http.StatusOK {
		t.Errorf("Expect %d status code, received %d", http.StatusOK, w.WrittenHeader)
	}

	if bytes.Compare(w.WrittenContent, userJSON) != 0 {
		t.Errorf("Expected JSON %s, received %s", w.WrittenContent, userJSON)
	}
}

func TestAuthenticateCorruptedJSON(t *testing.T) {
	u := createAdminUser()
	ctrl := &UserController{
		Interactor: &UserInteractorMock{
			AuthorizeImpl: func(string) (entities.User, error) {
				return u, nil
			},
			CreateImpl: func(entities.User) (entities.User, error) {
				return u, nil
			},
		},
	}

	r := NewHTTPRequestMock("124", []byte("Corrupted JSON"))

	w := NewResponseWriterMock()

	ctrl.Authenticate(w, r)

	if w.WrittenHeader != http.StatusBadRequest {
		t.Errorf("Expect %d status code, received %d", http.StatusBadRequest, w.WrittenHeader)
	}
}

func TestAuthenticateFailed(t *testing.T) {
	u := createAdminUser()
	ctrl := &UserController{
		Interactor: &UserInteractorMock{
			AuthorizeImpl: func(string) (entities.User, error) {
				return u, fmt.Errorf("User interactor creation failed")
			},
			AuthenticateImpl: func(string, string) (entities.User, string, error) {
				return u, "", fmt.Errorf("Authentication failed")
			},
		},
	}
	userJSON, _ := json.Marshal(u)

	r := NewHTTPRequestMock("124", userJSON)

	w := NewResponseWriterMock()

	ctrl.Authenticate(w, r)

	if w.WrittenHeader != http.StatusForbidden {
		t.Errorf("Expect %d status code, received %d", http.StatusForbidden, w.WrittenHeader)
	}
}

func TestAuthenticateSuccess(t *testing.T) {
	u := createAdminUser()
	token := "123"
	ctrl := &UserController{
		Interactor: &UserInteractorMock{
			AuthorizeImpl: func(string) (entities.User, error) {
				return u, fmt.Errorf("User interactor creation failed")
			},
			AuthenticateImpl: func(string, string) (entities.User, string, error) {
				return u, "123", nil
			},
		},
	}
	userJSON, _ := json.Marshal(u)

	r := NewHTTPRequestMock(token, userJSON)

	w := NewResponseWriterMock()

	ctrl.Authenticate(w, r)

	if w.WrittenHeader != http.StatusOK {
		t.Errorf("Expect %d status code, received %d", http.StatusOK, w.WrittenHeader)
	}

	givenToken := w.Header().Get(XAuthToken)
	if givenToken != token {
		t.Errorf("Expect token %s, received %s", token, w.WrittenHeader)
	}
}
