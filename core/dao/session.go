package dao

import "time"

// Session temporary implementation of session to store in SQL database,
// but should be moved to appropriate storage.
//go:generate
//reform:sessions
type Session struct {
	ID         string    `reform:"id,pk"`
	UserID     string    `reform:"user_id"`
	RemoteAddr string    `reform:"ip"`
	CreatedAt  time.Time `reform:"created_at"`
	ExpiresAt  time.Time `reform:"expires_at"`
}
