package log

// Log record keys
const (
	// LogController identifies controller.
	LogController = "controller"
	// LogMessage for info message.
	LogMessage = "msg"
	// LogHTTPStatus for HTTP response status.
	LogHTTPStatus = "code"
	// LogToken for token used for request if user is not identified.
	LogToken = "token"
	// LogUserID for user identifier.
	LogUserID = "uid"
	// LogRemoteAddr for remote ip of request.
	LogRemoteAddr = "ip"
	// LogRequestURI for request path
	LogRequestURI = "url"
)

// Logger interface to hide using external library from the code.
type Logger interface {
	Log(keyvals ...interface{}) error
}
