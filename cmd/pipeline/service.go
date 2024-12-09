package pipeline

import (
	"context"
	"errors"
	"strings"

	"github.com/alexfalkowski/go-service/meta"
	"github.com/alexfalkowski/go-service/net/http/rest"
	"github.com/alexfalkowski/go-service/os"
	"github.com/alexfalkowski/go-service/runtime"
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

// CreatePipeline for cmd.
func (s *Service) CreatePipeline(ctx context.Context) context.Context {
	pipeline, err := os.ReadFile(*CreateFlag)
	runtime.Must(err)

	res, err := s.request().SetBody(pipeline).Post(s.cfg.Host + "/pipelines")
	runtime.Must(err)
	runtime.Must(rest.Error(res))

	ctx = meta.WithAttribute(ctx, "response", meta.String(res.Body()))

	return ctx
}

// GetPipeline for cmd.
func (s *Service) GetPipeline(ctx context.Context) context.Context {
	res, err := s.request().Get(s.cfg.Host + "/pipelines/" + *GetFlag)
	runtime.Must(err)
	runtime.Must(rest.Error(res))

	ctx = meta.WithAttribute(ctx, "response", meta.String(res.Body()))

	return ctx
}

// UpdatePipeline for cmd.
func (s *Service) UpdatePipeline(ctx context.Context) context.Context {
	split := strings.Split(*UpdateFlag, ":")
	if len(split) != 2 {
		runtime.Must(ErrInvalidUpdateFormat)
	}

	pipeline, err := os.ReadFile(split[1])
	runtime.Must(err)

	res, err := s.request().SetAuthToken(s.token).SetBody(pipeline).Put(s.cfg.Host + "/pipelines/" + split[0])
	runtime.Must(err)
	runtime.Must(rest.Error(res))

	ctx = meta.WithAttribute(ctx, "response", meta.String(res.Body()))

	return ctx
}

// DeletePipeline for cmd.
func (s *Service) DeletePipeline(ctx context.Context) context.Context {
	res, err := s.request().Delete(s.cfg.Host + "/pipelines/" + *DeleteFlag)
	runtime.Must(err)
	runtime.Must(rest.Error(res))

	ctx = meta.WithAttribute(ctx, "response", meta.String(res.Body()))

	return ctx
}

// TriggerPipeline for cmd.
func (s *Service) TriggerPipeline(ctx context.Context) context.Context {
	res, err := s.request().Post(s.cfg.Host + "/pipelines/" + *TriggerFlag + "/trigger")
	runtime.Must(err)
	runtime.Must(rest.Error(res))

	ctx = meta.WithAttribute(ctx, "response", meta.String(res.Body()))

	return ctx
}

func (s *Service) request() *resty.Request {
	return s.client.R().
		SetHeaders(map[string]string{
			"Content-Type": "application/json",
			"Accept":       "application/json",
		}).
		SetAuthToken(s.token)
}
