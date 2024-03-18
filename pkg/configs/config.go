package configs

import (
	"github.com/joeshaw/envdecode"
	"github.com/pkg/errors"
)

// ServerConfig represents all HTTP server configuration options.
type ServerConfig struct {
	Host string `env:"SERVER_HOST"`
	Port int    `env:"SERVER_PORT"`
}

// NewServerConfig constructs a new instance of ServerConfig via decoding
// the mapped env vars with envdecode library.
func NewServerConfig() (*ServerConfig, error) {
	var config ServerConfig
	err := envdecode.Decode(&config)

	if err != nil {
		if !errors.Is(err, envdecode.ErrNoTargetFieldsAreSet) {
			return nil, errors.WithStack(err)
		}
	}

	return &config, nil
}
