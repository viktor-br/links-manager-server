package handlers

import (
	"github.com/viktor-br/links-manager-server/app/mocks"
	"net/http"
	"testing"
)

func TestUserHandlerSuccess(t *testing.T) {
	userControllerMock := mocks.NewUserControllerMock()
	userInteractorMock := mocks.NewUserInteractorMock()
	logger := mocks.NewLoggerMock()
	handler := NewUserHandler(userControllerMock, userInteractorMock, logger)
	r := mocks.NewHTTPRequestMock([]byte{})
	r.Method = http.MethodPut
	w := mocks.NewResponseWriterMock()
	handler.ServeHTTP(w, r)

	if len(userControllerMock.Calls) != 1 {
		t.Errorf("Expect 1 call to controller")
		if userControllerMock.Calls[0] != "authenticate" {
			t.Errorf("Expect controller.Authenticatr() call, received %s", userControllerMock.Calls[0])
		}
	}
}

func TestUserHandlerMethodNotAllowed(t *testing.T) {
	userControllerMock := mocks.NewUserControllerMock()
	userInteractorMock := mocks.NewUserInteractorMock()
	logger := mocks.NewLoggerMock()
	handler := NewUserHandler(userControllerMock, userInteractorMock, logger)

	checkNotAllowedMethods(handler, []string{http.MethodPut}, t)
}

func TestUserAuthenticateHandlerSuccess(t *testing.T) {
	userControllerMock := mocks.NewUserControllerMock()
	userInteractorMock := mocks.NewUserInteractorMock()
	logger := mocks.NewLoggerMock()
	handler := NewUserAuthenticateHandler(userControllerMock, userInteractorMock, logger)
	r := mocks.NewHTTPRequestMock([]byte{})
	r.Method = http.MethodPost
	w := mocks.NewResponseWriterMock()
	handler.ServeHTTP(w, r)

	if len(userControllerMock.Calls) != 1 {
		t.Errorf("Expect 1 call to controller")
		if userControllerMock.Calls[0] != "authenticate" {
			t.Errorf("Expect controller.Authenticatr() call, received %s", userControllerMock.Calls[0])
		}
	}
}

func TestUserAuthenticateHandlerMethodNotAllowed(t *testing.T) {
	userControllerMock := mocks.NewUserControllerMock()
	userInteractorMock := mocks.NewUserInteractorMock()
	logger := mocks.NewLoggerMock()
	handler := NewUserAuthenticateHandler(userControllerMock, userInteractorMock, logger)

	checkNotAllowedMethods(handler, []string{http.MethodPost}, t)
}

func checkNotAllowedMethods(handler http.Handler, allowedMethods []string, t *testing.T) {
	notAllowedMethods := []string{}
	r := mocks.NewHTTPRequestMock([]byte{})
	w := mocks.NewResponseWriterMock()

	methods := [...]string{
		http.MethodHead,
		http.MethodDelete,
		http.MethodGet,
		http.MethodOptions,
		http.MethodPatch,
		http.MethodPut,
		http.MethodPost,
		http.MethodTrace,
	}

	for i := 0; i < len(methods); i++ {
		found := false
		for j := 0; j < len(allowedMethods); j++ {
			if methods[i] == allowedMethods[j] {
				found = true
				break
			}
		}
		if !found {
			notAllowedMethods = append(notAllowedMethods, methods[i])
		}
	}
	for i := 0; i < len(notAllowedMethods); i++ {
		r.Method = notAllowedMethods[i]
		w.WriteHeader(0)
		handler.ServeHTTP(w, r)

		// Response should be OK
		if w.WrittenHeader != http.StatusMethodNotAllowed {
			t.Errorf(
				"Expect HTTP status %d, received %d for method %s",
				http.StatusMethodNotAllowed,
				w.WrittenHeader,
				notAllowedMethods[i],
			)
		}
	}
}
