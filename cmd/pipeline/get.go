package pipeline

import (
	"context"

	"github.com/alexfalkowski/go-service/meta"
	"github.com/alexfalkowski/go-service/net/http/rest"
	"github.com/alexfalkowski/go-service/runtime"
)

// GetPipeline for cmd.
func (s *Service) GetPipeline(ctx context.Context) context.Context {
	res, err := s.request().Get(s.cfg.Host + "/pipelines/" + *GetFlag)
	runtime.Must(err)
	runtime.Must(rest.Error(res))

	ctx = meta.WithAttribute(ctx, "response", meta.String(res.Body()))

	return ctx
}
