package controllers

import (
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

// NewResponseWriterMock constructs ResponseWriterMock instance.
func NewResponseWriterMock() *ResponseWriterMock {
	return &ResponseWriterMock{
		HeaderImpl: http.Header{},
	}
}

// NewHTTPRequestMock constructs http.Request instance.
func NewHTTPRequestMock(token string, userJSON []byte) *http.Request {
	r := &http.Request{Header: http.Header{}}
	r.Header.Set(XAuthToken, token)
	r.Body = &RequestBodyMock{
		ReadImpl: userJSON,
	}

	return r
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
