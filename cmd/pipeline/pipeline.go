package pipeline

import (
	"github.com/alexfalkowski/go-service/flags"
	"github.com/alexfalkowski/idpctl/cmd/runner"
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

	// DeleteFlag defines a pipeline to be deleted by the provided id.
	DeleteFlag = flags.String()

	// TriggerFlag defines a pipeline to be triggered by the provided id.
	TriggerFlag = flags.String()
)

// Start for redis.
func Start(lc fx.Lifecycle, service *Service, logger *zap.Logger) error {
	var (
		operation string
		fn        runner.StartFn
	)

	switch {
	case flags.IsStringSet(CreateFlag):
		fn = service.Create
		operation = "created the pipeline"
	case flags.IsStringSet(GetFlag):
		fn = service.Get
		operation = "retrieved the pipeline"
	case flags.IsStringSet(UpdateFlag):
		fn = service.Update
		operation = "updated the pipeline"
	case flags.IsStringSet(DeleteFlag):
		fn = service.Delete
		operation = "deleted the pipeline"
	case flags.IsStringSet(TriggerFlag):
		fn = service.Trigger
		operation = "triggered the pipeline"
	}

	opts := &runner.Options{Lifecycle: lc, Logger: logger, Fn: fn}
	runner.Start("pipeline", operation, opts)

	return nil
}
