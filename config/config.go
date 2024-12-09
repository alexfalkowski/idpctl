package config

import (
	"github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/config"
	"github.com/alexfalkowski/idpctl/pipeline"
)

// NewConfig for config.
func NewConfig(i *cmd.InputConfig) (*Config, error) {
	c := &Config{}

	return c, i.Decode(c)
}

// Config for the client.
type Config struct {
	Pipeline       *pipeline.Config `yaml:"pipeline,omitempty" json:"pipeline,omitempty" toml:"pipeline,omitempty"`
	*config.Config `yaml:",inline" json:",inline" toml:",inline"`
}

func decorateConfig(cfg *Config) *config.Config {
	return cfg.Config
}

func pipelineConfig(cfg *Config) *pipeline.Config {
	return cfg.Pipeline
}
