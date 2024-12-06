package cmd

import (
	"github.com/alexfalkowski/go-service/compress"
	"github.com/alexfalkowski/go-service/encoding"
	"github.com/alexfalkowski/go-service/feature"
	"github.com/alexfalkowski/go-service/sync"
	"github.com/alexfalkowski/go-service/telemetry"
	"github.com/alexfalkowski/idpctl/client"
	"github.com/alexfalkowski/idpctl/config"
	"go.uber.org/fx"
)

// ClientOptions for cmd.
var ClientOptions = []fx.Option{
	sync.Module, compress.Module, encoding.Module,
	feature.Module, telemetry.Module,
	config.Module, client.Module, Module,
}
