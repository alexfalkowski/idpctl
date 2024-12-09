package pipeline

import (
	"context"

	"github.com/alexfalkowski/go-service/meta"
	"github.com/alexfalkowski/go-service/net/http/rest"
	"github.com/alexfalkowski/go-service/runtime"
)

// TriggerPipeline for cmd.
func (s *Service) TriggerPipeline(ctx context.Context) context.Context {
	res, err := s.request().Post(s.cfg.Host + "/pipelines/" + *TriggerFlag + "/trigger")
	runtime.Must(err)
	runtime.Must(rest.Error(res))

	ctx = meta.WithAttribute(ctx, "response", meta.String(res.Body()))

	return ctx
}
