package pipeline

import (
	"context"

	"github.com/alexfalkowski/go-service/meta"
	"github.com/alexfalkowski/go-service/net/http/rest"
	"github.com/alexfalkowski/go-service/os"
	"github.com/alexfalkowski/go-service/runtime"
)

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
