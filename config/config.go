package config

import (
	"github.com/alexfalkowski/go-service/config"
	"github.com/alexfalkowski/idpctl/pipeline"
)

// Config for the client.
type Config struct {
	Pipeline       *pipeline.Config `yaml:"pipeline,omitempty" json:"pipeline,omitempty" toml:"pipeline,omitempty"`
	*config.Config `yaml:",inline" json:",inline" toml:",inline"`
}

// Valid or error.
func (c Config) Valid() error {
	if c.Pipeline == nil || c.Config == nil {
		return config.ErrInvalidConfig
	}

	return c.Config.Valid()
}

func decorateConfig(cfg *Config) *config.Config {
	return cfg.Config
}

func pipelineConfig(cfg *Config) *pipeline.Config {
	return cfg.Pipeline
}
