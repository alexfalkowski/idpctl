package pipeline

import (
	"context"

	"github.com/alexfalkowski/go-service/meta"
	"github.com/alexfalkowski/go-service/net/http/rest"
	"github.com/alexfalkowski/go-service/runtime"
)

// DeletePipeline for cmd.
func (s *Service) DeletePipeline(ctx context.Context) context.Context {
	res, err := s.request().Delete(s.cfg.Host + "/pipelines/" + *DeleteFlag)
	runtime.Must(err)
	runtime.Must(rest.Error(res))

	ctx = meta.WithAttribute(ctx, "response", meta.String(res.Body()))

	return ctx
}
