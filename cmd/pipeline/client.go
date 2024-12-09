package pipeline

import (
	"github.com/alexfalkowski/go-service/net/http/rest"
	"github.com/alexfalkowski/go-service/transport/http"
	"github.com/go-resty/resty/v2"
)

// NewClient for pipeline.
func NewClient() *resty.Client {
	rt := http.NewRoundTripper()
	client := rest.NewClient(rest.WithClientRoundTripper(rt), rest.WithClientTimeout("10s"))

	return client
}
