package pipeline

import (
	"github.com/alexfalkowski/go-service/os"
)

// Config for the pipeline.
type Config struct {
	Host  string `yaml:"host,omitempty" json:"host,omitempty" toml:"host,omitempty"`
	Token string `yaml:"token,omitempty" json:"token,omitempty" toml:"token,omitempty"`
}

// GetToken from the config.
func (c *Config) GetToken() (string, error) {
	return os.ReadBase64File(c.Token)
}
