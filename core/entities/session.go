package entities

import "time"

// Session entity.
type Session struct {
	ID         string
	User       *User
	UserAgent  string
	RemoteAddr string
	CreatedOn  time.Time
	UpdatedOn  time.Time
	ExpiresOn  time.Time
}

// IsValid checks if session instance is valid.
func (session *Session) IsValid() bool {
	return session.ID != "" && session.User.ID != ""
}

// IsExpired checks if session is expired.
func (session *Session) IsExpired() bool {
	return session.ExpiresOn.Sub(time.Now()) > 0
}
