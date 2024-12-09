package pipeline

import (
	"errors"

	"github.com/alexfalkowski/idpctl/pipeline"
	"github.com/go-resty/resty/v2"
)

// ErrInvalidUpdateFormat indicates that the format is not how it's expected.
var ErrInvalidUpdateFormat = errors.New("invalid update format, this should be id:path")

// Service for pipeline.
type Service struct {
	client *resty.Client
	cfg    *pipeline.Config
	token  string
}

// NewService for pipeline.
func NewService(client *resty.Client, cfg *pipeline.Config) (*Service, error) {
	token, err := cfg.GetToken()
	if err != nil {
		return nil, err
	}

	return &Service{client: client, cfg: cfg, token: token}, nil
}

func (s *Service) request() *resty.Request {
	return s.client.R().
		SetHeaders(map[string]string{
			"Content-Type": "application/json",
			"Accept":       "application/json",
		}).
		SetAuthToken(s.token)
}
