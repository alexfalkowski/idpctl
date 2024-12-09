package pipeline

import (
	"context"
	"strings"

	"github.com/alexfalkowski/go-service/flags"
	"github.com/alexfalkowski/go-service/meta"
	"github.com/alexfalkowski/go-service/net/http/rest"
	"github.com/alexfalkowski/go-service/os"
	"github.com/alexfalkowski/go-service/runtime"
	"github.com/alexfalkowski/go-service/transport/http"
	"github.com/alexfalkowski/idpctl/cmd/runner"
	"github.com/alexfalkowski/idpctl/pipeline"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var (
	// CreateFlag defines the pipeline path to create a pipeline.
	CreateFlag = flags.String()

	// GetFlag defines a pipeline id to be retrieved.
	GetFlag = flags.String()

	// GetFlag defines the id and pipeline path to update.
	UpdateFlag = flags.String()
)

// Start for redis.
func Start(lc fx.Lifecycle, cfg *pipeline.Config, logger *zap.Logger) {
	rt := http.NewRoundTripper()
	client := rest.NewClient(rest.WithClientRoundTripper(rt), rest.WithClientTimeout("10s"))

	var (
		operation string
		fn        runner.StartFn
	)

	switch {
	case flags.IsStringSet(CreateFlag):
		fn = func(ctx context.Context) context.Context {
			token, err := os.ReadBase64File("secrets/token")
			runtime.Must(err)

			pipeline, err := os.ReadFile(*CreateFlag)
			runtime.Must(err)

			res, err := client.R().SetHeader("Content-Type", "application/json").SetAuthToken(token).SetBody(pipeline).Post(cfg.Host + "/pipelines")
			runtime.Must(err)

			ctx = meta.WithAttribute(ctx, "response", meta.String(res.Body()))

			return ctx
		}
		operation = "created the pipeline"
	case flags.IsStringSet(GetFlag):
		fn = func(ctx context.Context) context.Context {
			token, err := os.ReadBase64File("secrets/token")
			runtime.Must(err)

			res, err := client.R().SetHeader("Content-Type", "application/json").SetAuthToken(token).Get(cfg.Host + "/pipelines/" + *GetFlag)
			runtime.Must(err)

			ctx = meta.WithAttribute(ctx, "response", meta.String(res.Body()))

			return ctx
		}
		operation = "retrieved the pipeline"
	case flags.IsStringSet(UpdateFlag):
		fn = func(ctx context.Context) context.Context {
			token, err := os.ReadBase64File("secrets/token")
			runtime.Must(err)

			split := strings.Split(*UpdateFlag, ":")

			pipeline, err := os.ReadFile(split[1])
			runtime.Must(err)

			res, err := client.R().SetHeader("Content-Type", "application/json").SetAuthToken(token).SetBody(pipeline).Put(cfg.Host + "/pipelines/" + split[0])
			runtime.Must(err)

			ctx = meta.WithAttribute(ctx, "response", meta.String(res.Body()))

			return ctx
		}
		operation = "updated the pipeline"
	}

	opts := &runner.Options{Lifecycle: lc, Logger: logger, Fn: fn}
	runner.Start("pipeline", operation, opts)
}
