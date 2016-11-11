package mocks

import (
	"github.com/viktor-br/links-manager-server/core/entities"
	"io"
	"net/http"
)

// ResponseWriterMock mocks http.ResponseWriter interface with ability to customise behaviour.
type ResponseWriterMock struct {
	WrittenHeader  int
	WrittenContent []byte
	HeaderImpl     http.Header
}

// RequestBodyMock mocks http.Request.Body with ability to customise behaviour.
type RequestBodyMock struct {
	done     bool
	ReadImpl []byte
}

// UserInteractorMock mocks UserInteractorImpl
type UserInteractorMock struct {
	AuthenticateImpl func(username, password string) (*entities.User, *entities.Session, error)
	AuthorizeImpl    func(string) (*entities.User, error)
	CreateImpl       func(*entities.User) error
}

// UserControllerMock mocks UserController
type UserControllerMock struct {
	Calls []string
}

// LoggerMock mocks logger.
type LoggerMock struct {
}

// UserRepositoryMock mocks UserRepository.
type UserRepositoryMock struct {
	FindByUsernameImpl func(username string) (*entities.User, error)
	StoreImpl          func(user *entities.User) error
}

// Log mocks logger main method.
func (loggerMock *LoggerMock) Log(keyvals ...interface{}) error {
	return nil
}

// Create mocks UserControllers method call.
func (userControllerMock *UserControllerMock) Create(w http.ResponseWriter, r *http.Request) {
	userControllerMock.Calls = append(userControllerMock.Calls, "create")
}

// Authenticate mocks UserControllers method call.
func (userControllerMock *UserControllerMock) Authenticate(w http.ResponseWriter, r *http.Request) {
	userControllerMock.Calls = append(userControllerMock.Calls, "authenticate")
}

// Log mocks UserControllers method call.
func (userControllerMock *UserControllerMock) Log(args ...interface{}) {
	userControllerMock.Calls = append(userControllerMock.Calls, "log")
}

// Authenticate mocks method via implementation method (allow simulate errors and etc).
func (userInteractorMock UserInteractorMock) Authenticate(username, password string) (*entities.User, *entities.Session, error) {
	return userInteractorMock.AuthenticateImpl(username, password)
}

// Authorize mocks method via implementation method (allow simulate errors and etc).
func (userInteractorMock UserInteractorMock) Authorize(token string) (*entities.User, error) {
	return userInteractorMock.AuthorizeImpl(token)
}

// Create mocks method via implementation method (allow simulate errors and etc).
func (userInteractorMock UserInteractorMock) Create(user *entities.User) error {
	return userInteractorMock.CreateImpl(user)
}

// NewLoggerMock creates logger mock instance.
func NewLoggerMock() *LoggerMock {
	return &LoggerMock{}
}

// NewResponseWriterMock constructs ResponseWriterMock instance.
func NewResponseWriterMock() *ResponseWriterMock {
	return &ResponseWriterMock{
		HeaderImpl: http.Header{},
	}
}

// NewHTTPRequestMock constructs http.Request instance.
func NewHTTPRequestMock(userJSON []byte) *http.Request {
	r := &http.Request{Header: http.Header{}}
	r.Body = &RequestBodyMock{
		ReadImpl: userJSON,
	}

	return r
}

// NewUserControllerMock creates UserController mock.
func NewUserControllerMock() *UserControllerMock {
	return &UserControllerMock{}
}

// NewUserInteractorMock creates new UserInteractor mock.
func NewUserInteractorMock() *UserInteractorMock {
	return &UserInteractorMock{}
}

// Read mocks ResponseWriterMock method.
func (body *RequestBodyMock) Read(p []byte) (int, error) {
	if body.done {
		return 0, io.EOF
	}
	for i, b := range []byte(body.ReadImpl) {
		p[i] = b
	}
	body.done = true

	return len(body.ReadImpl), nil
}

// Close mocks appropriate method of io.ReadCloser
func (body *RequestBodyMock) Close() error {
	return nil
}

// Write mocks appropriate method of http.ResponseWriter with custom behaviour.
func (writer *ResponseWriterMock) Write(content []byte) (int, error) {
	writer.WrittenContent = content

	return len(content), nil
}

// Header mocks appropriate method of http.ResponseWriter with custom behaviour.
func (writer *ResponseWriterMock) Header() http.Header {
	if writer.HeaderImpl == nil {
		writer.HeaderImpl = http.Header{}
	}
	return writer.HeaderImpl
}

// WriteHeader mocks appropriate method of http.ResponseWriter and save assigned status code to struct variable.
func (writer *ResponseWriterMock) WriteHeader(statusCode int) {
	writer.WrittenHeader = statusCode
}

// FindByUsername mocks UserRepository method
func (userRepositoryMock *UserRepositoryMock) FindByUsername(username string) (*entities.User, error) {
	return userRepositoryMock.FindByUsernameImpl(username)
}

// Store mocks UserRepository Store method
func (userRepositoryMock *UserRepositoryMock) Store(user *entities.User) error {
	return userRepositoryMock.StoreImpl(user)
}
