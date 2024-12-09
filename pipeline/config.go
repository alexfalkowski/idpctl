package pipeline

// Config for the pipeline.
type Config struct {
	Host string `yaml:"host,omitempty" json:"host,omitempty" toml:"host,omitempty"`
}
