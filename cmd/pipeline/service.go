package pipeline

import (
	"context"
	"strings"

	"github.com/alexfalkowski/go-service/meta"
	"github.com/alexfalkowski/go-service/os"
	"github.com/alexfalkowski/go-service/runtime"
	"github.com/alexfalkowski/idpctl/pipeline"
	"github.com/go-resty/resty/v2"
)

// Service for pipeline.
type Service struct {
	client *resty.Client
	cfg    *pipeline.Config
}

// NewService for pipeline.
func NewService(client *resty.Client, cfg *pipeline.Config) *Service {
	return &Service{client: client, cfg: cfg}
}

// Create a pipeline.
func (s *Service) Create(ctx context.Context) context.Context {
	token, err := os.ReadBase64File("secrets/token")
	runtime.Must(err)

	pipeline, err := os.ReadFile(*CreateFlag)
	runtime.Must(err)

	res, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(token).SetBody(pipeline).
		Post(s.cfg.Host + "/pipelines")
	runtime.Must(err)

	ctx = meta.WithAttribute(ctx, "response", meta.String(res.Body()))

	return ctx
}

// Get a pipelines.
func (s *Service) Get(ctx context.Context) context.Context {
	token, err := os.ReadBase64File("secrets/token")
	runtime.Must(err)

	res, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(token).
		Get(s.cfg.Host + "/pipelines/" + *GetFlag)
	runtime.Must(err)

	ctx = meta.WithAttribute(ctx, "response", meta.String(res.Body()))

	return ctx
}

// Update a pipeline.
func (s *Service) Update(ctx context.Context) context.Context {
	token, err := os.ReadBase64File("secrets/token")
	runtime.Must(err)

	split := strings.Split(*UpdateFlag, ":")

	pipeline, err := os.ReadFile(split[1])
	runtime.Must(err)

	res, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(token).SetBody(pipeline).
		Put(s.cfg.Host + "/pipelines/" + split[0])
	runtime.Must(err)

	ctx = meta.WithAttribute(ctx, "response", meta.String(res.Body()))

	return ctx
}

// Delete a pipeline.
func (s *Service) Delete(ctx context.Context) context.Context {
	token, err := os.ReadBase64File("secrets/token")
	runtime.Must(err)

	res, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(token).
		Delete(s.cfg.Host + "/pipelines/" + *DeleteFlag)
	runtime.Must(err)

	ctx = meta.WithAttribute(ctx, "response", meta.String(res.Body()))

	return ctx
}

// Delete a pipeline.
func (s *Service) Trigger(ctx context.Context) context.Context {
	token, err := os.ReadBase64File("secrets/token")
	runtime.Must(err)

	res, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(token).
		Post(s.cfg.Host + "/pipelines/" + *TriggerFlag + "/trigger")
	runtime.Must(err)

	ctx = meta.WithAttribute(ctx, "response", meta.String(res.Body()))

	return ctx
}
