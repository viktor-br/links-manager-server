package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/satori/go.uuid"
	"github.com/viktor-br/links-manager-server/app/mocks"
	"github.com/viktor-br/links-manager-server/core/entities"
	"net/http"
	"testing"
)

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
	ctrl := &UserControllerImpl{
		Interactor: &mocks.UserInteractorMock{
			AuthorizeImpl: func(string) (entities.User, error) {
				return entities.User{}, nil
			},
		},
	}

	r := mocks.NewHTTPRequestMock([]byte{})

	w := mocks.NewResponseWriterMock()

	ctrl.Create(w, r)

	if w.WrittenHeader != http.StatusForbidden {
		t.Errorf("Expect %d status code, received %d", http.StatusForbidden, w.WrittenHeader)
	}
}

func TestUserCreateAuthorizeFailed(t *testing.T) {
	ctrl := &UserControllerImpl{
		Interactor: &mocks.UserInteractorMock{
			AuthorizeImpl: func(string) (entities.User, error) {
				return entities.User{}, fmt.Errorf("Authorization failed")
			},
		},
	}

	r := mocks.NewHTTPRequestMock([]byte{})
	r.Header.Set(XAuthToken, "124")

	w := mocks.NewResponseWriterMock()

	ctrl.Create(w, r)

	if w.WrittenHeader != http.StatusForbidden {
		t.Errorf("Expect %d status code, received %d", http.StatusForbidden, w.WrittenHeader)
	}
}

func TestUserCreateByRegularUserFailed(t *testing.T) {
	u := createRegularUser()
	ctrl := &UserControllerImpl{
		Interactor: &mocks.UserInteractorMock{
			AuthorizeImpl: func(string) (entities.User, error) {
				return u, nil
			},
			CreateImpl: func(entities.User) (entities.User, error) {
				return u, nil
			},
		},
	}

	userJSON, _ := json.Marshal(u)

	r := mocks.NewHTTPRequestMock(userJSON)
	r.Header.Set(XAuthToken, "124")

	w := mocks.NewResponseWriterMock()

	ctrl.Create(w, r)

	if w.WrittenHeader != http.StatusForbidden {
		t.Errorf("Expect %d status code, received %d", http.StatusForbidden, w.WrittenHeader)
	}
}

func TestUserCreateCorruptedJSON(t *testing.T) {
	u := createAdminUser()
	ctrl := &UserControllerImpl{
		Interactor: &mocks.UserInteractorMock{
			AuthorizeImpl: func(string) (entities.User, error) {
				return u, nil
			},
			CreateImpl: func(entities.User) (entities.User, error) {
				return u, nil
			},
		},
	}

	r := mocks.NewHTTPRequestMock([]byte("corrupted json"))
	r.Header.Set(XAuthToken, "124")

	w := mocks.NewResponseWriterMock()

	ctrl.Create(w, r)

	if w.WrittenHeader != http.StatusBadRequest {
		t.Errorf("Expect %d status code, received %d", http.StatusBadRequest, w.WrittenHeader)
	}
}

func TestUserCreateFailed(t *testing.T) {
	u := createAdminUser()
	ctrl := &UserControllerImpl{
		Interactor: &mocks.UserInteractorMock{
			AuthorizeImpl: func(string) (entities.User, error) {
				return u, nil
			},
			CreateImpl: func(entities.User) (entities.User, error) {
				return u, fmt.Errorf("Unable to create user")
			},
		},
	}
	userJSON, _ := json.Marshal(u)

	r := mocks.NewHTTPRequestMock(userJSON)
	r.Header.Set(XAuthToken, "124")

	w := mocks.NewResponseWriterMock()

	ctrl.Create(w, r)

	if w.WrittenHeader != http.StatusNotFound {
		t.Errorf("Expect %d status code, received %d", http.StatusNotFound, w.WrittenHeader)
	}
}

func TestUserCreateSuccess(t *testing.T) {
	u := createAdminUser()
	ctrl := &UserControllerImpl{
		Interactor: &mocks.UserInteractorMock{
			AuthorizeImpl: func(string) (entities.User, error) {
				return u, nil
			},
			CreateImpl: func(entities.User) (entities.User, error) {
				return u, nil
			},
		},
	}
	userJSON, _ := json.Marshal(u)

	r := mocks.NewHTTPRequestMock(userJSON)
	r.Header.Set(XAuthToken, "124")

	w := mocks.NewResponseWriterMock()

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
	ctrl := &UserControllerImpl{
		Interactor: &mocks.UserInteractorMock{
			AuthorizeImpl: func(string) (entities.User, error) {
				return u, nil
			},
			CreateImpl: func(entities.User) (entities.User, error) {
				return u, nil
			},
		},
	}

	r := mocks.NewHTTPRequestMock([]byte("Corrupted JSON"))

	w := mocks.NewResponseWriterMock()

	ctrl.Authenticate(w, r)

	if w.WrittenHeader != http.StatusBadRequest {
		t.Errorf("Expect %d status code, received %d", http.StatusBadRequest, w.WrittenHeader)
	}
}

func TestAuthenticateFailed(t *testing.T) {
	u := createAdminUser()
	ctrl := &UserControllerImpl{
		Interactor: &mocks.UserInteractorMock{
			AuthorizeImpl: func(string) (entities.User, error) {
				return u, fmt.Errorf("User interactor creation failed")
			},
			AuthenticateImpl: func(string, string) (entities.User, string, error) {
				return u, "", fmt.Errorf("Authentication failed")
			},
		},
	}
	userJSON, _ := json.Marshal(u)

	r := mocks.NewHTTPRequestMock(userJSON)
	r.Header.Set(XAuthToken, "124")

	w := mocks.NewResponseWriterMock()

	ctrl.Authenticate(w, r)

	if w.WrittenHeader != http.StatusForbidden {
		t.Errorf("Expect %d status code, received %d", http.StatusForbidden, w.WrittenHeader)
	}
}

func TestAuthenticateSuccess(t *testing.T) {
	u := createAdminUser()
	token := "123"
	ctrl := &UserControllerImpl{
		Interactor: &mocks.UserInteractorMock{
			AuthorizeImpl: func(string) (entities.User, error) {
				return u, fmt.Errorf("User interactor creation failed")
			},
			AuthenticateImpl: func(string, string) (entities.User, string, error) {
				return u, "123", nil
			},
		},
	}
	userJSON, _ := json.Marshal(u)

	r := mocks.NewHTTPRequestMock(userJSON)
	r.Header.Set(XAuthToken, token)

	w := mocks.NewResponseWriterMock()

	ctrl.Authenticate(w, r)

	if w.WrittenHeader != http.StatusOK {
		t.Errorf("Expect %d status code, received %d", http.StatusOK, w.WrittenHeader)
	}

	givenToken := w.Header().Get(XAuthToken)
	if givenToken != token {
		t.Errorf("Expect token %s, received %s", token, w.WrittenHeader)
	}
}
