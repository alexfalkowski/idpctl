package cmd

import (
	"github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/compress"
	"github.com/alexfalkowski/go-service/debug"
	"github.com/alexfalkowski/go-service/encoding"
	"github.com/alexfalkowski/go-service/env"
	"github.com/alexfalkowski/go-service/feature"
	"github.com/alexfalkowski/go-service/sync"
	"github.com/alexfalkowski/go-service/telemetry"
	"github.com/alexfalkowski/go-service/transport"
	"github.com/alexfalkowski/idpctl/config"
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	cmd.Module, env.Module,
	transport.Module, debug.Module,
	sync.Module, compress.Module, encoding.Module,
	feature.Module, telemetry.Module, config.Module,
	fx.Provide(NewVersion),
)
