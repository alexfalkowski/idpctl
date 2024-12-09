package pipeline

import (
	"context"
	"strings"

	"github.com/alexfalkowski/go-service/meta"
	"github.com/alexfalkowski/go-service/net/http/rest"
	"github.com/alexfalkowski/go-service/os"
	"github.com/alexfalkowski/go-service/runtime"
)

// UpdatePipeline for cmd.
func (s *Service) UpdatePipeline(ctx context.Context) context.Context {
	split := strings.Split(*UpdateFlag, ":")
	if len(split) != 2 {
		runtime.Must(ErrInvalidUpdateFormat)
	}

	pipeline, err := os.ReadFile(split[1])
	runtime.Must(err)

	res, err := s.request().SetBody(pipeline).Put(s.cfg.Host + "/pipelines/" + split[0])
	runtime.Must(err)
	runtime.Must(rest.Error(res))

	ctx = meta.WithAttribute(ctx, "response", meta.String(res.Body()))

	return ctx
}
