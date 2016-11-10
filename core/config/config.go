package config

// AppConfig application config interface.
type AppConfig interface {
	Secret() string
}

// AppConfigImpl application config implementation.
type AppConfigImpl struct {
	SecretVal string
}

// Secret returns hashing secret for app.
func (config *AppConfigImpl) Secret() string {
	return config.SecretVal
}
