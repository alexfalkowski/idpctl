package pipeline

import (
	"context"
	"strings"

	"github.com/alexfalkowski/go-service/meta"
	"github.com/alexfalkowski/go-service/net/http/rest"
	"github.com/alexfalkowski/go-service/os"
	"github.com/alexfalkowski/go-service/runtime"
	"github.com/alexfalkowski/idpctl/pipeline"
	"github.com/go-resty/resty/v2"
)

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

// Create a pipeline.
func (s *Service) Create(ctx context.Context) context.Context {
	pipeline, err := os.ReadFile(*CreateFlag)
	runtime.Must(err)

	res, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(s.token).
		SetBody(pipeline).
		Post(s.cfg.Host + "/pipelines")
	runtime.Must(err)
	runtime.Must(rest.Error(res))

	ctx = meta.WithAttribute(ctx, "response", meta.String(res.Body()))

	return ctx
}

// Get a pipelines.
func (s *Service) Get(ctx context.Context) context.Context {
	res, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(s.token).
		Get(s.cfg.Host + "/pipelines/" + *GetFlag)
	runtime.Must(err)
	runtime.Must(rest.Error(res))

	ctx = meta.WithAttribute(ctx, "response", meta.String(res.Body()))

	return ctx
}

// Update a pipeline.
func (s *Service) Update(ctx context.Context) context.Context {
	split := strings.Split(*UpdateFlag, ":")

	pipeline, err := os.ReadFile(split[1])
	runtime.Must(err)

	res, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(s.token).
		SetBody(pipeline).
		Put(s.cfg.Host + "/pipelines/" + split[0])
	runtime.Must(err)
	runtime.Must(rest.Error(res))

	ctx = meta.WithAttribute(ctx, "response", meta.String(res.Body()))

	return ctx
}

// Delete a pipeline.
func (s *Service) Delete(ctx context.Context) context.Context {
	res, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(s.token).
		Delete(s.cfg.Host + "/pipelines/" + *DeleteFlag)
	runtime.Must(err)
	runtime.Must(rest.Error(res))

	ctx = meta.WithAttribute(ctx, "response", meta.String(res.Body()))

	return ctx
}

// Delete a pipeline.
func (s *Service) Trigger(ctx context.Context) context.Context {
	res, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(s.token).
		Post(s.cfg.Host + "/pipelines/" + *TriggerFlag + "/trigger")
	runtime.Must(err)
	runtime.Must(rest.Error(res))

	ctx = meta.WithAttribute(ctx, "response", meta.String(res.Body()))

	return ctx
}
